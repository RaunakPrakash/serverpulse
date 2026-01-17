package main

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestAgentStopsOnContextCancel(t *testing.T) {
	var called int32
	done := make(chan bool, 1)

	a := &agent{
		intervalSeconds: 1,
		debug:           false,
		getCpuPercent: func() (float64, error) {
			if atomic.AddInt32(&called, 1) == 1 {
				done <- true
			}
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

	go func() {
		err := a.start(ctx)
		if err != nil {
			t.Errorf("agent start returned error: %v", err)
		}
	}()

	select {
	case <-done:
		// First collection happened
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for first collection")
	}

	cancel()

	time.Sleep(100 * time.Millisecond)

	if atomic.LoadInt32(&called) == 0 {
		t.Fatal("expected agent to collect metrics")
	}
}
