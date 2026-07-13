---
"@chainlink/svr": minor
---

Add optional `gas_limit` field (9) to `TxMessage` proto. Populated with the per-OFA "accountability" gas limit (the OFA-tiered gas limit) so on-chain transactions can be attributed to a specific OFA (e.g. Titan/MEV-Share) for auction and performance comparison.
