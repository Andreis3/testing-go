package tax

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount must be greater than zero"))

	err := CalculateTaxAndSave(10000, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(-100, repository)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than zero")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)

}
