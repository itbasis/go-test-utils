package files

import (
	"io"
	"log/slog"

	. "github.com/onsi/gomega"
	"golang.org/x/tools/godoc/vfs"
)

func Close(closer io.Closer) {
	if closer == nil {
		slog.Warn("closer is nil")

		return
	}

	Expect(closer.Close()).To(Succeed())
}

func ReadFile(fs vfs.Opener, filePath string) []byte {
	fileReader := FileReader(fs, filePath)

	defer Close(fileReader)

	bytes, err := io.ReadAll(fileReader)
	Expect(err).To(Succeed())

	return bytes
}

func FileReader(fs vfs.Opener, filePath string) vfs.ReadSeekCloser {
	slog.Info("reading file", slog.String("file", filePath))

	rsc, err := fs.Open(filePath)
	Expect(err).To(Succeed())

	return rsc
}
