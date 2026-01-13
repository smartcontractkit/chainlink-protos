---
"@smartcontractkit/chainlink-protos": minor
---

Add Streams LLO Trigger NoDAG API

Migrates the existing Streams LLO Trigger to the NoDAG capability API.
This replaces the legacy RegisterTrigger/UnregisterTrigger API with a
proto-based streaming RPC pattern.

- Add Streams Trigger capability proto (streams/v1/trigger.proto)
- Defines Config message for stream subscription configuration
- Defines Report message matching existing OCRTriggerEvent structure
- Support for Data Feeds migration to CRE DON-to-DON architecture
