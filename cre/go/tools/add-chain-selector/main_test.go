package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupChainSelector(t *testing.T) {
	tests := []struct {
		name        string
		selector    uint64
		wantName    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid ethereum mainnet selector",
			selector: 5009297550715157269,
			wantName: "ethereum-mainnet",
			wantErr:  false,
		},
		{
			name:     "valid avalanche mainnet selector",
			selector: 6433500567565415381,
			wantName: "avalanche-mainnet",
			wantErr:  false,
		},
		{
			name:     "valid polygon mainnet selector",
			selector: 4051577828743386545,
			wantName: "polygon-mainnet",
			wantErr:  false,
		},
		{
			name:        "invalid selector",
			selector:    999999999999,
			wantErr:     true,
			errContains: "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := lookupChainSelector(tt.selector)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantName, name)
		})
	}
}

func TestParseProtoEntries(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		wantEntries []ChainEntry
		wantErr     bool
		errContains string
	}{
		{
			name: "valid proto content",
			content: `
defaults: [
            {
              key: "ethereum-mainnet"
              value: 5009297550715157269
            },
            {
              key: "polygon-mainnet"
              value: 4051577828743386545
            }
          ]`,
			wantEntries: []ChainEntry{
				{Key: "ethereum-mainnet", Value: 5009297550715157269},
				{Key: "polygon-mainnet", Value: 4051577828743386545},
			},
			wantErr: false,
		},
		{
			name:        "missing defaults array",
			content:     `some other content`,
			wantErr:     true,
			errContains: "could not find defaults array",
		},
		{
			name: "empty defaults array",
			content: `
defaults: [
          ]`,
			wantErr:     true,
			errContains: "no entries found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := parseProtoEntries(tt.content)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantEntries, entries)
		})
	}
}

func TestBuildDefaultsArray(t *testing.T) {
	entries := []ChainEntry{
		{Key: "avalanche-mainnet", Value: 6433500567565415381},
		{Key: "ethereum-mainnet", Value: 5009297550715157269},
	}

	result := buildDefaultsArray(entries)

	// Check that it contains the expected structure
	assert.Contains(t, result, "defaults: [")
	assert.Contains(t, result, `key: "avalanche-mainnet"`)
	assert.Contains(t, result, `value: 6433500567565415381`)
	assert.Contains(t, result, `key: "ethereum-mainnet"`)
	assert.Contains(t, result, `value: 5009297550715157269`)
	assert.True(t, strings.HasSuffix(result, "]"))
}

func TestChainSelectorExists(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.proto")

	content := `
service Client {
  option (tools.generator.v1alpha.capability) = {
    labels: {
      key: "ChainSelector"
      value: {
        uint64_label: {
          defaults: [
            {
              key: "ethereum-mainnet"
              value: 5009297550715157269
            }
          ]
        }
      }
    }
  };
}`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	tests := []struct {
		name       string
		chainName  string
		wantExists bool
	}{
		{
			name:       "chain exists",
			chainName:  "ethereum-mainnet",
			wantExists: true,
		},
		{
			name:       "chain does not exist",
			chainName:  "polygon-mainnet",
			wantExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists, err := chainSelectorExists(testFile, tt.chainName)
			require.NoError(t, err)
			assert.Equal(t, tt.wantExists, exists)
		})
	}
}

func TestUpdateProtoFile(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.proto")

	originalContent := `syntax = "proto3";

service Client {
  option (tools.generator.v1alpha.capability) = {
    labels: {
      key: "ChainSelector"
      value: {
        uint64_label: {
          defaults: [
            {
              key: "ethereum-mainnet"
              value: 5009297550715157269
            }
          ]
        }
      }
    }
  };
}`

	t.Run("add new chain", func(t *testing.T) {
		err := os.WriteFile(testFile, []byte(originalContent), 0644)
		require.NoError(t, err)

		err = updateProtoFile(testFile, "avalanche-mainnet", 6433500567565415381)
		require.NoError(t, err)

		// Read the updated content
		updatedContent, err := os.ReadFile(testFile)
		require.NoError(t, err)

		// Verify the new chain was added
		assert.Contains(t, string(updatedContent), `key: "avalanche-mainnet"`)
		assert.Contains(t, string(updatedContent), `value: 6433500567565415381`)

		// Verify original chain still exists
		assert.Contains(t, string(updatedContent), `key: "ethereum-mainnet"`)

		// Verify alphabetical ordering (avalanche comes before ethereum)
		avaxIdx := strings.Index(string(updatedContent), "avalanche-mainnet")
		ethIdx := strings.Index(string(updatedContent), "ethereum-mainnet")
		assert.Less(t, avaxIdx, ethIdx, "avalanche should come before ethereum alphabetically")
	})

	t.Run("add chain alphabetically after existing", func(t *testing.T) {
		err := os.WriteFile(testFile, []byte(originalContent), 0644)
		require.NoError(t, err)

		err = updateProtoFile(testFile, "polygon-mainnet", 4051577828743386545)
		require.NoError(t, err)

		// Read the updated content
		updatedContent, err := os.ReadFile(testFile)
		require.NoError(t, err)

		// Verify the new chain was added
		assert.Contains(t, string(updatedContent), `key: "polygon-mainnet"`)

		// Verify ordering (ethereum comes before polygon)
		ethIdx := strings.Index(string(updatedContent), "ethereum-mainnet")
		polyIdx := strings.Index(string(updatedContent), "polygon-mainnet")
		assert.Less(t, ethIdx, polyIdx, "ethereum should come before polygon alphabetically")
	})

	t.Run("reject duplicate chain", func(t *testing.T) {
		err := os.WriteFile(testFile, []byte(originalContent), 0644)
		require.NoError(t, err)

		err = updateProtoFile(testFile, "ethereum-mainnet", 5009297550715157269)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "already exists")
	})
}

func TestSanitizeBranchName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple name",
			input: "ethereum-mainnet",
			want:  "ethereum-mainnet",
		},
		{
			name:  "name with underscores",
			input: "binance_smart_chain-mainnet",
			want:  "binance-smart-chain-mainnet",
		},
		{
			name:  "name with special chars",
			input: "chain@1.0.0",
			want:  "chain100",
		},
		{
			name:  "uppercase",
			input: "ETHEREUM-MAINNET",
			want:  "ethereum-mainnet",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sanitizeBranchName(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntegration_DryRun(t *testing.T) {
	// This test verifies the dry-run flow works correctly
	// by using a real proto file from testdata

	testdataDir := filepath.Join("testdata")
	testFile := filepath.Join(testdataDir, "test_client.proto")

	// Skip if testdata file doesn't exist
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skip("testdata/test_client.proto not found")
	}

	// Read original content to verify it's not modified
	originalContent, err := os.ReadFile(testFile)
	require.NoError(t, err)

	// Verify a chain we know exists
	exists, err := chainSelectorExists(testFile, "ethereum-mainnet")
	require.NoError(t, err)
	assert.True(t, exists, "ethereum-mainnet should exist in test file")

	// Verify a chain that doesn't exist
	exists, err = chainSelectorExists(testFile, "some-new-chain")
	require.NoError(t, err)
	assert.False(t, exists, "some-new-chain should not exist in test file")

	// Verify file wasn't modified
	afterContent, err := os.ReadFile(testFile)
	require.NoError(t, err)
	assert.Equal(t, string(originalContent), string(afterContent), "file should not be modified during read operations")
}

