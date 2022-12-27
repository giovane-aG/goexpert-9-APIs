package database

import (
	"time"

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
	db.Last(&productCreated, "id = ?", newProduct.ID.String())

	assert.Nil(t, err)

	assert.NotNil(t, productCreated)
	assert.NotNil(t, productCreated.CreatedAt)
	assert.NotNil(t, productCreated.ID)
	assert.NotNil(t, productCreated.Name)
	assert.NotNil(t, productCreated.Price)

	assert.Equal(t, productCreated.ID, newProduct.ID)
	assert.Equal(t, productCreated.Name, newProduct.Name)
	assert.Equal(t, productCreated.Price, newProduct.Price)
	assert.EqualValues(t, productCreated.CreatedAt.Format(time.Stamp), newProduct.CreatedAt.Format(time.Stamp))
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var products []entity.Product
	p := NewProduct(db)

	products, _ = p.FindAll(2, 3, "")
	assert.NotNil(t, products)
}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	p := NewProduct(db)

	var newProduct *entity.Product
	newProduct, _ = entity.NewProduct("Monitor Husky Tempest 34'", 2299.99)

	err = p.Create(newProduct)
	if err != nil {
		panic(err)
	}

	foundProduct, err := p.FindById(newProduct.ID.String())
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, foundProduct)
	assert.Equal(t, newProduct.ID, foundProduct.ID)
	assert.Equal(t, newProduct.CreatedAt.Format(time.Stamp), foundProduct.CreatedAt.Format(time.Stamp))
	assert.Equal(t, newProduct.Name, foundProduct.Name)
	assert.Equal(t, newProduct.Price, foundProduct.Price)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	p := NewProduct(db)

	var newProduct *entity.Product
	newProduct, _ = entity.NewProduct("Monitor Husky Tempest 34'", 2299.99)
	err = p.Create(newProduct)

	newProduct.Price = 2184.99
	newProduct.Name = "Monitor Husky Tempest 34 polegadas"

	err = p.Update(newProduct)

	productUpdated, err := p.FindById(newProduct.ID.String())
	if err != nil {
		panic(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, productUpdated)
	assert.Equal(t, newProduct.Price, productUpdated.Price)
	assert.Equal(t, newProduct.Name, productUpdated.Name)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	p := NewProduct(db)

	var newProduct *entity.Product
	newProduct, _ = entity.NewProduct("Monitor Husky Tempest 34'", 2299.99)
	err = p.Create(newProduct)

	err = p.Delete(newProduct.ID.String())

	productDeleted, err := p.FindById(newProduct.ID.String())
	if err != nil {
		panic(err)
	}

	assert.Nil(t, productDeleted)
}
