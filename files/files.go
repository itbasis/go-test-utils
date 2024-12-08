package files

import (
	"log/slog"

	. "github.com/onsi/gomega"
)

type ReadFileFn func(filename string) ([]byte, error)

func ReadFile(readFileFn ReadFileFn, filePath string) []byte {
	slog.Info("reading file", slog.String("file", filePath))

	bytes, err := readFileFn(filePath)
	Expect(err).To(Succeed())

	return bytes
}
