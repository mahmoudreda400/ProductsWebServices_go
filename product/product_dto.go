package product

type ProdutDTO struct {
	ID    int     `json:"id,string,omitEmpty"`
	Code  string  `json:"code"`
	Price float32 `json:"price,string"`
}
