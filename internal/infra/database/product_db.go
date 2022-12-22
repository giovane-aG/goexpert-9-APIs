package database

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(p).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product

	err := p.DB.Offset(
		page * limit,
	).Limit(
		limit,
	).Order(sort).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
