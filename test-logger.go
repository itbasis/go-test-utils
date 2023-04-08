package testutils

import (
	"os"
	"time"

	"github.com/onsi/ginkgo/v2"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var TestLogger = log.
	Output(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		},
	).
	Level(zerolog.TraceLevel).
	With().
	Caller().
	Logger()

func ConfigureTestLoggerForGinkgo() {
	TestLogger.Output(ginkgo.GinkgoWriter)
	log.Logger = TestLogger
	zerolog.DefaultContextLogger = &TestLogger
}
