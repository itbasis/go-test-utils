package test_utils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadFile(t *testing.T, filePath string) []byte {
	TestLogger.Info().Msgf("read file: %s", filePath)

	assert.FileExists(t, filePath)

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}

	return bytes
}
