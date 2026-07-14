# @chainlink/svr

## 1.3.0

### Minor Changes

- [#429](https://github.com/smartcontractkit/chainlink-protos/pull/429) [`8c19365`](https://github.com/smartcontractkit/chainlink-protos/commit/8c19365a8f8ea315f3205e290524481fa0512183) Thanks [@mostlyconsistent](https://github.com/mostlyconsistent)! - Add optional `gas_limit` field (9) to `TxMessage` proto. Populated with the per-OFA "accountability" gas limit (the OFA-tiered gas limit) so on-chain transactions can be attributed to a specific OFA (e.g. Titan/MEV-Share) for auction and performance comparison.

## 1.2.0

### Minor Changes

- OEV-851: Add optional `dual_broadcast_params` field (8) to `TxMessage` proto. Populated with the URL-encoded MEVShare/Atlas params when a secondary (dual-broadcast) transaction is emitted.

## 1.1.0

### Minor Changes

- [#60](https://github.com/smartcontractkit/chainlink-protos/pull/60) [`5063bca`](https://github.com/smartcontractkit/chainlink-protos/commit/5063bca287485b7ad6db05b87a4f2731ac514a48) Thanks [@cll-gg](https://github.com/cll-gg)! - Add information about chain + feed to SVR TxMessage

## 1.0.0

### Major Changes

- [#54](https://github.com/smartcontractkit/chainlink-protos/pull/54) [`26a79f9`](https://github.com/smartcontractkit/chainlink-protos/commit/26a79f9ba4bd44a6f0f9138c193f52c87cc7e1aa) Thanks [@eduard-cl](https://github.com/eduard-cl)! - SVR package first release, includes v1 tx proto message.
