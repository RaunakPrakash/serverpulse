package metrics

import "testing"

func TestSystemCollector_Calls(t *testing.T) {
	c := NewSystemCollector()

	if _, err := c.Memory(); err != nil {
		t.Fatal(err)
	}

	if _, err := c.Disk(); err != nil {
		t.Fatal(err)
	}
}
