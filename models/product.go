package models

type Product struct {
	ID              int32  `id:json`
	product_id      string `product_id:json`
	product_name    string `product_name:json`
	product_desc    string `product_desc:json`
	product_img_url string `product_img_url:json`
}
