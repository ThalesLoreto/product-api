package database

import "github.com/ThalesLoreto/product-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
