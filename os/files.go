package files

import (
	"io"
	"log/slog"

	"github.com/onsi/gomega"
	"golang.org/x/tools/godoc/vfs"
)

func Close(closer io.Closer) {
	if closer == nil {
		slog.Warn("closer is nil")

		return
	}

	gomega.Expect(closer.Close()).To(gomega.Succeed())
}

func ReadFile(fs vfs.Opener, filePath string) []byte {
	fileReader := FileReader(fs, filePath)

	defer Close(fileReader)

	bytes, err := io.ReadAll(fileReader)
	gomega.Expect(err).To(gomega.Succeed())

	return bytes
}

func FileReader(fs vfs.Opener, filePath string) vfs.ReadSeekCloser {
	slog.Info("reading file", slog.String("file", filePath))

	rsc, err := fs.Open(filePath)
	gomega.Expect(err).To(gomega.Succeed())

	return rsc
}
