package convert

import (
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/response"
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
