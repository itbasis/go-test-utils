package files_test

import (
	"embed"
	"os"

	"github.com/itbasis/go-test-utils/v4/files"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

//go:embed testdata/*
var testData embed.FS

var _ = ginkgo.DescribeTableSubtree(
	"ReadFile", func(fileName, wantContent string) {
		ginkgo.It(
			"Real OS", func() {
				gomega.Expect(fileName).To(gomega.BeARegularFile())
				gomega.Expect(files.ReadFile(os.ReadFile, fileName)).To(gomega.BeEquivalentTo(wantContent))
			},
		)

		ginkgo.It(
			"Embedded FS", func() {
				gomega.Expect(files.ReadFile(testData.ReadFile, fileName)).To(gomega.BeEquivalentTo(wantContent))
			},
		)
	},
	ginkgo.Entry(nil, "testdata/001.txt", "test 001\n"),
	ginkgo.Entry(nil, "testdata/002.txt", "test 002\n"),
)
