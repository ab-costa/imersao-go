package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCaseShouldReturnAnErrorWhenCreateANewOrderIsCalledWithAnEmptyId(t *testing.T) {
	order := Order{}

	assert.Error(t, order.IsValid(), "Invalid id")
}

func TestCaseShouldReturnAnErrorWhenCreateANewOrderIsCalledWithAnEmptyPrice(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.IsValid(), "Invalid price")
}

func TestCaseShouldReturnAnErrorWhenCreateANewOrderIsCalledWithAnEmptyTax(t *testing.T) {
	order := Order{ID: "123", Price: 10}

	assert.Error(t, order.IsValid(), "Invalid tax")
}

func TestCaseShouldReturnNilWhenCreateANewOrderIsCalledWithValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.0}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}