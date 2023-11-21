package request

type ConsumedProductAdd struct {
	Email    string `json:"email"`
	Barcode  string `json:"barcode"`
	Quantity string `json:"quantity"`
}
