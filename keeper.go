package keeper

import (
	"context"
	"errors"
	"time"
)

var FailedGetChannel = errors.New("failed to get Channel")

type result struct {
	value interface{}
	err   error
}

// ExecWithContext wait result of f() until context canceled
func ExecWithContext(ctx context.Context, f func() (interface{}, error)) (interface{}, error) {
	resultCh := make(chan result)

	go func() {
		defer close(resultCh)
		resultCh <- func() result {
			i, e := f()
			return result{value: i, err: e}
		}()
	}()
	return waitResult(ctx, resultCh)
}

// wait channel result until context done
func waitResult(ctx context.Context, ch chan result) (interface{}, error) {
	var i result
	for {
		select {
		case <-ctx.Done():
			return i.value, ctx.Err()
		case i, ok := <-ch:
			if !ok {
				return nil, FailedGetChannel
			}
			return i.value, i.err
		default:
		}
		time.Sleep(1 * time.Millisecond)
	}
}
