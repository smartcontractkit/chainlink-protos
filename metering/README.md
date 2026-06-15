# Metering

Protobuf definitions for the billable events emitted by services that manage
durable resources (trigger registrations, log filters, workflow registrations,
...). Records are published via Beholder and consumed by the billing pipeline.

There are two messages, and **each carries exactly one resource**, identified
entirely by its `ResourceIdentity`:

- `MeterRecord` — a single **state-transition event** (RESERVE / RELEASE /
  UPDATE / USAGE) describing one lifecycle edge of one durable resource.
- `MeterSnapshot` — the **action-less periodic utilization** of one active
  resource (the liveness / utilization-over-time signal). The resource manager
  emits one `MeterSnapshot` per active resource each interval.

Both embed a structured `ResourceIdentity`.

Derived Kafka subjects (domain `beholder__platform__messages` from
`chip-platform.json` + schema entity):

- `beholder__platform__messages-metering.v1.MeterRecord`
- `beholder__platform__messages-metering.v1.MeterSnapshot`

## Resource identity

`ResourceIdentity` (in `identity.proto`) is the first-class, structured
identity embedded by every metering message. Downstream aggregators, analytics,
and UI treat each dimension as a discrete column rather than parsing a dotted
string or carrying values out-of-band on telemetry. Its nine fields:

| Field           | Meaning                                                                                  |
| --------------- | ---------------------------------------------------------------------------------------- |
| `product`       | Deployment product, e.g. `cre-mainline`. Coarse billing-rollup dimension.                |
| `environment`   | Deployment environment, e.g. `production`, `staging`. Coarse billing-rollup dimension.   |
| `zone`          | Deployment zone, e.g. `wf-zone-a`. Coarse billing-rollup dimension.                      |
| `don_id`        | DON the emitting service belongs to. Coarse billing-rollup dimension.                    |
| `node_id`       | Node identity (the node's CSA public key). Coarse billing-rollup dimension.              |
| `service`       | Stable service constant (the old `entity`), e.g. `cron-trigger`. Coarse rollup dimension.|
| `resource`      | Resource pool, e.g. `trigger_registrations`, `log_filters`.                              |
| `resource_type` | Billing unit for the value, e.g. `operations`, `log_filter_addresses`.                   |
| `resource_id`   | The **physical/logical resource identity** (see below).                                  |

`product` / `environment` / `zone` / `don_id` / `node_id` are the
deployment+DON+node dimensions used for coarse billing rollup. `service` is the
stable service constant that replaced the old opaque `entity`.

`resource_id` is the **physical/logical resource identity, workflow-independent
where a shared physical resource exists**:

- **EVM log filters** — content hash of `chain_selector` + canonicalized
  addresses + event signatures + positional topics, so identical filters from
  different workflows share one `resource_id`.
- **cron / http / syncer** — no shared physical resource exists, so it is the
  workflow-scoped `trigger_id` / `workflow_id`.

`ResourceIdentity` is the **sole** identity of a metered resource: `Utilization`
carries no labels. For workflow-scoped resources the workflow is recoverable
from `resource_id` (it is the `trigger_id` / `workflow_id`) and the owner is
resolved downstream from the workflow; shared resources (EVM filters) have no
single owner and are billed by the DON / node dimensions.

## MeterRecord idempotency key contract (level-triggered)

`Utilization.idempotency_key` is the lowercase hex SHA-256 over the full
structured identity (**including `node_id`**) plus the action and
event-identity:

```
product|environment|zone|don_id|node_id|service|resource|resource_type|action|resource_id|event-identity
```

where `action` is the `MeterAction` enum name (e.g. `METER_ACTION_RESERVE`).
Producers derive keys exclusively through the canonical helper
(`resourcemanager.IdempotencyKey` in chainlink-common); inputs are identifiers
and must not contain `|`.

Because `node_id` is in the preimage, keys are **unique per node**: billing
dedups a single node's retries by key and counts distinct nodes for quorum.
Cross-node grouping / convergence is the consumer's job on `resource_id` +
dimensions, which is independent of the key.

Keys are **level-triggered by design**: producers intentionally re-emit a
record with an identical key whenever they re-observe the same resource
lifecycle edge — for example re-registering all active triggers after a
restart. An identical key therefore means "the same logical event", not "a
producer bug".

`RESERVE` / `RELEASE` are emitted **only for genuine allocation and
deallocation** (register / unregister, workflow create / delete / pause).
Producers never synthesize a RELEASE for process-lifecycle cleanup of a
leaked or orphaned resource. A reservation lost without a RELEASE — e.g. a node
crash, or a log-poller filter orphaned by a failed unregister — is reconciled
by the resource's **absence from subsequent `MeterSnapshot`s** (see the
MeterSnapshot contract below), not by a special cleanup record. Consequently a
RELEASE always carries the same `value` as its paired RESERVE.

Consumer rules:

1. Use `idempotency_key` for **exact-duplicate suppression only** (per node).
   Records with equal keys are one logical event; bill it once.
2. Do NOT infer lifecycle state from key (re)appearance. Derive lifecycle
   state by ordering records by `timestamp` per `resource_id` + dimensions
   (last-write-wins), and treat `MeterSnapshot` as the authoritative liveness
   signal — a resource absent from snapshots is no longer active.
3. Pair `METER_ACTION_RESERVE` / `METER_ACTION_RELEASE` records by their
   `resource_id` + dimensions.

## MeterSnapshot contract

`MeterSnapshot` (in `snapshot.proto`) is the **action-less** periodic
utilization of **exactly one** active resource on one node. It exists because
pure RESERVE/RELEASE state transitions have no liveness signal — a node panic
would otherwise leak a reservation forever — and because periodic utilization is
the magnitude the billing median-across-nodes reducer consumes. `MeterAction`
does **not** apply to snapshots.

- `identity` is the **full** identity of the one resource (the six coarse
  dimensions plus `resource` / `resource_type` / `resource_id`).
- `utilization.value` is the resource's current level; the per-interval
  increment is `value` over `interval`.
- `interval` is the nominal period the snapshot covers, for staleness detection
  and for computing the increment.
- There is **no batch and no sequence**: the resource manager emits one
  `MeterSnapshot` per active resource per interval. A resource that **stops
  being snapshotted is no longer active** — billing zeroes it out by that
  absence rather than from an explicit empty record.

### MeterSnapshot idempotency key

Snapshots key per interval rather than per lifecycle edge. The key is the
lowercase hex SHA-256 of:

```
snapshot|product|environment|zone|don_id|node_id|service|resource|resource_id|interval-bucket
```

where `interval-bucket` is the snapshot timestamp truncated to `interval`. The
bucket makes each interval's snapshot distinct (so per-interval increments are
not collapsed) while deduping retries of the same interval; consumers aggregate
across nodes by `resource_id` + dimensions.

## Code generation

```bash
cd metering && make generate
# or, from the repo root:
task proto:gen:metering
```
