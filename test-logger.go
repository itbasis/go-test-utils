package testutils

import (
	"github.com/fgrosse/zaptest"
	"github.com/juju/zaputil/zapctx"
	"github.com/onsi/ginkgo/v2"
)

var TestLogger = zaptest.LoggerWriter(ginkgo.GinkgoWriter)

func init() {
	zapctx.Default = TestLogger
}
