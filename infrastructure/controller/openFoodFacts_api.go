package controller

import (
	"gaia-api/domain/entity"
	"gaia-api/infrastructure/error/openFoodFacts_api_error"
	"github.com/openfoodfacts/openfoodfacts-go"
)

type OpenFoodFactsAPI struct {
	Client *openfoodfacts.Client
}

func NewOpenFoodFactsAPI() *OpenFoodFactsAPI {
	api := openfoodfacts.NewClient("world", "", "")
	return &OpenFoodFactsAPI{&api}
}

func (openFoodFactsAPI *OpenFoodFactsAPI) retrieveAndMapProduct(barcode string) (entity.Product, error) {
	client := openFoodFactsAPI.Client
	product, err := client.Product(barcode)

	if err != nil {
		return entity.Product{}, openFoodFacts_api_error.ProductNotFoundError{Barcode: barcode}
	}

	mappedProduct, err := mapOpenFoodFactsProductToEntitiesProduct(product)

	if err != nil {
		return entity.Product{}, err
	}

	return mappedProduct, nil
}

func mapOpenFoodFactsProductToEntitiesProduct(product *openfoodfacts.Product) (entity.Product, error) {
	nutrients := product.Nutriments

	mappedNutriscore := entity.NutriScore{Score: product.Nutriments.NutritionScoreFr100G, Grade: product.NutritionGradeFr}

	mappedNutrients := entity.Nutrients{
		EnergyKj:      nutrients.Energy100G,
		EnergyKcal:    nutrients.EnergyKcal100G,
		Fat:           nutrients.Fat100G,
		SaturatedFat:  nutrients.SaturatedFat100G,
		Carbohydrates: nutrients.Carbohydrates100G,
		Sugar:         nutrients.Sugars100G,
		Fiber:         nutrients.Fiber100G,
		Proteins:      nutrients.Proteins100G,
		Salt:          nutrients.Salt100G,
	}

	mappedProduct := entity.Product{
		Name:       product.ProductName,
		Nutrients:  mappedNutrients,
		ImageURL:   product.ImageURL.URL.String(),
		NutriScore: mappedNutriscore,
		EcoScore:   product.EcoscoreGrade,
	}

	return mappedProduct, nil
}
