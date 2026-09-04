---
"@chainlink/node-platform": minor
---

Add `common.v1.CLJobInfo`: a job-type-agnostic snapshot of a single Chainlink job, carrying
the job's common identity fields, the complete definition as a raw TOML string, and
optional Job Distributor provenance (`feeds_manager_id`, `remote_uuid`, `spec_version`,
`proposed_at`, `approved_at`). Lets any job type be reported through one schema without a
dedicated message per type.
