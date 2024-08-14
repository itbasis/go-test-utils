package files

import (
	"log/slog"
	"path/filepath"

	. "github.com/onsi/gomega"
)

type ReadFileFn func(filename string) ([]byte, error)

func ReadFile(readFileFn ReadFileFn, filePath string) []byte {
	slog.Info("reading file", slog.String("file", filePath))

	bytes, err := readFileFn(filepath.Clean(filePath))
	Expect(err).To(Succeed())

	return bytes
}
