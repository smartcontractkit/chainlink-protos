# chainlink-protos

> **Note**
>
> _This demo represents an educational example to use a Chainlink system, product, or service and is provided to demonstrate how to interact with Chainlink’s systems, products, and services to integrate them into your own. This template is provided “AS IS” and “AS AVAILABLE” without warranties of any kind, it has not been audited, and it may be missing key checks or error handling to make the usage of the system, product or service more clear. Do not use the code in this example in a production environment without completing your own audits and application of best practices. Neither Chainlink Labs, the Chainlink Foundation, nor Chainlink node operators are responsible for unintended outputs that are generated due to errors in code._

This repository serves as a central hub for shared protobuf definitions used across various services. Currently, it includes protobuf definitions and generated Go SDKs for the job-distributor and orchestrator services. Contributions and additional proto files for other services are welcome.

## Usage

Go applications can import only the necessary SDKs.

```bash
$ go get github.com/smartcontractkit/chainlink-protos/job-distributor
$ go get github.com/smartcontractkit/chainlink-protos/orchestrator
```

## Dependencies

Dependencies are managed via [asdf](https://asdf-vm.com/guide/getting-started.html).

```bash
$ asdf install
```

### Installing wsRPC

Communication between core node and job distributor requires the library [wsRPC](https://github.com/smartcontractkit/wsrpc). To generate protos that are compatible with wsRPC, we will need to install the CLI.

Follow the instructions [here](https://github.com/smartcontractkit/wsrpc?tab=readme-ov-file#set-up) to install it.

## Generating GO SDKs

Generate the GO SDKs to implement gRPC services or clients via [task](https://taskfile.dev/installation/)

```bash
$ task proto:all
$ task proto:gen:job-distributor # only run for job-distributor
$ task proto:gen:orchestrator: # only run for orchestrator
```

## Contributing

### Filing a PR

Before creating a PR with your change, you should generate a "changeset" file.

Let's assume that you've made some local changes in one of the protos.
Before filing a PR you need to generate a "changeset" description required for
the automated release process. Follow the steps below:

- Run `pnpm changset` in the git top level directory.
- This repo contains multiple packages, so it will ask you for which package it
  should generate changeset update.
- Answer remaining questions. At the end, you will have a new
  `.changeset/<random-name>.md` file generated.
- Now you need to commit and push your changes
- Create a Pull request which includes your code change and generated
  "changeset" file.

#### Preparing a release

After merging your PR, a changesets CI job will create or update a "Version
Packages" PR which contains a release bump.

#### Merging Version Packages PR

Now you can Approve/Request approval and Merge the PR from the previous step.
After merging, it will kick off the push-main.yml workflow and that will release
a new versions and push tags automatically. You can navigate to the
[tags view](https://github.com/smartcontractkit/chainlink-protos/tags), to check if the
latest tag is available.
