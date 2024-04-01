package models

type Product struct {
	SKU         string `db:"sku"`
	FullSKU     string `db:"full_sku"`
	Name        string `db:"name"`
	PriceAmount int64  `db:"price_amount"`
	Thumbnail   string `db:"thumbnail"`
	Gallery     string `db:"gallery"`
	Color       string `db:"colors"`
}
