package main

import (
	"context"
	"log"
	"time"
)

type agent struct {
	intervalSeconds  int
	debug            bool
	getCpuPercent    func() (float64, error)
	getMemoryPercent func() (float64, error)
	getDiskPercent   func() (float64, error)
}

func newAgent(c *Container) *agent {
	d := c.systemCollector
	return &agent{
		intervalSeconds:  c.flag.intervalSeconds,
		debug:            c.flag.debug,
		getCpuPercent:    d.CPU,
		getMemoryPercent: d.Memory,
		getDiskPercent:   func() (float64, error) { return d.Disk(c.flag.diskPath) },
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
		log.Printf("failed to collect CPU metrics: %v", err)
		return
	}

	mem, err := a.getMemoryPercent()
	if err != nil {
		log.Printf("failed to collect memory metrics: %v", err)
		return
	}

	disk, err := a.getDiskPercent()
	if err != nil {
		log.Printf("failed to collect disk metrics: %v", err)
		return
	}

	if a.debug {
		log.Printf("metrics collected: CPU=%.2f%%, Memory=%.2f%%, Disk=%.2f%%", cpu, mem, disk)
	} else {
		log.Printf("CPU: %.2f%%, Memory: %.2f%%, Disk: %.2f%%", cpu, mem, disk)
	}
}
