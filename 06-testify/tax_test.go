package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(100)
	assert.Equal(t, 5.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(10000)
	assert.Equal(t, 10.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(100000)
	assert.Equal(t, 20.0, tax)
	assert.Nil(t, err)

	tax, err = CalculateTax(-100)
	assert.Equal(t, 0.0, tax)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}
