package keeper

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecWithContext(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
		f       func() (interface{}, error)
		want    interface{}
		wantErr error
	}{
		{
			name:    "timeout",
			timeout: 400 * time.Millisecond,
			f: func() (interface{}, error) {
				time.Sleep(1 * time.Second)
				return []int{1, 2, 3}, nil
			},
			want:    nil,
			wantErr: context.DeadlineExceeded,
		},
		{
			name:    "finish in time",
			timeout: 400 * time.Millisecond,
			f: func() (interface{}, error) {
				return []int{1, 2, 3}, nil
			},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "finish in time 2",
			timeout: 400 * time.Millisecond,
			f: func() (interface{}, error) {
				time.Sleep(300 * time.Millisecond)
				return []int{1, 2, 3}, nil
			},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "fail func",
			timeout: 400 * time.Millisecond,
			f: func() (interface{}, error) {
				time.Sleep(300 * time.Millisecond)
				return nil, errors.New("error")
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()
			got, err := ExecWithContext(ctx, tt.f)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
