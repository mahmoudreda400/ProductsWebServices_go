package product

import "github.com/jinzhu/gorm"

type product struct {
	gorm.Model

	code  string
	price float32
}
