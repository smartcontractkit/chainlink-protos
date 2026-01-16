// Package main adds a new EVM chain selector to client.proto.
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

const protoPath = "../capabilities/blockchain/evm/v1alpha/client.proto"

func main() {
	selector := flag.Uint64("selector", 0, "chain selector value (required)")
	flag.Parse()

	if *selector == 0 {
		fatal("selector is required")
	}

	// Look up chain ID from selector
	chainId, err := chain_selectors.ChainIdFromSelector(*selector)
	if err != nil {
		fatal("selector %d not found: %v", *selector, err)
	}

	// Get chain name from chain ID
	chainName, err := chain_selectors.NameFromChainId(chainId)
	if err != nil {
		fatal("failed to get chain name for chain ID %d: %v", chainId, err)
	}

	// Read proto file
	content, err := os.ReadFile(protoPath)
	if err != nil {
		fatal("failed to read %s: %v", protoPath, err)
	}

	// Check if already exists
	if strings.Contains(string(content), fmt.Sprintf(`key: "%s"`, chainName)) {
		fmt.Printf("chain %s already exists\n", chainName)
		return
	}

	// Parse, add, sort, rebuild
	newContent, err := addChain(string(content), chainName, *selector)
	if err != nil {
		fatal("failed to add chain: %v", err)
	}

	if err := os.WriteFile(protoPath, []byte(newContent), 0644); err != nil {
		fatal("failed to write file: %v", err)
	}

	fmt.Printf("added %s (selector: %d)\n", chainName, *selector)
}

func addChain(content, name string, selector uint64) (string, error) {
	// Find defaults array
	re := regexp.MustCompile(`defaults:\s*\[([\s\S]*?)\n\s*\]`)
	match := re.FindStringSubmatch(content)
	if len(match) < 2 {
		return "", fmt.Errorf("defaults array not found")
	}

	// Parse entries
	entryRe := regexp.MustCompile(`\{\s*key:\s*"([^"]+)"\s*value:\s*(\d+)\s*\}`)
	entries := entryRe.FindAllStringSubmatch(match[1], -1)

	type entry struct {
		key string
		val uint64
	}
	var list []entry
	for _, e := range entries {
		v, _ := strconv.ParseUint(e[2], 10, 64)
		list = append(list, entry{e[1], v})
	}
	list = append(list, entry{name, selector})

	sort.Slice(list, func(i, j int) bool { return list[i].key < list[j].key })

	var b strings.Builder
	b.WriteString("defaults: [\n")
	for i, e := range list {
		b.WriteString(fmt.Sprintf("            {\n              key: \"%s\"\n              value: %d\n            }", e.key, e.val))
		if i < len(list)-1 {
			b.WriteString(",\n")
		} else {
			b.WriteString("\n")
		}
	}
	b.WriteString("          ]")

	return re.ReplaceAllString(content, b.String()), nil
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}

