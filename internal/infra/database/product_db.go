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
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product

	if limit > 0 && page > 0 {
		p.DB.Offset(
			(page - 1) * limit,
		).Limit(limit)
	}

	if sort != "" {
		p.DB.Order("created_at " + sort)
	}

	err := p.DB.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product *entity.Product

	tx := p.DB.Find(&product, "id = ?", id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, nil
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

func (p *Product) Delete(id string) error {
	_, err := p.FindById(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}
