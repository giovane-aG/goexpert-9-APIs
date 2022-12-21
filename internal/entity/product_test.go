package entity

import (
	"testing"

	"github.com/giovane-aG/goexpert/9-APIs/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("product", 10.0)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, product, "Product should not be nil")
	assert.Equal(t, "product", product.Name, "Name should be product")
	assert.Equal(t, 10.0, product.Price, "Price should be 10.0")
	assert.NotEmpty(t, product.ID, "ID should not be empty")
	assert.NotEmpty(t, product.CreatedAt, "CreatedAt should not be empty")
}
func TestNewProductPriceRequired(t *testing.T) {
	product, err := NewProduct("product", 0.0)

	assert.Nil(t, product)
	assert.Equal(t, errors.ErrRequiredPrice, err, "Price is required")
}
func TestNewProductInvalidPrice(t *testing.T) {
	product, err := NewProduct("product", 10.0)

	assert.Nil(t, product)
	assert.Equal(t, errors.ErrInvalidPrice, err, "Price should be greater than 0")
}

func TestNewProductInvalidName(t *testing.T) {
	product, err := NewProduct("", 10.0)

	assert.Nil(t, product)
	assert.Equal(t, errors.ErrRequiredName, err, "Name should not be empty")
}
