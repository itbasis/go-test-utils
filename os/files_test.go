package files_test

// tag::snippet[]
import (
	"embed"
	"io"
	"io/fs"

	itbasisTestOs "github.com/itbasis/go-test-utils/v5/os"
	"golang.org/x/tools/godoc/vfs"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _entries = []ginkgo.TableEntry{
	ginkgo.Entry(nil, "testdata/001.txt", "test 001\n"),
	ginkgo.Entry(nil, "testdata/002.txt", "test 002\n"),
}

// end::snippet[]

// tag::snippet_os[]
var _ = ginkgo.DescribeTable(
	"ReadFile :: OS", func(fileName, wantContent string) {
		gomega.Expect(fileName).To(gomega.BeARegularFile())

		gomega.Expect(itbasisTestOs.ReadFile(vfs.OS("."), fileName)).To(gomega.BeEquivalentTo(wantContent))
	},
	_entries,
)

// end::snippet_os[]

// tag::snippet_reader_os[]
var _ = ginkgo.DescribeTable(
	"FileReader :: OS", func(fileName, wantContent string) {
		var reader = itbasisTestOs.FileReader(vfs.OS("."), fileName)

		defer itbasisTestOs.Close(reader)

		gomega.Expect(io.ReadAll(reader)).To(gomega.BeEquivalentTo(wantContent))
	},
	_entries,
)

// end::snippet_reader_os[]

// tag::snippet_embedded[]

//go:embed testdata/*
var testData embed.FS

var _ = ginkgo.DescribeTable(
	"ReadFile :: Embedded FS", func(fileName, wantContent string) {
		gomega.Expect(itbasisTestOs.ReadFile(vfs.FromFS(testData), fileName)).To(gomega.BeEquivalentTo(wantContent))
	},
	_entries,
)

// end::snippet_embedded[]

// tag::snippet_reader_embedded[]
var _ = ginkgo.DescribeTable(
	"FileReader :: Embedded FS", func(fileName, wantContent string) {
		var reader = itbasisTestOs.FileReader(vfs.FromFS(testData), fileName)

		defer itbasisTestOs.Close(reader)

		gomega.Expect(io.ReadAll(reader)).To(gomega.BeEquivalentTo(wantContent))
	},
	_entries,
)

// end::snippet_reader_embedded[]

var _ = ginkgo.Describe(
	"Close", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.It(
			"Close nil", func() {
				gomega.Expect(
					gomega.InterceptGomegaFailure(
						func() {
							itbasisTestOs.Close(io.Closer(nil))
						},
					),
				).To(gomega.Succeed())
			},
		)

		ginkgo.It(
			"Fail close", func() {
				var reader = itbasisTestOs.FileReader(vfs.OS("."), "testdata/001.txt")

				gomega.Expect(reader.Close()).To(gomega.Succeed())
				gomega.Expect(
					gomega.InterceptGomegaFailure(
						func() {
							itbasisTestOs.Close(reader)
						},
					),
				).Error().To(gomega.MatchError(gomega.ContainSubstring(fs.ErrClosed.Error())))
			},
		)
	},
)
