package time

import (
	"context"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/onsi/gomega"
)

func NewContextWithFakeClock(date, layout string) context.Context {
	return FromContextWithFakeClock(context.Background(), date, layout)
}

func FromContextWithFakeClock(ctx context.Context, date, layout string) context.Context {
	var t, err = time.Parse(layout, date)
	gomega.Expect(err).To(gomega.Succeed())

	return FromContextWithFakeClockAt(ctx, t)
}

func FromContextWithFakeClockAt(ctx context.Context, t time.Time) context.Context {
	return clockwork.AddToContext(ctx, clockwork.NewFakeClockAt(t))
}
