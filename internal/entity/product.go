package entity

import (
	"errors"
	"time"

	"github.com/ThalesLoreto/product-api/pkg/entity"
)

var (
	ErrIdRequired    = errors.New("ID is required")
	ErrNameRequired  = errors.New("Name is required")
	ErrPriceRequired = errors.New("Price is required")
	ErrInvalidID     = errors.New("Invalid ID")
	ErrInvalidPrice  = errors.New("Invalid Price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price == 0 {
		return ErrPriceRequired
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
