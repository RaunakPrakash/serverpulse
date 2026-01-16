package main

import "testing"

func TestNewContainer(t *testing.T) {
	c, _ := NewContainer(newFlags())

	if c == nil {
		t.Fatal("container is nil")
	}

	if c.agent == nil {
		t.Fatal("agent not initialized")
	}

	if c.systemCollector == nil {
		t.Fatal("systemCollector not initialized")
	}
}
