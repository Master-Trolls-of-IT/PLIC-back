package product

import (
	"encoding/json"
	"gaia-api/domain/entity/mapping"
	"github.com/openfoodfacts/openfoodfacts-go"
	"golang.org/x/exp/slices"
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
	defer response.Body.Close()

	var responseJSON Response
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return "", err
	}

	ecoscoreString := strconv.FormatFloat(responseJSON.Product.EcoscoreScore, 'f', -1, 64)
	return ecoscoreString, nil
}

func (openFoodFactsAPI *OpenFoodFactsAPI) retrieveAndMapProduct(barcode string) (mapping.Product, error) {
	client := openFoodFactsAPI.Client
	product, err := client.Product(barcode)
	if err != nil {
		return mapping.Product{}, err
	}

	mappedProduct, err := mapOpenFoodFactsProductToEntitiesProduct(product)
	if err != nil {
		return mapping.Product{}, err
	}
	mappedProduct.EcoScore, err = getEcoScore(barcode)
	if err != nil {
		return mapping.Product{}, err
	}

	return mappedProduct, nil
}

func mapOpenFoodFactsProductToEntitiesProduct(product *openfoodfacts.Product) (mapping.Product, error) {
	nutrients := product.Nutriments
	isWater := false
	if slices.Contains(product.CategoriesTags, "en:waters") {
		isWater = true
	}

	mappedNutriscore := mapping.NutriScore{Score: product.Nutriments.NutritionScoreFr100G, Grade: product.NutritionGradeFr}

	mappedNutrients := mapping.Nutrients{
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

	mappedNutrients100g := mapping.Nutrients100g{
		mapping.Nutrients{
			EnergyKj:      nutrients.Energy100G,
			EnergyKcal:    nutrients.EnergyKcal100G,
			Fat:           nutrients.Fat100G,
			SaturatedFat:  nutrients.SaturatedFat100G,
			Carbohydrates: nutrients.Carbohydrates100G,
			Sugar:         nutrients.Sugars100G,
			Fiber:         nutrients.Fiber100G,
			Proteins:      nutrients.Proteins100G,
			Salt:          nutrients.Salt100G,
		},
	}

	mappedNutrientsValue := mapping.NutrientsValue{
		mapping.Nutrients{
			EnergyKj:      nutrients.EnergyValue,
			EnergyKcal:    nutrients.EnergyKcalValue,
			Fat:           nutrients.FatValue,
			SaturatedFat:  nutrients.SaturatedFatValue,
			Carbohydrates: nutrients.CarbohydratesValue,
			Sugar:         nutrients.SugarsValue,
			Fiber:         nutrients.FiberValue,
			Proteins:      nutrients.ProteinsValue,
			Salt:          nutrients.SaltValue,
		},
	}

	mappedNutrientsServing := mapping.NutrientsServing{
		mapping.Nutrients{
			EnergyKj:      nutrients.EnergyServing,
			EnergyKcal:    nutrients.EnergyKcalServing,
			Fat:           nutrients.FatServing,
			SaturatedFat:  nutrients.SaturatedFatServing,
			Carbohydrates: nutrients.CarbohydratesServing,
			Sugar:         nutrients.SugarsServing,
			Fiber:         nutrients.FiberServing,
			Proteins:      nutrients.ProteinsServing,
			Salt:          nutrients.SaltServing,
		},
	}

	mappedProduct := mapping.Product{
		Brand:            product.Brands,
		Name:             product.ProductName,
		Nutrients:        mappedNutrients,
		Nutrients100g:    mappedNutrients100g,
		NutrientsServing: mappedNutrientsServing,
		NutrientsValue:   mappedNutrientsValue,
		ImageURL:         product.ImageFrontURL.URL.String(),
		NutriScore:       mappedNutriscore,
		EcoScore:         product.EcoscoreGrade,
		IsWater:          isWater,
		Quantity:         product.Quantity,
		ServingQuantity:  product.ServingQuantity,
		ServingSize:      product.ServingSize,
	}

	return mappedProduct, nil
}
