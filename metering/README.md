# Metering

Protobuf definitions for `MeterRecord`, the billable event emitted by services
that manage durable resources (trigger registrations, log filters, workflow
registrations, ...). Records are published via Beholder and consumed by the
billing pipeline.

Derived Kafka subject: `beholder__platform__messages-metering.v1.MeterRecord`
(domain `beholder__platform__messages` from `chip-platform.json` + schema
entity `metering.v1.MeterRecord`).

## Entity naming

`MeterRecord.entity` is a stable service constant identifying the emitting
service, e.g. `"cron-trigger"`, `"http-trigger"`, `"evm-log-trigger"`,
`"workflow-syncer-v2"`. Deployment environment and zone are deliberately NOT
encoded in the entity: they ride on transport-level attributes (OTel resource
attributes attached by the LOOP server) and on the per-environment CHiP hosts
and pipelines.

## Idempotency key contract (level-triggered)

`Utilization.idempotency_key` is the lowercase hex SHA-256 of

```
entity|resource|action|resource-id|event-identity
```

where `action` is the `MeterAction` enum name (e.g. `METER_ACTION_RESERVE`).
Producers derive keys exclusively through the canonical helper
(`resourcemanager.IdempotencyKey` in chainlink-common); inputs are identifiers
and must not contain `|`.

Keys are **level-triggered by design**: producers intentionally re-emit a
record with an identical key whenever they re-observe the same resource
lifecycle edge — for example re-registering all active triggers after a
restart, or a cleanup pass that overlaps an explicit release. An identical key
therefore means "the same logical event", not "a producer bug".

Consumer rules:

1. Use `idempotency_key` for **exact-duplicate suppression only**. Records
   with equal keys are one logical event; bill it once.
2. Do NOT infer lifecycle state from key (re)appearance. Derive lifecycle
   state by ordering records by `timestamp` per resource-id label
   (last-write-wins).
3. Pair `METER_ACTION_RESERVE` / `METER_ACTION_RELEASE` records by their
   resource-id labels (`trigger_id` / `workflow_id` / `filter_id`), not by
   key or value.
4. A RELEASE's `value` may differ from its paired RESERVE's. In particular,
   the EVM log-trigger's orphan cleanup emits RELEASE records with `value` 0;
   the pairing labels, not the value, identify which reservation ended.

## Code generation

```bash
cd metering && make generate
# or, from the repo root:
task proto:gen:metering
```
