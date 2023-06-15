package controllers

import (
	"gaia-api/domain/entities"
	"github.com/openfoodfacts/openfoodfacts-go"
)

type OpenFoodFactsAPI struct {
	Client *openfoodfacts.Client
}

func NewOpenFoodFactsAPI() *OpenFoodFactsAPI {
	api := openfoodfacts.NewClient("world", "", "")
	return &OpenFoodFactsAPI{&api}
}

func (openFoodFactsAPI *OpenFoodFactsAPI) retrieveAndMapProduct(barcode string) (entities.Product, error) {
	client := openFoodFactsAPI.Client
	product, err := client.Product(barcode)

	if err != nil {
		return entities.Product{}, err
	}

	mappedProduct, err := mapOpenFoodFactsProductToEntitiesProduct(product)

	if err != nil {
		return entities.Product{}, err
	}

	return mappedProduct, nil
}

func mapOpenFoodFactsProductToEntitiesProduct(product *openfoodfacts.Product) (entities.Product, error) {
	nutrients := product.Nutriments

	mappedNutriscore := entities.NutriScore{Score: product.Nutriments.NutritionScoreFr100G, Grade: product.NutritionGradeFr}

	mappedNutrients := entities.Nutrients{
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

	mappedProduct := entities.Product{
		Name:       product.ProductNameEn,
		Nutrients:  mappedNutrients,
		ImageURL:   product.ImageURL.URL.String(),
		NutriScore: mappedNutriscore,
		EcoScore:   product.EcoscoreGrade,
	}

	return mappedProduct, nil
}
