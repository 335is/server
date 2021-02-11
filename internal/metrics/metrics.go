package metrics

import (
	"sync"
	"time"
)

// This bare bones metrics package accumulates and shares integer and duration values.

// Organization stores data about this service instance
type Organization struct {
	Region      string // US, EU, TW, CH
	Datacenter  string // WEST, EAST, OC, LAX
	Network     string // GARBOLIC, CENTRAL, DEPAK
	Environment string // DEV, TEST, PROD
	Hostname    string // SOLONAKER, DEEMSTAM
	IPAddress   string // 10.25.28.17, 192.168.23.55
	Service     string // GYMSNAPPER, TELEDIRK
	Instance    string // A, B
}

// Statistics calculated
type Statistics struct {
	IntegerStats  map[string]IStats
	DurationStats map[string]DStats
}

// IStats holds the statistics on the integer metrics
type IStats struct {
	Count   int
	Mean    int
	Maximum int
	Minimum int
}

// DStats holds the statistics on the duration metrics
type DStats struct {
	Count   int
	Mean    time.Duration
	Maximum time.Duration
	Minimum time.Duration
}

// Data holds the metrics and calculated statistics on the metrics
type data struct {
	integers  map[string]*[]int
	durations map[string]*[]time.Duration
}

var (
	// Metrics and calculated statistics are the accumulated values
	metrics = data{
		map[string]*[]int{},
		map[string]*[]time.Duration{},
	}
	mutexIntegers  = &sync.Mutex{}
	mutexDurations = &sync.Mutex{}
)

// AddInteger accumulates integer samples by label
func AddInteger(label string, value int) {
	mutexIntegers.Lock()
	if v, ok := metrics.integers[label]; ok {
		*v = append(*v, value)
	} else {
		metrics.integers[label] = &[]int{value}
	}
	mutexIntegers.Unlock()
}

// AddDuration accumulates time.Duration samples by label
func AddDuration(label string, value time.Duration) {
	mutexDurations.Lock()
	if v, ok := metrics.durations[label]; ok {
		*v = append(*v, value)
	} else {
		metrics.durations[label] = &[]time.Duration{value}
	}
	mutexDurations.Unlock()
}

// CalculateStatistics crunches the numbers and returns the stats
func CalculateStatistics() Statistics {
	stats := Statistics{
		map[string]IStats{},
		map[string]DStats{},
	}

	if len(metrics.integers) > 0 {
		mutexIntegers.Lock()
		for k, v := range metrics.integers {
			sum := 0
			max := 0
			min := 0
			for i, d := range *v {
				sum += d
				if i == 0 {
					max = d
					min = d
				} else {
					if d > max {
						max = d
					}
					if d < min {
						min = d
					}
				}
			}
			stats.IntegerStats[k] = IStats{len(*v), sum / len(*v), max, min}
		}
		mutexIntegers.Unlock()
	}

	if len(metrics.durations) > 0 {
		mutexDurations.Lock()
		for k, v := range metrics.durations {
			sum := time.Duration(0)
			max := time.Duration(0)
			min := time.Duration(0)
			for i, d := range *v {
				sum += d
				if i == 0 {
					max = d
					min = d
				} else {
					if d > max {
						max = d
					}
					if d < min {
						min = d
					}
				}
			}
			stats.DurationStats[k] = DStats{len(*v), sum / time.Duration(len(*v)), max, min}
		}
		mutexDurations.Unlock()
	}

	return stats
}
