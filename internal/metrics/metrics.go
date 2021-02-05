package metrics

import (
	"time"
)

// This bare bones metrics package accumulates and shares integer and duration values.
// It is NOT thread safe.
// It globally shares the metrics set.

// Data holds the metrics and calculated statistics on the metrics
type Data struct {
	Integers           map[string]*[]int
	Durations          map[string]*[]time.Duration
	IntegerStatistics  map[string]IntegerStats
	DurationStatistics map[string]DurationStats
}

// IntegerStats hold the statistics on the integer metrics
type IntegerStats struct {
	Count   int
	Mean    int
	Maximum int
	Minimum int
}

// DurationStats hold the statistics on the duration metrics
type DurationStats struct {
	Count   int
	Mean    time.Duration
	Maximum time.Duration
	Minimum time.Duration
}

var (
	// Metrics and calculated statistics are the accumulated values
	Metrics = Data{
		map[string]*[]int{},
		map[string]*[]time.Duration{},
		map[string]IntegerStats{},
		map[string]DurationStats{},
	}
)

// AddInteger accumulates integer samples by label
func AddInteger(label string, value int) {
	if v, ok := Metrics.Integers[label]; ok {
		*v = append(*v, value)
	} else {
		Metrics.Integers[label] = &[]int{value}
	}
}

// AddDuration accumulates time.Duration samples by label
func AddDuration(label string, value time.Duration) {
	if v, ok := Metrics.Durations[label]; ok {
		*v = append(*v, value)
	} else {
		Metrics.Durations[label] = &[]time.Duration{value}
	}
}

// CalculateStatistics crunches the numbers
func CalculateStatistics() {
	Metrics.IntegerStatistics = map[string]IntegerStats{}
	if len(Metrics.Integers) > 0 {
		for k, v := range Metrics.Integers {
			stats := IntegerStats{}
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
			stats.Count = len(*v)
			stats.Mean = sum / len(*v)
			stats.Maximum = max
			stats.Minimum = min
			Metrics.IntegerStatistics[k] = stats
		}
	}

	Metrics.DurationStatistics = map[string]DurationStats{}
	if len(Metrics.Durations) > 0 {
		for k, v := range Metrics.Durations {
			stats := DurationStats{}
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
			stats.Count = len(*v)
			stats.Mean = sum / time.Duration(len(*v))
			stats.Maximum = max
			stats.Minimum = min
			Metrics.DurationStatistics[k] = stats
		}
	}
}
