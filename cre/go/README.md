Run `make` to generate all the necessary files for the Go SDK.

Ideally, all genreated files would live in their SDKs repository (or a shared one like chainlink-common).
The basic values, sdk, and tools are installed here in Go validation on protos CRE metadata.

Furthermore, since the value wrappers should never drift from the values protos, and are simply thin wrappers, it make sense to have them here.

Anything beyond that should be in the SDKs repository, and not in chainlink protos.
