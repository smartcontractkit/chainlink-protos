---
"@chainlink/job-distributor": minor
"@chainlink/orchestrator": minor
---

Add `CHAIN_TYPE_STELLAR` (8) and `CHAIN_TYPE_CANTON` (9) to the `ChainType` enum in
`feedsmanager.proto` and `node.proto`. These chain families are needed for standalone
committee verifier nodes to publish their onchain signing keys to JD via `UpdateNode`,
unblocking deployment tooling that reads signing addresses back via
`ListNodeChainConfigs`.
