package time

import (
	"context"
	"time"
)

type TimeHelper interface {
	SleepOrWaitForContextGetCancelled(ctx context.Context, delay time.Duration)
}

type timeHelper struct {
}

func NewTimeHelper() (TimeHelper, error) {
	return &timeHelper{}, nil
}

func (th *timeHelper) SleepOrWaitForContextGetCancelled(
	ctx context.Context,
	delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
