package time

import (
	"context"
	"time"
)

type TimeHelper interface {
	WaitUntilCancelled(ctx context.Context)
	SleepOrWaitForContextGetCancelled(ctx context.Context, delay time.Duration)
}

type timeHelper struct {
}

func NewTimeHelper() (TimeHelper, error) {
	return &timeHelper{}, nil
}

func (th *timeHelper) WaitUntilCancelled(ctx context.Context) {
	for {
		if ctx.Err() == context.Canceled {
			break
		}

		th.SleepOrWaitForContextGetCancelled(
			ctx,
			time.Second)
	}
}

func (th *timeHelper) SleepOrWaitForContextGetCancelled(
	ctx context.Context,
	delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
