package metrics

import (
	"testing"
)

func TestSystemCollector_Calls(t *testing.T) {
	c := NewSystemCollector()

	if _, err := c.CPU(); err != nil {
		t.Fatalf("CPU collection failed: %v", err)
	}

	if _, err := c.Memory(); err != nil {
		t.Fatalf("Memory collection failed: %v", err)
	}

	if _, err := c.Disk("/"); err != nil {
		t.Fatalf("Disk collection failed: %v", err)
	}
}

func TestSystemCollector_Disk_InvalidPath(t *testing.T) {
	c := NewSystemCollector()

	_, err := c.Disk("/nonexistent/path")
	if err == nil {
		t.Error("expected error for invalid disk path")
	}
}
