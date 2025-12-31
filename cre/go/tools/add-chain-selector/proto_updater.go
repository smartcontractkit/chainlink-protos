package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ChainEntry represents a key-value entry in the proto file.
type ChainEntry struct {
	Key   string
	Value uint64
}

// chainSelectorExists checks if a chain selector already exists in the proto file.
func chainSelectorExists(protoFile, chainName string) (bool, error) {
	data, err := os.ReadFile(protoFile)
	if err != nil {
		return false, fmt.Errorf("failed to read proto file: %w", err)
	}

	content := string(data)
	// Check if the chain name already exists in the defaults array
	pattern := fmt.Sprintf(`key:\s*"%s"`, regexp.QuoteMeta(chainName))
	matched, err := regexp.MatchString(pattern, content)
	if err != nil {
		return false, fmt.Errorf("failed to match pattern: %w", err)
	}

	return matched, nil
}

// updateProtoFile updates the proto file with a new chain selector entry.
func updateProtoFile(protoFile, chainName string, selector uint64) error {
	data, err := os.ReadFile(protoFile)
	if err != nil {
		return fmt.Errorf("failed to read proto file: %w", err)
	}

	content := string(data)

	// Parse existing entries
	entries, err := parseProtoEntries(content)
	if err != nil {
		return fmt.Errorf("failed to parse proto entries: %w", err)
	}

	// Check if entry already exists
	for _, entry := range entries {
		if entry.Key == chainName {
			return fmt.Errorf("chain %s already exists in proto file", chainName)
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
	re := regexp.MustCompile(`defaults:\s*\[[\s\S]*?\n\s*\]`)
	newContent := re.ReplaceAllString(content, newDefaults)
	if newContent == content {
		return fmt.Errorf("failed to replace defaults array - pattern may not match")
	}

	// Write back to file
	if err := os.WriteFile(protoFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write proto file: %w", err)
	}

	return nil
}

// parseProtoEntries parses the defaults array from the proto file.
func parseProtoEntries(content string) ([]ChainEntry, error) {
	entries := []ChainEntry{}

	// Find the defaults array - match from "defaults: [" to the closing "]"
	re := regexp.MustCompile(`defaults:\s*\[([\s\S]*?)\n\s*\]`)
	matches := re.FindStringSubmatch(content)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not find defaults array in proto file")
	}

	defaultsContent := matches[1]

	// Parse each entry - match the exact format with proper whitespace
	// Pattern: { key: "..." value: ... }
	entryRe := regexp.MustCompile(`\{\s*key:\s*"([^"]+)"\s*value:\s*(\d+)\s*\}`)
	entryMatches := entryRe.FindAllStringSubmatch(defaultsContent, -1)

	if len(entryMatches) == 0 {
		return nil, fmt.Errorf("no entries found in defaults array")
	}

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

