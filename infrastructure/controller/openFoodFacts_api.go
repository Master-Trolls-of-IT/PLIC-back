package controller

import (
	"gaia-api/domain/entity"
	"gaia-api/infrastructure/error/openFoodFacts_api_error"
	"github.com/openfoodfacts/openfoodfacts-go"
	"golang.org/x/exp/slices"
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
	isWater := false
	if slices.Contains(product.CategoriesTags, "en:waters") {
		isWater = true
	}

	mappedNutriscore := entity.NutriScore{Score: product.Nutriments.NutritionScoreFr100G, Grade: product.NutritionGradeFr}

	mappedNutrients := entity.Nutrients{
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

	mappedNutrients100g := entity.Nutrients100g{
		entity.Nutrients{
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

	mappedNutrientsValue := entity.NutrientsValue{
		entity.Nutrients{
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

	mappedNutrientsServing := entity.NutrientsServing{
		entity.Nutrients{
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

	//mappedNutrientsUnit := entity.NutrientsUnit{
	//	EnergyKj:      nutrients.EnergyUnit,
	//	EnergyKcal:    nutrients.EnergyKcalUnit,
	//	Fat:           nutrients.FatUnit,
	//	SaturatedFat:  nutrients.SaturatedFatUnit,
	//	Carbohydrates: nutrients.CarbohydratesUnit,
	//	Sugar:         nutrients.SugarsUnit,
	//	Fiber:         nutrients.FiberUnit,
	//	Proteins:      nutrients.ProteinsUnit,
	//	Salt:          nutrients.SaltUnit,
	//}

	mappedProduct := entity.Product{
		Brand:            product.Brands,
		Name:             product.ProductName,
		Nutrients:        mappedNutrients,
		Nutrients100g:    mappedNutrients100g,
		NutrientsServing: mappedNutrientsServing,
		NutrientsValue:   mappedNutrientsValue,
		//NutrientsUnit:    mappedNutrientsUnit,
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
