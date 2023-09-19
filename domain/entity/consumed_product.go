package entity

type ConsumedProduct struct {
	Product  Product `json:"product"`
	Quantity int
}
