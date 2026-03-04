package dto

type ProductsDto struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Rating      float32 `json:"rating"`
}
