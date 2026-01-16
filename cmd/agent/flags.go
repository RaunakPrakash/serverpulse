package main

import (
	"flag"
	"os"
)

type flags struct {
	intervalSeconds int
	endpoint        string
	apiKey          string
	debug           bool
}

func newFlags() *flags {
	return &flags{
		intervalSeconds: 60,
		endpoint:        "",
		apiKey:          "",
		debug:           false,
	}
}

func getFlags() *flags {
	flg := newFlags()

	fs := flag.NewFlagSet("serverpulse-agent", flag.ExitOnError)

	fs.IntVar(
		&flg.intervalSeconds,
		"interval",
		flg.intervalSeconds,
		"metrics collection interval in seconds",
	)

	fs.StringVar(
		&flg.endpoint,
		"endpoint",
		flg.endpoint,
		"server endpoint URL",
	)

	fs.StringVar(
		&flg.apiKey,
		"apikey",
		flg.apiKey,
		"API key for authentication",
	)

	fs.BoolVar(
		&flg.debug,
		"debug",
		flg.debug,
		"enable debug logging",
	)

	_ = fs.Parse(os.Args[1:])
	return flg
}
