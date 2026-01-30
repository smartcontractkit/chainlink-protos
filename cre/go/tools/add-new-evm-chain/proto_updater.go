package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Sentinel errors for common error conditions.
var (
	ErrChainExists       = errors.New("chain already exists in proto file")
	ErrNoDefaultsArray   = errors.New("could not find defaults array in proto file")
	ErrNoEntriesFound    = errors.New("no entries found in defaults array")
	ErrPatternNotMatched = errors.New("failed to replace defaults array - pattern may not match")
)

// Pre-compiled regex patterns for better performance.
// See: 100 Go Mistakes #42 - Not using compiled regular expressions.
var (
	// defaultsArrayRe matches the entire defaults array block.
	defaultsArrayRe = regexp.MustCompile(`defaults:\s*\[[\s\S]*?\n\s*\]`)

	// defaultsContentRe extracts the content inside the defaults array.
	defaultsContentRe = regexp.MustCompile(`defaults:\s*\[([\s\S]*?)\n\s*\]`)

	// entryRe matches individual chain entries within the defaults array.
	entryRe = regexp.MustCompile(`\{\s*key:\s*"([^"]+)"\s*value:\s*(\d+)\s*\}`)
)

// ChainEntry represents a key-value entry in the proto file.
type ChainEntry struct {
	Key   string
	Value uint64
}

// chainSelectorExists checks if a chain selector already exists in the proto file.
func chainSelectorExists(protoFile, chainName string) (bool, error) {
	fileContent, err := os.ReadFile(protoFile)
	if err != nil {
		return false, fmt.Errorf("failed to read proto file: %w", err)
	}

	// Build pattern to check if the chain name already exists
	pattern := regexp.MustCompile(fmt.Sprintf(`key:\s*"%s"`, regexp.QuoteMeta(chainName)))
	return pattern.MatchString(string(fileContent)), nil
}

// updateProtoFile updates the proto file with a new chain selector entry.
func updateProtoFile(protoFile, chainName string, selector uint64) error {
	fileContent, err := os.ReadFile(protoFile)
	if err != nil {
		return fmt.Errorf("failed to read proto file: %w", err)
	}

	protoContent := string(fileContent)

	// Parse existing entries
	entries, err := parseProtoEntries(protoContent)
	if err != nil {
		return fmt.Errorf("failed to parse proto entries: %w", err)
	}

	// Check if entry already exists
	for _, entry := range entries {
		if entry.Key == chainName {
			return fmt.Errorf("%w: %s", ErrChainExists, chainName)
		}
	}

	// Add new entry
	entries = append(entries, ChainEntry{
		Key:   chainName,
		Value: selector,
	})

	// Sort entries alphabetically by key
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Key < entries[j].Key
	})

	// Rebuild the defaults array
	newDefaults := buildDefaultsArray(entries)

	// Replace the defaults array in the content
	newContent := defaultsArrayRe.ReplaceAllString(protoContent, newDefaults)
	if newContent == protoContent {
		return ErrPatternNotMatched
	}

	// Write back to file
	if err := os.WriteFile(protoFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write proto file: %w", err)
	}

	return nil
}

// parseProtoEntries parses the defaults array from the proto file.
func parseProtoEntries(protoContent string) ([]ChainEntry, error) {
	// Find the defaults array content
	matches := defaultsContentRe.FindStringSubmatch(protoContent)
	if len(matches) < 2 {
		return nil, ErrNoDefaultsArray
	}

	defaultsContent := matches[1]

	// Parse each entry
	entryMatches := entryRe.FindAllStringSubmatch(defaultsContent, -1)
	if len(entryMatches) == 0 {
		return nil, ErrNoEntriesFound
	}

	entries := make([]ChainEntry, 0, len(entryMatches))
	for _, match := range entryMatches {
		if len(match) < 3 {
			continue
		}

		key := match[1]
		value, err := strconv.ParseUint(match[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse selector value %s for key %s: %w", match[2], key, err)
		}

		entries = append(entries, ChainEntry{
			Key:   key,
			Value: value,
		})
	}

	return entries, nil
}

// buildDefaultsArray builds the defaults array string from entries.
func buildDefaultsArray(entries []ChainEntry) string {
	var builder strings.Builder
	builder.WriteString("defaults: [\n")

	for i, entry := range entries {
		builder.WriteString(fmt.Sprintf("            {\n              key: \"%s\"\n              value: %d\n            }", entry.Key, entry.Value))
		if i < len(entries)-1 {
			builder.WriteString(",\n")
		} else {
			builder.WriteString("\n")
		}
	}

	builder.WriteString("          ]")
	return builder.String()
}
