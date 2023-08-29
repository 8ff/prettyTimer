package main

import (
	"fmt"
	"time"

	"github.com/8ff/prettyTimer"
)

func main() {
	ts := prettyTimer.NewTimingStats()
	ts.RecordTiming(1 * time.Second)
	ts.RecordTiming(2 * time.Second)
	ts.RecordTiming(3 * time.Second)
	ts.PrintStats()

	// Return stats and get single value
	stats := ts.GetStats()
	fmt.Printf("99th percentile: %s\n", stats.Percent99)
}
