package product

func toProduct(productDTO ProductDTO) Product {
	return Product{Code: productDTO.Code, Price: productDTO.Price}
}

func toProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func toProductDTOs(products []Product) []ProductDTO {
	productDTOs := make([]ProductDTO, len(products))

	for i, item := range products {
		productDTOs[i] = toProductDTO(item)
	}
	return productDTOs
}
