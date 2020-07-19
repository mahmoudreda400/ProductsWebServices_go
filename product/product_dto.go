package product

type ProductDTO struct {
	ID    uint    `json:"id,string,omitEmpty"`
	Code  string  `json:"code"`
	Price float32 `json:"price,string"`
}
