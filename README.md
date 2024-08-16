# tooling-protos

This repo holds the protobuf definitions for the `job-distributor` and `orchestrator` services.

## dependencies

Dependencies are managed via [asdf](https://asdf-vm.com/guide/getting-started.html).

```bash
$ asdf install
```

## generating GO SDKs

Generate the GO SDKs to implement gRPC services or clients via [task](https://taskfile.dev/installation/)

```bash
$ task generate
```
