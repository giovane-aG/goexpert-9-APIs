package database

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})
	p := NewProduct(db)

	var newProduct *entity.Product
	newProduct, _ = entity.NewProduct("Monitor Husky Tempest 34'", 2299.99)

	err = p.Create(newProduct)

	var productCreated *entity.Product
	db.Find(&productCreated, "id = ?", newProduct.ID.String())

	assert.Nil(t, err)
	assert.NotNil(t, productCreated)
	assert.NotNil(t, productCreated.CreatedAt)
	assert.NotNil(t, productCreated.ID)
	assert.NotNil(t, productCreated.Name)
	assert.NotNil(t, productCreated.Price)
	assert.Equal(t, productCreated.ID, newProduct.ID)
}
