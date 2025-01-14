package database

import (
	"github.com/ThalesLoreto/product-api/internal/entity"
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
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err := p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&products).Error
		return products, err
	}
	err := p.DB.Order("created_at " + sort).Find(&products).Error
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Update(id string, fields interface{}) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Model(&product).Updates(fields).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
