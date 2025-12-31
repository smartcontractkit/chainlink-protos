package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupChainSelector(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

	tests := []struct {
		name        string
		content     string
		wantEntries []ChainEntry
		wantErr     error
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
			wantErr: nil,
		},
		{
			name:    "missing defaults array",
			content: `some other content`,
			wantErr: ErrNoDefaultsArray,
		},
		{
			name: "empty defaults array",
			content: `
defaults: [
          ]`,
			wantErr: ErrNoEntriesFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			entries, err := parseProtoEntries(tt.content)
			if tt.wantErr != nil {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantEntries, entries)
		})
	}
}

func TestBuildDefaultsArray(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
			t.Parallel()

			exists, err := chainSelectorExists(testFile, tt.chainName)
			require.NoError(t, err)
			assert.Equal(t, tt.wantExists, exists)
		})
	}
}

func TestUpdateProtoFile(t *testing.T) {
	t.Parallel()

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
		t.Parallel()

		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "test.proto")

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
		t.Parallel()

		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "test.proto")

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
		t.Parallel()

		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "test.proto")

		err := os.WriteFile(testFile, []byte(originalContent), 0644)
		require.NoError(t, err)

		err = updateProtoFile(testFile, "ethereum-mainnet", 5009297550715157269)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrChainExists), "expected ErrChainExists, got %v", err)
	})
}

func TestIntegration_DryRun(t *testing.T) {
	t.Parallel()

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
