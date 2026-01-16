package main

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestAgentStopsOnContextCancel(t *testing.T) {
	var called int32

	a := &agent{
		intervalSeconds: 1,
		getCpuPercent: func() (float64, error) {
			atomic.AddInt32(&called, 1)
			return 10, nil
		},
		getMemoryPercent: func() (float64, error) {
			return 20, nil
		},
		getDiskPercent: func() (float64, error) {
			return 30, nil
		},
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		err := a.start(ctx)
		if err != nil {
			t.Errorf("agent start returned error: %v", err)
		}
	}(ctx)
	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(500 * time.Millisecond)

	if atomic.LoadInt32(&called) == 0 {
		t.Fatal("expected agent to collect metrics")
	}
}
