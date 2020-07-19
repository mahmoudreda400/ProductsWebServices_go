package product

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	code  string
	price float32
}
