package ginkgo

import (
	"log/slog"
	"testing"

	"github.com/dusted-go/logging/prettylog"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func InitGinkgoSuite(t *testing.T, name string) {
	InitGinkgoSuiteWithSlogOptions(
		t, name,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	)
}

func InitGinkgoSuiteWithSlogOptions(t *testing.T, name string, slogOptions *slog.HandlerOptions) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	slog.SetDefault(
		slog.New(
			prettylog.New(slogOptions, prettylog.WithDestinationWriter(ginkgo.GinkgoWriter)),
		),
	)

	ginkgo.RunSpecs(t, name)
}
