package main

import (
	"ProductsWebServices/product"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {

	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductAPI)
	return product.ProductAPI{}
}
