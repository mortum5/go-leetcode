package task1603

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesignPark(t *testing.T) {
	dp := Constructor(1, 1, 0)
	assert.Equal(t, true, dp.AddCar(1))
	assert.Equal(t, true, dp.AddCar(2))
	assert.Equal(t, false, dp.AddCar(3))
	assert.Equal(t, false, dp.AddCar(1))
}
