package files

import (
	"log/slog"
	"os"
	"path/filepath"

	. "github.com/onsi/gomega"
)

func ReadFile(filePath string) []byte {
	slog.Info("reading file", slog.String("file", filePath))

	Expect(filePath).To(BeARegularFile())

	bytes, err := os.ReadFile(filepath.Clean(filePath))
	Expect(err).To(Succeed())

	return bytes
}
