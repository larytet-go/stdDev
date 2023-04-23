package stddev

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdDev(t *testing.T) {
	t.Parallel()

	mean, variance := stdDevSlice([]float64{1, 2, 3})
	assert.Equal(t, 2.0, mean)
	assert.Equal(t, 2.0, variance)

	mean, variance = stdDevSlice([]float64{2, 3, 4})
	assert.Equal(t, 3.0, mean)
	assert.Equal(t, 2.0, variance)

	rs := newStdDev(3)
	rs.addPoint(1.0)
	rs.addPoint(2.0)
	assert.True(t, math.IsNaN(rs.stdDev()))
	rs.addPoint(3.0)

	actual := rs.stdDev()
	assert.Equal(t, 1.0, actual)

	rs.addPoint(4.0)
	actual = rs.stdDev()
	assert.Equal(t, 1.0, actual)
}
