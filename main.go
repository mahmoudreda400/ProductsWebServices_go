package main

import (
	"fmt"
	"os"

	"ProductsWebServices/product"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	// db, err = gorm.Open(“mysql”, “user:pass@tcp(127.0.0.1:3306)/samples?charset=utf8&parseTime=True&loc=Local”)
	db, err := gorm.Open("mysql", os.Getenv(""))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&product.Product{})
	return db
}

func main() {
	fmt.Println("hello...")
	db := initDB()
	defer db.Close()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products/save", productAPI.Save)
	r.PUT("/products/update/:id", productAPI.Update)
	r.DELETE("products/delete/:id", productAPI.Delete)

	err := r.Run()
	// err := r.Run(“:8080”)
	if err != nil {
		panic(err)
	}

}
