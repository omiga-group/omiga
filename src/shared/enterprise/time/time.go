package time

import (
	"context"
	"time"
)

func SleepOrWaitForContextGetCancelled(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
