package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("Create a new product", func(t *testing.T) {
		name := "Product Test"
		price := 1000

		product, err := NewProduct(name, price)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.NotEmpty(t, product.ID)
		assert.Equal(t, product.Name, name)
		assert.Equal(t, product.Price, price)
	})
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	t.Run("Name is required", func(t *testing.T) {
		name := ""
		price := 1000

		product, err := NewProduct(name, price)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		assert.Equal(t, err, ErrNameRequired)
	})
}

func TestProductWhenPriceIsZero(t *testing.T) {
	t.Run("Price is required", func(t *testing.T) {
		name := "Product Test"
		price := 0

		product, err := NewProduct(name, price)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		assert.Equal(t, err, ErrPriceRequired)
	})
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	t.Run("Price is invalid", func(t *testing.T) {
		name := "Product Test"
		price := -1000

		product, err := NewProduct(name, price)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		assert.Equal(t, err, ErrInvalidPrice)
	})
}
