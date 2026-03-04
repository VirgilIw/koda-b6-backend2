package model

type ProductModel struct {
	Id          int     `db:"id"`
	ProductName string  `db:"product_name"`
	Rating      float32 `db:"rating"`
}
