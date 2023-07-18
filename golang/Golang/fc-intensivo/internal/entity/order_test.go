package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := &Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfItGetsAnErrorIfPriceIsLessThanZero(t *testing.T) {
	order := &Order{ID: "1"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErrorIfTaxIsLessThanZero(t *testing.T) {
	order := &Order{ID: "1", Price: 10}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestIfItGetsAnErrorIfTaxIsZero(t *testing.T) {
	order := &Order{ID: "1", Price: 10, Tax: 0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := &Order{ID: "1", Price: 10, Tax: 1}
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, 11.0, order.FinalPrice)
}

func TestIfItGetsAnErrorIfIDIsBlankWhenCalculatingFinalPrice(t *testing.T) {
	order := &Order{}
	assert.Error(t, order.CalculateFinalPrice(), "id is required")
}