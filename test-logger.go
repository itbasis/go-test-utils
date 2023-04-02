package test_utils

import (
	"context"
	"os"
	"time"

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

func TestLoggerWithContext(ctx context.Context) context.Context {
	return TestLogger.WithContext(ctx)
}
