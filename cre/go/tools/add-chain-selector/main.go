// Package main provides a tool for adding new chain selectors to the chainlink-protos EVM client.proto file.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

// Config holds the tool configuration from command-line flags.
type Config struct {
	ChainSelector uint64
	DryRun        bool
	GithubToken   string
	BaseBranch    string
	ProtoFile     string
	RepoOwner     string
	RepoName      string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := parseFlags()
	if err != nil {
		return err
	}

	// Look up chain information from chain-selectors package
	chainName, err := lookupChainSelector(config.ChainSelector)
	if err != nil {
		return fmt.Errorf("failed to look up chain selector: %w", err)
	}

	fmt.Printf("Found chain: %s (selector: %d)\n", chainName, config.ChainSelector)

	// Check if chain already exists in proto file
	exists, err := chainSelectorExists(config.ProtoFile, chainName)
	if err != nil {
		return fmt.Errorf("failed to check if chain exists: %w", err)
	}

	if exists {
		fmt.Printf("Chain %s already exists in proto file, nothing to do\n", chainName)
		return nil
	}

	if config.DryRun {
		fmt.Println("\n=== DRY RUN MODE ===")
		fmt.Printf("Would add chain entry:\n")
		fmt.Printf("  Key: %s\n", chainName)
		fmt.Printf("  Value: %d\n", config.ChainSelector)
		fmt.Printf("To file: %s\n", config.ProtoFile)
		return nil
	}

	// Update proto file with new chain
	if err := updateProtoFile(config.ProtoFile, chainName, config.ChainSelector); err != nil {
		return fmt.Errorf("failed to update proto file: %w", err)
	}

	fmt.Printf("Successfully added chain %s to proto file\n", chainName)

	// Create GitHub PR
	if config.GithubToken == "" {
		fmt.Println("No GitHub token provided, skipping PR creation")
		fmt.Println("Please commit and push changes manually")
		return nil
	}

	ctx := context.Background()
	prURL, err := createPR(ctx, PRConfig{
		Token:      config.GithubToken,
		Owner:      config.RepoOwner,
		Repo:       config.RepoName,
		BaseBranch: config.BaseBranch,
		ChainName:  chainName,
		Selector:   config.ChainSelector,
		ProtoFile:  config.ProtoFile,
	})
	if err != nil {
		return fmt.Errorf("failed to create PR: %w", err)
	}

	fmt.Printf("Created PR: %s\n", prURL)
	return nil
}

func parseFlags() (*Config, error) {
	var chainSelector uint64
	var dryRun bool
	var githubToken string
	var baseBranch string
	var protoFile string
	var repoOwner string
	var repoName string

	flag.Uint64Var(&chainSelector, "chain-selector", 0, "The chain selector value (required)")
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be done without making changes")
	flag.StringVar(&githubToken, "github-token", "", "GitHub token for PR creation (required if not dry-run)")
	flag.StringVar(&baseBranch, "base-branch", "main", "Base branch for PR")
	flag.StringVar(&protoFile, "proto-file", "", "Path to client.proto (default: auto-detect)")
	flag.StringVar(&repoOwner, "repo-owner", "smartcontractkit", "GitHub repository owner")
	flag.StringVar(&repoName, "repo-name", "chainlink-protos", "GitHub repository name")

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
		GithubToken:   githubToken,
		BaseBranch:    baseBranch,
		ProtoFile:     protoFile,
		RepoOwner:     repoOwner,
		RepoName:      repoName,
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

