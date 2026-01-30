# Add New EVM Chain Tool

A CLI tool for adding new EVM chain selectors to the chainlink-protos EVM client.proto file.

## Overview

This tool automates the process of adding a new EVM chain to the CRE (Chainlink Runtime Environment) by:

1. Looking up chain information from the [chain-selectors](https://github.com/smartcontractkit/chain-selectors) package
2. Validating that the chain is an EVM chain
3. Checking if the chain already exists in the proto file
4. Adding the chain entry to `cre/capabilities/blockchain/evm/v1alpha/client.proto` (sorted alphabetically)

## Installation

Build the tool from the `cre/go` directory:

```bash
cd cre/go
go build -o add-new-evm-chain ./tools/add-new-evm-chain
```

## Usage

### Basic Usage

```bash
# From the cre/go directory
./add-new-evm-chain -chain-selector <SELECTOR_VALUE>
```

### Flags

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `-chain-selector` | Yes | - | The chain selector value (uint64) |
| `-proto-file` | No | Auto-detect | Path to client.proto file |
| `-dry-run` | No | `false` | Preview changes without modifying files |

### Examples

#### Dry Run (Preview Changes)

```bash
./add-new-evm-chain -chain-selector 5009297550715157269 -dry-run
```

Output:
```
Found chain: ethereum-mainnet (selector: 5009297550715157269)

=== DRY RUN MODE ===
Would add chain entry:
  Key: ethereum-mainnet
  Value: 5009297550715157269
To file: ../capabilities/blockchain/evm/v1alpha/client.proto

--- OUTPUT ---
CHAIN_NAME=ethereum-mainnet
CHAIN_SELECTOR=5009297550715157269
ACTION=dry-run
```

#### Add a New Chain

```bash
./add-new-evm-chain -chain-selector 5009297550715157269
```

Output:
```
Found chain: ethereum-mainnet (selector: 5009297550715157269)
Successfully added chain ethereum-mainnet to proto file

--- OUTPUT ---
CHAIN_NAME=ethereum-mainnet
CHAIN_SELECTOR=5009297550715157269
ACTION=added
```

#### Existing Chain

If the chain already exists:
```
Found chain: ethereum-mainnet (selector: 5009297550715157269)
Chain ethereum-mainnet already exists in proto file, nothing to do

--- OUTPUT ---
CHAIN_NAME=ethereum-mainnet
CHAIN_SELECTOR=5009297550715157269
ACTION=exists
```

## Output Format

The tool outputs structured data for workflow consumption:

```
--- OUTPUT ---
CHAIN_NAME=<chain-name>
CHAIN_SELECTOR=<selector-value>
ACTION=<added|exists|dry-run|error>
```

| Action | Description |
|--------|-------------|
| `added` | Chain was successfully added to the proto file |
| `exists` | Chain already exists in the proto file |
| `dry-run` | Dry run mode - no changes made |
| `error` | An error occurred |

## Finding Chain Selectors

Chain selectors can be found in the [chain-selectors repository](https://github.com/smartcontractkit/chain-selectors):

- EVM chains: `selectors.yml`
- Look for the `selector` field for each chain

Example from `selectors.yml`:
```yaml
selectors:
  1:  # Ethereum Mainnet chain ID
    selector: 5009297550715157269
    name: "ethereum-mainnet"
```

## Running Tests

```bash
cd cre/go
go test ./tools/add-new-evm-chain/... -v
```

## GitHub Workflow

This tool is also available as a GitHub Actions workflow that can be triggered manually:

1. Go to **Actions** → **Add Chain Selector**
2. Click **Run workflow**
3. Enter the chain selector value
4. Optionally enable dry-run mode
5. Click **Run workflow**

The workflow will:
- Run this tool to update the proto file
- Create a branch with the changes
- Open a Pull Request for review

## Proto File Format

The tool adds entries to the `defaults` array in `client.proto`:

```protobuf
defaults: [
  {
    key: "avalanche-mainnet"
    value: 6433500567565415381
  },
  {
    key: "ethereum-mainnet"
    value: 5009297550715157269
  }
]
```

Entries are automatically sorted alphabetically by chain name.

## Code Practices & References

This tool follows Go best practices and idioms from the following resources:

- [Ardan Labs / Bill Kennedy's Ultimate Go Training](https://www.ardanlabs.com/training/ultimate-go/)
- [Go in Action](https://www.manning.com/books/go-in-action) (Manning Publications)
- [The Go Blog](https://go.dev/blog/)
- [Effective Go](https://go.dev/doc/effective_go) (Official)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) (Official Wiki)
- [The Go Memory Model](https://go.dev/ref/mem) (Official)
- [Go Proverbs](https://go-proverbs.github.io/) (Rob Pike)
- [Google Go Style Guide](https://google.github.io/styleguide/go/) (Style Guide + Best Practices)
- [100 Go Mistakes and How to Avoid Them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) (Teiva Harsanyi)
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) (Katherine Cox-Buday)
