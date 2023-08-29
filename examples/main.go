package main

import (
	"time"

	"github.com/8ff/prettyTimer"
)

func main() {
	ts := prettyTimer.NewTimingStats()
	ts.RecordTiming(1 * time.Second)
	ts.RecordTiming(2 * time.Second)
	ts.RecordTiming(3 * time.Second)
	ts.PrintStats()
}
