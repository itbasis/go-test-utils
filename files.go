package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadFile(t *testing.T, filePath string) []byte {
	TestLogger.Info().Msgf("read file: %s", filePath)

	assert.FileExists(t, filePath)

	bytes, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		t.Error(err)
	}

	return bytes
}
