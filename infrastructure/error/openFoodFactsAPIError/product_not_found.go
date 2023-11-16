package openFoodFactsAPIError

import "fmt"

type ProductNotFoundError struct {
	Barcode string
}

func (err ProductNotFoundError) Error() string {
	return fmt.Sprintf("Product with barcode: %s not found.", err.Barcode)
}
