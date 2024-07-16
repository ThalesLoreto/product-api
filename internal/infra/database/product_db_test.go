package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ThalesLoreto/product-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	product, _ := entity.NewProduct("Product 1", 10)
	productDB := NewProduct(db)
	if err := productDB.Create(product); err != nil {
		t.Errorf("could not create product: %v", err)
	}
	var productFound entity.Product
	if err := db.First(&productFound, product.ID).Error; err != nil {
		t.Errorf("could not find product: %v", err)
	}
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	for i := 0; i < 10; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Int())
		db.Create(&product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 5, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 5", products[4].Name)

	products, err = productDB.FindAll(2, 5, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 6", products[0].Name)
	assert.Equal(t, "Product 10", products[4].Name)
}

func TestProduct_FindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	product, _ := entity.NewProduct("Product 1", 10)
	db.Create(&product)
	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	product, _ := entity.NewProduct("Product 1", 10)
	db.Create(&product)
	product.Name = "Product 2"
	productDB := NewProduct(db)
	err = productDB.Update(product.ID.String(), product)
	assert.NoError(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to db: %v", err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("could not migrate db: %v", err)
	}
	product, _ := entity.NewProduct("Product 1", 10)
	db.Create(&product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	var productFound entity.Product
	err = db.First(&productFound, product.ID).Error
	assert.Error(t, err)
}
