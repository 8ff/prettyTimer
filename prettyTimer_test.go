package prettyTimer

import (
	"testing"
	"time"
)

// TestNewTimingStats checks if a new TimingStats instance is properly initialized
func TestNewTimingStats(t *testing.T) {
	ts := NewTimingStats()
	if ts.Count != 0 || ts.TotalTime != 0 || ts.MinTime != time.Duration(1<<63-1) || ts.MaxTime != time.Duration(-1<<63) || len(ts.Timings) != 0 {
		t.Errorf("NewTimingStats() did not initialize properly")
	}
}

// TestRecordTiming checks if RecordTiming function works correctly
func TestRecordTiming(t *testing.T) {
	ts := NewTimingStats()
	ts.RecordTiming(10)
	if ts.Count != 1 || ts.TotalTime != 10 || ts.MinTime != 10 || ts.MaxTime != 10 || len(ts.Timings) != 1 {
		t.Errorf("RecordTiming() did not record timing properly")
	}
}

// TestStartAndFinish checks if Start and Finish functions work correctly
func TestStartAndFinish(t *testing.T) {
	ts := NewTimingStats()
	ts.Start()
	time.Sleep(10 * time.Millisecond)
	ts.Finish()
	if ts.Count != 1 || ts.TotalTime == 0 || ts.MinTime == 0 || ts.MaxTime == 0 || len(ts.Timings) != 1 {
		t.Errorf("Start() and Finish() did not work properly")
	}
}

// TestCalculatePercentile checks if CalculatePercentile function works correctly
func TestCalculatePercentile(t *testing.T) {
	ts := NewTimingStats()
	ts.RecordTiming(10)
	ts.RecordTiming(20)
	ts.RecordTiming(30)
	if ts.CalculatePercentile(50) != 20 {
		t.Errorf("CalculatePercentile() did not calculate correctly")
	}
}

// TestPrintStats checks if PrintStats function works, but it does not check the output
func TestPrintStats(t *testing.T) {
	ts := NewTimingStats()
	ts.PrintStats() // Just to cover the function, checking the output is not straightforward
}

// TestGetStats checks if GetStats function works correctly
func TestGetStats(t *testing.T) {
	ts := NewTimingStats()
	ts.RecordTiming(10)
	ts.RecordTiming(20)
	stats := ts.GetStats()
	if stats.MinTime != 10 || stats.MaxTime != 20 || stats.AvgTime != 15 || stats.Count != 2 || stats.Percent50 != 10 || stats.Percent90 != 20 || stats.Percent99 != 20 {
		t.Errorf("GetStats() did not return correct stats. Got: %+v", stats)
	}
}
