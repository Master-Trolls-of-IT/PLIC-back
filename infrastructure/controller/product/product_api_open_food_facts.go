package product

import (
	"encoding/json"
	"gaia-api/domain/entity/convert"
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/response"
	"io"
	"net/http"
)

func (product *Product) RetrieveOpenFoodFactsProduct(barcode string) (*response.Product, error) {
	openFoodFactsResponse, err := http.Get("https://world.openfoodfacts.org/api/v2/product/" + barcode)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(openFoodFactsResponse.Body)

	var responseJSON mapping.OpenFoodFactsProduct
	err = json.NewDecoder(openFoodFactsResponse.Body).Decode(&responseJSON)
	if err != nil {
		return nil, err
	}
	return convert.OpenFoodFactsProductToProduct(responseJSON), nil
}
