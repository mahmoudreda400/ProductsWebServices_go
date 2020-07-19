package product

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository struct {
	DB *gorm.DB
}

func providePeoductRepository(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) findAll() []Product {
	var products []Product
	p.DB.Find(&products)
	return products
}

func (p *ProductRepository) findByID(id int) Product {
	var product Product
	p.DB.First(&product, id)
	return product
}

func (p *ProductRepository) save(product Product) Product {
	p.DB.Save(&product)
	return product
}

func (p *ProductRepository) delete(prodcut Product) {
	p.DB.Delete(&prodcut)
}
