package convert

import (
	"encoding/json"
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/response"
	"golang.org/x/exp/slices"
	"strconv"
)

func ProductMappingToProduct(productMapping mapping.Product) response.Product {
	return response.Product{
		ID:        productMapping.ID,
		Brand:     productMapping.Brand,
		Name:      productMapping.Name,
		Barcode:   productMapping.Barcode,
		Nutrients: response.Nutrients{},
		ImageURL:  productMapping.ImageURL,
		NutriScore: response.NutriScore{
			Score: productMapping.NutriScoreScore,
			Grade: productMapping.NutriScoreGrade,
		},
		EcoScore:        productMapping.EcoScore,
		IsWater:         productMapping.IsWater,
		Quantity:        productMapping.Quantity,
		ServingQuantity: productMapping.ServingQuantity,
		ServingSize:     productMapping.ServingSize,
	}
}

func NutrientsMappingToNutrients(nutrientsMapping mapping.Nutrients) response.Nutrients {
	return response.Nutrients{
		EnergyKj:      nutrientsMapping.EnergyKj,
		EnergyKcal:    nutrientsMapping.EnergyKcal,
		Fat:           nutrientsMapping.Fat,
		SaturatedFat:  nutrientsMapping.SaturatedFat,
		Carbohydrates: nutrientsMapping.Carbohydrates,
		Sugar:         nutrientsMapping.Sugar,
		Fiber:         nutrientsMapping.Fiber,
		Proteins:      nutrientsMapping.Proteins,
		Salt:          nutrientsMapping.Salt,
	}
}

func OpenFoodFactsProductToProduct(openFoodFactsProduct mapping.OpenFoodFactsProduct) *response.Product {
	ecoscore := strconv.FormatFloat(openFoodFactsProduct.Product.Ecoscore, 'f', -1, 64)
	isWater := slices.Contains(openFoodFactsProduct.Product.CategoriesTags, "en:waters")
	servingQuantity := json.Number(openFoodFactsProduct.Product.ServingQuantity)

	nutriScore := response.NutriScore{
		Score: 0,
		Grade: openFoodFactsProduct.Product.Nutriscore,
	}

	return &response.Product{
		Brand:           openFoodFactsProduct.Product.Brand,
		Name:            openFoodFactsProduct.Product.Name,
		Barcode:         openFoodFactsProduct.Product.Barcode,
		Nutrients:       OpenFoodFactsNutrientsToNutrients(openFoodFactsProduct.Product.Nutrients),
		ImageURL:        openFoodFactsProduct.Product.ImageURL,
		NutriScore:      nutriScore,
		EcoScore:        ecoscore,
		IsWater:         isWater,
		Quantity:        "",
		ServingQuantity: servingQuantity,
		ServingSize:     "",
	}
}

func OpenFoodFactsNutrientsToNutrients(openFoodFactsNutrients mapping.OpenFoodFactsNutrients) response.Nutrients {
	return response.Nutrients{
		EnergyKj:      openFoodFactsNutrients.EnergyKj,
		EnergyKcal:    openFoodFactsNutrients.EnergyKcal,
		Fat:           openFoodFactsNutrients.Fat,
		SaturatedFat:  openFoodFactsNutrients.SaturatedFat,
		Carbohydrates: openFoodFactsNutrients.Carbohydrates,
		Sugar:         openFoodFactsNutrients.Sugar,
		Fiber:         openFoodFactsNutrients.Fiber,
		Proteins:      openFoodFactsNutrients.Proteins,
		Salt:          openFoodFactsNutrients.Salt,
	}
}
