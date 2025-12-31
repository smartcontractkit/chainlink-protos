// Package main provides a tool for adding new chain selectors to the chainlink-protos EVM client.proto file.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

// Action represents the outcome action of the tool execution.
type Action string

// Action constants define the possible outcomes of the tool execution.
const (
	ActionAdded  Action = "added"
	ActionExists Action = "exists"
	ActionDryRun Action = "dry-run"
	ActionError  Action = "error"
)

// Config holds the tool configuration from command-line flags.
type Config struct {
	ChainSelector uint64
	DryRun        bool
	ProtoFile     string
}

// Result represents the outcome of the tool execution.
type Result struct {
	ChainName     string
	ChainSelector uint64
	Action        Action
}

func main() {
	result, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// Output structured result for workflow
		fmt.Println("\n--- OUTPUT ---")
		fmt.Printf("ACTION=%s\n", ActionError)
		os.Exit(1)
	}

	// Output structured result for workflow consumption
	fmt.Println("\n--- OUTPUT ---")
	fmt.Printf("CHAIN_NAME=%s\n", result.ChainName)
	fmt.Printf("CHAIN_SELECTOR=%d\n", result.ChainSelector)
	fmt.Printf("ACTION=%s\n", result.Action)
}

func run() (*Result, error) {
	config, err := parseFlags()
	if err != nil {
		return nil, err
	}

	// Look up chain information from chain-selectors package
	chainName, err := lookupChainSelector(config.ChainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to look up chain selector: %w", err)
	}

	fmt.Printf("Found chain: %s (selector: %d)\n", chainName, config.ChainSelector)

	// Check if chain already exists in proto file
	exists, err := chainSelectorExists(config.ProtoFile, chainName)
	if err != nil {
		return nil, fmt.Errorf("failed to check if chain exists: %w", err)
	}

	if exists {
		fmt.Printf("Chain %s already exists in proto file, nothing to do\n", chainName)
		return &Result{
			ChainName:     chainName,
			ChainSelector: config.ChainSelector,
			Action:        ActionExists,
		}, nil
	}

	if config.DryRun {
		fmt.Println("\n=== DRY RUN MODE ===")
		fmt.Printf("Would add chain entry:\n")
		fmt.Printf("  Key: %s\n", chainName)
		fmt.Printf("  Value: %d\n", config.ChainSelector)
		fmt.Printf("To file: %s\n", config.ProtoFile)
		return &Result{
			ChainName:     chainName,
			ChainSelector: config.ChainSelector,
			Action:        ActionDryRun,
		}, nil
	}

	// Update proto file with new chain
	if err := updateProtoFile(config.ProtoFile, chainName, config.ChainSelector); err != nil {
		return nil, fmt.Errorf("failed to update proto file: %w", err)
	}

	fmt.Printf("Successfully added chain %s to proto file\n", chainName)

	return &Result{
		ChainName:     chainName,
		ChainSelector: config.ChainSelector,
		Action:        ActionAdded,
	}, nil
}

func parseFlags() (*Config, error) {
	var chainSelector uint64
	var dryRun bool
	var protoFile string

	flag.Uint64Var(&chainSelector, "chain-selector", 0, "The chain selector value (required)")
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be done without making changes")
	flag.StringVar(&protoFile, "proto-file", "", "Path to client.proto (default: auto-detect)")

	flag.Parse()

	if chainSelector == 0 {
		return nil, fmt.Errorf("chain-selector is required")
	}

	// Auto-detect proto file path if not provided
	if protoFile == "" {
		protoFile = findProtoFile()
		if protoFile == "" {
			return nil, fmt.Errorf("could not auto-detect proto file path, please provide -proto-file flag")
		}
	}

	// Validate proto file exists
	if _, err := os.Stat(protoFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("proto file not found: %s", protoFile)
	}

	return &Config{
		ChainSelector: chainSelector,
		DryRun:        dryRun,
		ProtoFile:     protoFile,
	}, nil
}

// findProtoFile attempts to find the client.proto file relative to the tool location.
func findProtoFile() string {
	// Try relative paths from likely locations
	candidates := []string{
		"cre/capabilities/blockchain/evm/v1alpha/client.proto",
		"../../../capabilities/blockchain/evm/v1alpha/client.proto",
		"../../../../cre/capabilities/blockchain/evm/v1alpha/client.proto",
	}

	// First try from current working directory
	cwd, err := os.Getwd()
	if err == nil {
		for _, candidate := range candidates {
			path := filepath.Join(cwd, candidate)
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	// Try to find it relative to the executable
	exe, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exe)
		for _, candidate := range candidates {
			path := filepath.Join(exeDir, candidate)
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	return ""
}

// lookupChainSelector looks up chain information from the chain-selectors package.
func lookupChainSelector(selector uint64) (string, error) {
	// Verify it's an EVM chain
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return "", fmt.Errorf("chain selector %d not found: %w", selector, err)
	}

	if family != chain_selectors.FamilyEVM {
		return "", fmt.Errorf("chain selector %d is not an EVM chain (family: %s), only EVM chains are supported", selector, family)
	}

	// Get chain name
	chainName, err := chain_selectors.GetChainNameFromSelector(selector)
	if err != nil {
		return "", fmt.Errorf("failed to get chain name for selector %d: %w", selector, err)
	}

	return chainName, nil
}
