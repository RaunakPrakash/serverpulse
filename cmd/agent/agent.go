package main

import (
	"context"
	"log"
	"time"
)

type agent struct {
	intervalSeconds  int
	getCpuPercent    func() (float64, error)
	getMemoryPercent func() (float64, error)
	getDiskPercent   func() (float64, error)
}

func newAgent(c *Container) *agent {
	d := c.systemCollector
	return &agent{
		intervalSeconds:  c.flag.intervalSeconds,
		getCpuPercent:    d.CPU,
		getMemoryPercent: d.Memory,
		getDiskPercent:   d.Disk,
	}
}

func (a *agent) start(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(a.intervalSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.collectOnce()
		case <-ctx.Done():
			log.Println("agent stopped")
			return nil
		}
	}
}

func (a *agent) collectOnce() {
	cpu, err := a.getCpuPercent()
	if err != nil {
		log.Println("cpu error:", err)
		return
	}

	mem, err := a.getMemoryPercent()
	if err != nil {
		log.Println("memory error:", err)
		return
	}

	disk, err := a.getDiskPercent()
	if err != nil {
		log.Println("disk error:", err)
		return
	}

	log.Printf(
		"CPU: %.2f%%, Memory: %.2f%%, Disk: %.2f%%",
		cpu, mem, disk,
	)
}
