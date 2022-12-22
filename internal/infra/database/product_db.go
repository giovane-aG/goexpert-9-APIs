package database

import (
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/google/uuid"
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

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product *entity.Product
	parsedID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	err = p.DB.Find(product, "id = ?", parsedID).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}
