package prettyTimer

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type TimingStats struct {
	Count     int
	TotalTime time.Duration
	MinTime   time.Duration
	MaxTime   time.Duration
	Timings   []time.Duration // To keep track of individual timings for percentile calculation
	startTime time.Time       // To keep track of the time when Start() is called
}

// Initialize a TimingStats instance with reasonable defaults
func NewTimingStats() *TimingStats {
	return &TimingStats{
		MinTime: time.Duration(math.MaxInt64), // Initialize to a high value
		MaxTime: time.Duration(math.MinInt64), // Initialize to a low value
	}
}

// Record a new timing
func (t *TimingStats) RecordTiming(duration time.Duration) {
	t.Count++
	t.TotalTime += duration
	if duration < t.MinTime {
		t.MinTime = duration
	}
	if duration > t.MaxTime {
		t.MaxTime = duration
	}
	t.Timings = append(t.Timings, duration)
}

// Start the timer
func (t *TimingStats) Start() {
	t.startTime = time.Now()
}

// Finish the timer and record the timing
func (t *TimingStats) Finish() {
	elapsed := time.Since(t.startTime)
	t.RecordTiming(elapsed)
}

// Calculate percentile
func (t *TimingStats) CalculatePercentile(percentile float64) time.Duration {
	if t.Count == 0 {
		return 0
	}
	sort.Slice(t.Timings, func(i, j int) bool { return t.Timings[i] < t.Timings[j] })

	index := int(math.Ceil((percentile/100.0)*float64(len(t.Timings)))) - 1
	if index < 0 {
		return t.Timings[0]
	}
	return t.Timings[index]
}

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
)

// Print stats to stdout
func (t *TimingStats) PrintStats() {
	if t.Count == 0 {
		fmt.Println(Red + "No timings recorded" + Reset)
		return
	}

	avgTime := t.TotalTime / time.Duration(t.Count)
	fmt.Printf(Green+"Min Time: %s, "+Yellow+"Max Time: %s, "+Purple+"Avg Time: %s, "+Blue+"Count: %d\n"+Reset, t.MinTime, t.MaxTime, avgTime, t.Count)

	// Percentile calculations
	fmt.Printf(Red+"50th: %s, "+Green+"90th: %s, "+Purple+"99th: %s\n"+Reset, t.CalculatePercentile(50), t.CalculatePercentile(90), t.CalculatePercentile(99))
}

// Return stats as struct
type Stats struct {
	MinTime   time.Duration
	MaxTime   time.Duration
	AvgTime   time.Duration
	Count     int
	Percent50 time.Duration
	Percent90 time.Duration
	Percent99 time.Duration
}

// Return stats as struct
func (t *TimingStats) GetStats() Stats {
	if t.Count == 0 {
		return Stats{}
	}

	return Stats{
		MinTime:   t.MinTime,
		MaxTime:   t.MaxTime,
		AvgTime:   t.TotalTime / time.Duration(t.Count),
		Count:     t.Count,
		Percent50: t.CalculatePercentile(50),
		Percent90: t.CalculatePercentile(90),
		Percent99: t.CalculatePercentile(99),
	}
}
