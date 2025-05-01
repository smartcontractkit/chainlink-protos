<div align="center">
  <h1>Chainlink Protos</h1>
  <a><img src="https://github.com/smartcontractkit/chainlink-protos/actions/workflows/push-main.yml/badge.svg" /></a>
  <br/>
  <br/>
</div>

This repository serves as a central hub for shared protobuf definitions used across various services. Currently, it includes protobuf definitions and generated Go SDKs for the job-distributor and orchestrator services. Contributions and additional proto files for other services are welcome.

## Why use Chainlink Protos?

- **No More Copy-Pasting Code**: Centralized proto definitions eliminate the need to copy-paste code across different repositories.
- **Publicly Accessible Repository**: No need for tokens, credentials, or GOPRIVATE settings.
- **Proto Breaking Changes Validation with Buf**: Ensures backward compatibility and consistent formatting.
- **Automated Golang SDK Generation**: GitHub workflows keep the generated code up-to-date.
- **Independent Versioning in a Monorepo**: Each package is versioned separately for better control and integration.

## Usage

Go applications can import only the necessary SDKs.

```bash
$ go get github.com/smartcontractkit/chainlink-protos/job-distributor@v<LATEST_VERSION>
$ go get github.com/smartcontractkit/chainlink-protos/orchestrator@v<LATEST_VERSION>
```

### Import

The import varies depending on the `go_package` option defined in the protos.
Below is an example when `go_package` is set to `github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node`

```go
import "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
```

## Development

### Getting Started

#### Setup

[asdf](https://asdf-vm.com/) is a tool version manager. All dependencies used for local development of this repo are
managed through `asdf`. To install `asdf`:

1. [Install asdf](https://asdf-vm.com/guide/getting-started.html)
2. Follow the instructions to ensure `asdf` is shimmed into your terminal or development environment

#### Installing Dependencies

Install the required tools using [asdf](https://asdf-vm.com/guide/getting-started.html):

```bash
./scripts/setup-asdf-plugin.sh
asdf install
```

#### Installing wsRPC

Communication between core node and job distributor requires the library [wsRPC](https://github.com/smartcontractkit/wsrpc). To generate protos that are compatible with wsRPC, we will need to install the CLI.

Follow the instructions [here](https://github.com/smartcontractkit/wsrpc?tab=readme-ov-file#set-up) to install it.

### Formatting

Ensure [buf](https://buf.build/product/cli) is installed following the dependencies above.

```bash
task fmt
```

### Linting

```bash
task lint
```

### Generating GO SDKs

> [!Note]
> Commiting the generated code resulting from the proto changes is optional. The CI will automatically update the pull request with the generated files through the GitHub workflow.

Generate the GO SDKs to implement gRPC services or clients via [task](https://taskfile.dev/installation/)

```bash
$ task proto:all
$ task proto:gen:job-distributor # only run for job-distributor
$ task proto:gen:orchestrator: # only run for orchestrator
$ task proto:gen:svr: # only run for svr
$ task proto:gen:rmn
$ task proto:gen:workflows
```

## Contributing

For instructions on how to contribute to `chainlink-protos` and the release process,
see [CONTRIBUTING.md](https://github.com/smartcontractkit/chainlink-protos/blob/main/CONTRIBUTING.md)

## Releasing

For instructions on how to release `chainlink-protos`,
see [RELEASE.md](https://github.com/smartcontractkit/chainlink-protos/blob/main/RELEASE.md)
