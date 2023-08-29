package prettyTimer

import (
	"testing"
	"time"
)

func TestTimingStats(t *testing.T) {
	t.Run("Test empty stats", func(t *testing.T) {
		ts := NewTimingStats()
		ts.PrintStats()
	})

	t.Run("Test single timing", func(t *testing.T) {
		ts := NewTimingStats()
		ts.RecordTiming(1 * time.Second)
		ts.PrintStats()
	})

	t.Run("Test multiple timings", func(t *testing.T) {
		ts := NewTimingStats()
		ts.RecordTiming(1 * time.Second)
		ts.RecordTiming(2 * time.Second)
		ts.RecordTiming(3 * time.Second)
		ts.PrintStats()
	})

	t.Run("Test percentile calculation", func(t *testing.T) {
		ts := NewTimingStats()
		ts.RecordTiming(1 * time.Second)
		ts.RecordTiming(2 * time.Second)
		ts.RecordTiming(3 * time.Second)
		ts.RecordTiming(4 * time.Second)
		ts.RecordTiming(5 * time.Second)

		if ts.CalculatePercentile(50) != 3*time.Second {
			t.Errorf("Expected 50th percentile to be 3s, but got %s", ts.CalculatePercentile(50))
		}

		if ts.CalculatePercentile(90) != 5*time.Second {
			t.Errorf("Expected 90th percentile to be 5s, but got %s", ts.CalculatePercentile(90))
		}

		if ts.CalculatePercentile(99) != 5*time.Second {
			t.Errorf("Expected 99th percentile to be 5s, but got %s", ts.CalculatePercentile(99))
		}
	})
}
