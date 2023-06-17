package testutils

import (
	"github.com/fgrosse/zaptest"
	"github.com/onsi/ginkgo/v2"
)

var TestLogger = zaptest.LoggerWriter(ginkgo.GinkgoWriter)
