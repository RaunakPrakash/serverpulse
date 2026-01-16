package main

import (
	"context"

	"github.com/RaunakPrakash/serverpulse/internal/metrics"
)

type Container struct {
	flag            *flags
	systemCollector *metrics.SystemCollector
	agent           *agent
}

func NewContainer(flg *flags) (*Container, func(c context.Context)) {
	c := &Container{
		flag: flg,
	}
	c.systemCollector = metrics.NewSystemCollector()
	c.agent = newAgent(c)
	cleanup := func(c context.Context) {}
	return c, cleanup
}
