package stddev

import "math"

type rollingStdDev struct {
	periods    int
	dataPoints []float64
	average    float64
	variance   float64
	lastStdDev *float64
}

func newStdDev(periods int) *rollingStdDev {
	return &rollingStdDev{periods: periods, dataPoints: make([]float64, 0, periods+1)}
}

func (rs *rollingStdDev) ready() bool {
	return len(rs.dataPoints) >= rs.periods
}

func (rs *rollingStdDev) addPoint(new float64) {
	rs.dataPoints = append(rs.dataPoints, new)
	if len(rs.dataPoints) < rs.periods {
		return
	}
	n := float64(rs.periods)

	if rs.lastStdDev != nil {
		old := rs.dataPoints[0]
		delta := new - old
		average := rs.average + delta/n
		rs.variance += delta * (new - average + old - rs.average)
		rs.average = average
	} else {
		// this is the first time I collected slice
		rs.average, rs.variance = stdDevSlice(rs.dataPoints)
	}
	stdDev := math.Sqrt(rs.variance / (n - 1))
	rs.lastStdDev = &stdDev

	if len(rs.dataPoints) > rs.periods {
		rs.dataPoints = rs.dataPoints[1:]
	}
}

func (rs *rollingStdDev) stdDev() float64 {
	if len(rs.dataPoints) < rs.periods {
		return math.NaN()
	}
	return *rs.lastStdDev
}

// This function calculates the standard deviation of the input slice
// by iterating over each value in the slice and computing the mean and
// variance of the data in a single pass.
func stdDevSlice(dataPoints []float64) (float64, float64) {
	n := len(dataPoints)
	if n <= 1 {
		return dataPoints[0], 0.0
	}
	mean := dataPoints[0]
	m2 := 0.0
	for i := 1; i < n; i++ {
		delta := dataPoints[i] - mean
		mean += delta / float64(i+1)
		m2 += delta * (dataPoints[i] - mean)
	}
	return mean, m2
}
