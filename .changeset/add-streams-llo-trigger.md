---
"@smartcontractkit/chainlink-protos": minor
---

Add Streams LLO NoDAG Trigger capability proto

- Add streams/v1/trigger.proto with NoDAG API
- Config message (stream_ids, max_frequency_ms)
- Report message matching OCRTriggerEvent structure
- Streams service with Trigger RPC (streams-trigger@2.0.0)
- Supports Data Feeds migration to CRE NoDAG architecture
