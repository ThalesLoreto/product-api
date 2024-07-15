package database

import (
	"testing"

	"github.com/ThalesLoreto/product-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)
	if err := userDB.Create(user); err != nil {
		t.Errorf("could not create user: %v", err)
	}

	var userFound entity.User
	if err := db.First(&userFound, user.ID).Error; err != nil {
		t.Errorf("could not find user: %v", err)
	}
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)
	if err := userDB.Create(user); err != nil {
		t.Errorf("could not create user: %v", err)
	}

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
