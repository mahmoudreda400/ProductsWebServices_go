package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	productService ProductService
}

func ProvideProductAPI(ps ProductService) ProductAPI {
	return ProductAPI{productService: ps}
}

func (api *ProductAPI) FindAll(c *gin.Context) {
	products := api.productService.findAll()
	c.JSON(http.StatusOK, gin.H{"products": toProductDTOs(products)})
}

func (api *ProductAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := api.productService.findByID(id)
	c.JSON(http.StatusFound, gin.H{"product": toProductDTO(product)})
}

func (api *ProductAPI) Save(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	createdProduct := api.productService.save(toProduct(productDTO))
	c.JSON(http.StatusCreated, gin.H{"product": createdProduct})
}

func (api *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := api.productService.findByID(id)
	if product == (Product{}) {
		log.Fatalln("product not found to update")
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	api.productService.save(product)
	c.Status(http.StatusOK)
}

func (api *ProductAPI) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	product := api.productService.findByID(id)
	if product == (Product{}) {
		log.Fatalln(" product not find to delete")
		c.Status(http.StatusBadRequest)
		return
	}

	api.productService.delete(product)
	c.Status(http.StatusOK)
}
