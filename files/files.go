package files

import (
	"fmt"
	"os"
	"path/filepath"

	itbasisTestUtils "github.com/itbasis/go-test-utils/v2"
	. "github.com/onsi/gomega"
)

func ReadFile(filePath string) []byte {
	itbasisTestUtils.TestLogger.Info(fmt.Sprintf("read file: %s", filePath))

	Ω(filePath).Should(BeARegularFile())

	bytes, err := os.ReadFile(filepath.Clean(filePath))
	Ω(err).Should(Succeed())

	return bytes
}
