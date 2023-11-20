package request

type ConsumedProductUpdateQuantity struct {
	UserEmail string `json:"email"`
	Barcode   string `json:"barcode"`
	Quantity  int    `json:"quantity"`
}
