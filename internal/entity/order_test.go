package entity_test

import (
	"tax-service/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreateValidOrder tests correct creation of order
func TestCreateValidOrder(t *testing.T) {

	// create order with invalid id should return error
	func() {
		order := entity.Order{}

		assert.Error(t, order.IsValid(), "invalid id")
	}()

	// create order with invalid price should return error
	func() {
		order := entity.Order{
			ID: "123",
		}

		assert.Error(t, order.IsValid(), "invalid price")
	}()

	// create order with invalid tax should return error
	func() {
		order := entity.Order{
			ID:    "123",
			Price: 100,
		}

		assert.Error(t, order.IsValid(), "invalid tax")
	}()

	// create order with valid id, price and tax should return no error
	func() {
		order, err := entity.NewOrder("123", 10, 2)

		assert.NoError(t, err)

		assert.Equal(t, "123", order.ID)

		assert.Equal(t, 10.0, order.Price)

		assert.Equal(t, 2.0, order.Tax)
	}()
}

// TestCalculateFinalPrice tests correct calculation of final price
func TestCalculateFinalPrice(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.NoError(t, err)

	err = order.CalculateFinalPrice()

	assert.NoError(t, err)

	assert.Equal(t, 12.0, order.FinalPrice)
}
