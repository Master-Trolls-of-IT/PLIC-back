package product

import (
	"encoding/json"
	"gaia-api/domain/entity/response"
	"github.com/openfoodfacts/openfoodfacts-go"
	"golang.org/x/exp/slices"
	"io"
	"net/http"
	"strconv"
)

type OpenFoodFactsAPI struct {
	Client *openfoodfacts.Client
}

func NewOpenFoodFactsAPI() *OpenFoodFactsAPI {
	api := openfoodfacts.NewClient("world", "", "")
	return &OpenFoodFactsAPI{&api}
}

type Response struct {
	Product struct {
		EcoscoreScore float64 `json:"ecoscore_score"`
	} `json:"product"`
}

func getEcoScore(barcode string) (string, error) {
	response, err := http.Get("https://world.openfoodfacts.org/api/v0/product/" + barcode)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	var responseJSON Response
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return "", err
	}

	ecoscoreString := strconv.FormatFloat(responseJSON.Product.EcoscoreScore, 'f', -1, 64)
	return ecoscoreString, nil
}

func (openFoodFactsAPI *OpenFoodFactsAPI) retrieveAndMapProduct(barcode string) (response.Product, error) {
	client := openFoodFactsAPI.Client
	product, err := client.Product(barcode)
	if err != nil {
		return response.Product{}, err
	}

	mappedProduct, err := mapOpenFoodFactsProductToEntitiesProduct(product)
	if err != nil {
		return response.Product{}, err
	}
	mappedProduct.Barcode = barcode
	mappedProduct.EcoScore, err = getEcoScore(barcode)
	if err != nil {
		return response.Product{}, err
	}

	return mappedProduct, nil
}

func mapOpenFoodFactsProductToEntitiesProduct(product *openfoodfacts.Product) (response.Product, error) {
	nutrients := product.Nutriments
	isWater := false
	if slices.Contains(product.CategoriesTags, "en:waters") {
		isWater = true
	}

	mappedNutriscore := response.NutriScore{Score: 5, Grade: product.NutritionGradeFr}

	mappedNutrients := response.Nutrients{
		EnergyKj:      nutrients.Energy,
		EnergyKcal:    nutrients.EnergyKcal,
		Fat:           nutrients.Fat,
		SaturatedFat:  nutrients.SaturatedFat,
		Carbohydrates: nutrients.Carbohydrates,
		Sugar:         nutrients.Sugars,
		Fiber:         nutrients.Fiber,
		Proteins:      nutrients.Proteins,
		Salt:          nutrients.Salt,
	}

	mappedProduct := response.Product{
		Brand:           product.Brands,
		Name:            product.ProductName,
		Nutrients:       mappedNutrients,
		ImageURL:        product.ImageFrontURL.URL.String(),
		NutriScore:      mappedNutriscore,
		EcoScore:        product.EcoscoreGrade,
		IsWater:         isWater,
		Quantity:        product.Quantity,
		ServingQuantity: product.ServingQuantity,
		ServingSize:     product.ServingSize,
	}

	return mappedProduct, nil
}
