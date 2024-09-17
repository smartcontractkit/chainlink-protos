# chainlink-protos

This repo holds the protobuf definitions and generated Go SDKs for the `job-distributor` and `orchestrator` services.

Go applications should depend on the necessary SDKs directly

```bash
$ go get github.com/smartcontractkit/chainlink-protos/job-distributor
$ go get github.com/smartcontractkit/chainlink-protos/orchestrator
```

Other applications may build their own SDKs directly from the provided protobufs.

## Dependencies

Dependencies are managed via [asdf](https://asdf-vm.com/guide/getting-started.html).

```bash
$ asdf install
```

### Installing wsRPC

Communication between core node and job distributor requires the library [wsRPC](https://github.com/smartcontractkit/wsrpc). To generate protos that are compatible with wsRPC, we will need to install the CLI.

Follow the instructions [here](https://github.com/smartcontractkit/wsrpc?tab=readme-ov-file#set-up) to install it.

## Generating GO SDKs Manually

Generate the GO SDKs to implement gRPC services or clients via [task](https://taskfile.dev/installation/)

```bash
$ task proto:all
$ task proto:gen:job-distributor # only run for job-distributor
```
