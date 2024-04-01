package responses

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/dto"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/generated/server/models"

	"strings"
)

func BuildNeurocapsuleGetResponse(neurocapsule *dto.Neurocapsule) *models.NeurocapsuleResponse {
	if neurocapsule == nil {
		return nil
	}

	var trousers, jeans, sweatpants, shirt1, shirt2, tShirt1, tShirt2, hoodie, sweater, blazer *models.Product

	if neurocapsule.Bottom.Trousers != nil {
		trousers = &models.Product{
			Sku:         neurocapsule.Bottom.Trousers.SKU,
			FullSku:     neurocapsule.Bottom.Trousers.FullSKU,
			Name:        neurocapsule.Bottom.Trousers.Name,
			PriceAmount: neurocapsule.Bottom.Trousers.PriceAmount,
			Thumbnail:   neurocapsule.Bottom.Trousers.Thumbnail,
			Gallery:     getGallery(neurocapsule.Bottom.Trousers.Gallery),
			Color:       neurocapsule.Bottom.Trousers.Color,
		}
	}

	if neurocapsule.Bottom.Jeans != nil {
		jeans = &models.Product{
			Sku:         neurocapsule.Bottom.Jeans.SKU,
			FullSku:     neurocapsule.Bottom.Jeans.FullSKU,
			Name:        neurocapsule.Bottom.Jeans.Name,
			PriceAmount: neurocapsule.Bottom.Jeans.PriceAmount,
			Thumbnail:   neurocapsule.Bottom.Jeans.Thumbnail,
			Gallery:     getGallery(neurocapsule.Bottom.Jeans.Gallery),
			Color:       neurocapsule.Bottom.Jeans.Color,
		}
	}

	if neurocapsule.Bottom.Sweatpants != nil {
		sweatpants = &models.Product{
			Sku:         neurocapsule.Bottom.Sweatpants.SKU,
			FullSku:     neurocapsule.Bottom.Sweatpants.FullSKU,
			Name:        neurocapsule.Bottom.Sweatpants.Name,
			PriceAmount: neurocapsule.Bottom.Sweatpants.PriceAmount,
			Thumbnail:   neurocapsule.Bottom.Sweatpants.Thumbnail,
			Gallery:     getGallery(neurocapsule.Bottom.Sweatpants.Gallery),
			Color:       neurocapsule.Bottom.Sweatpants.Color,
		}
	}

	if neurocapsule.Top1.Shirt1 != nil {
		shirt1 = &models.Product{
			Sku:         neurocapsule.Top1.Shirt1.SKU,
			FullSku:     neurocapsule.Top1.Shirt1.FullSKU,
			Name:        neurocapsule.Top1.Shirt1.Name,
			PriceAmount: neurocapsule.Top1.Shirt1.PriceAmount,
			Thumbnail:   neurocapsule.Top1.Shirt1.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top1.Shirt1.Gallery),
			Color:       neurocapsule.Top1.Shirt1.Color,
		}
	}

	if neurocapsule.Top1.Shirt2 != nil {
		shirt2 = &models.Product{
			Sku:         neurocapsule.Top1.Shirt2.SKU,
			FullSku:     neurocapsule.Top1.Shirt2.FullSKU,
			Name:        neurocapsule.Top1.Shirt2.Name,
			PriceAmount: neurocapsule.Top1.Shirt2.PriceAmount,
			Thumbnail:   neurocapsule.Top1.Shirt2.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top1.Shirt2.Gallery),
			Color:       neurocapsule.Top1.Shirt2.Color,
		}
	}

	if neurocapsule.Top1.TShirt1 != nil {
		tShirt1 = &models.Product{
			Sku:         neurocapsule.Top1.TShirt1.SKU,
			FullSku:     neurocapsule.Top1.TShirt1.FullSKU,
			Name:        neurocapsule.Top1.TShirt1.Name,
			PriceAmount: neurocapsule.Top1.TShirt1.PriceAmount,
			Thumbnail:   neurocapsule.Top1.TShirt1.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top1.TShirt1.Gallery),
			Color:       neurocapsule.Top1.TShirt1.Color,
		}
	}

	if neurocapsule.Top1.TShirt2 != nil {
		tShirt2 = &models.Product{
			Sku:         neurocapsule.Top1.TShirt2.SKU,
			FullSku:     neurocapsule.Top1.TShirt2.FullSKU,
			Name:        neurocapsule.Top1.TShirt2.Name,
			PriceAmount: neurocapsule.Top1.TShirt2.PriceAmount,
			Thumbnail:   neurocapsule.Top1.TShirt2.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top1.TShirt2.Gallery),
			Color:       neurocapsule.Top1.TShirt2.Color,
		}
	}

	if neurocapsule.Top2.Hoodie != nil {
		hoodie = &models.Product{
			Sku:         neurocapsule.Top2.Hoodie.SKU,
			FullSku:     neurocapsule.Top2.Hoodie.FullSKU,
			Name:        neurocapsule.Top2.Hoodie.Name,
			PriceAmount: neurocapsule.Top2.Hoodie.PriceAmount,
			Thumbnail:   neurocapsule.Top2.Hoodie.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top2.Hoodie.Gallery),
			Color:       neurocapsule.Top2.Hoodie.Color,
		}
	}

	if neurocapsule.Top2.Sweater != nil {
		sweater = &models.Product{
			Sku:         neurocapsule.Top2.Sweater.SKU,
			FullSku:     neurocapsule.Top2.Sweater.FullSKU,
			Name:        neurocapsule.Top2.Sweater.Name,
			PriceAmount: neurocapsule.Top2.Sweater.PriceAmount,
			Thumbnail:   neurocapsule.Top2.Sweater.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top2.Sweater.Gallery),
			Color:       neurocapsule.Top2.Sweater.Color,
		}
	}

	if neurocapsule.Top2.Blazer != nil {
		blazer = &models.Product{
			Sku:         neurocapsule.Top2.Blazer.SKU,
			FullSku:     neurocapsule.Top2.Blazer.FullSKU,
			Name:        neurocapsule.Top2.Blazer.Name,
			PriceAmount: neurocapsule.Top2.Blazer.PriceAmount,
			Thumbnail:   neurocapsule.Top2.Blazer.Thumbnail,
			Gallery:     getGallery(neurocapsule.Top2.Blazer.Gallery),
			Color:       neurocapsule.Top2.Blazer.Color,
		}
	}

	return &models.NeurocapsuleResponse{
		Neurocapsule: &models.NeurocapsuleResponseNeurocapsule{
			Bottom: &models.NeurocapsuleResponseNeurocapsuleBottom{
				Trousers:   trousers,
				Jeans:      jeans,
				Sweatpants: sweatpants,
			},
			Top1: &models.NeurocapsuleResponseNeurocapsuleTop1{
				Shirt1:  shirt1,
				Shirt2:  shirt2,
				TShirt1: tShirt1,
				TShirt2: tShirt2,
			},
			Top2: &models.NeurocapsuleResponseNeurocapsuleTop2{
				Hoodie:  hoodie,
				Sweater: sweater,
				Blazer:  blazer,
			},
		},
	}
}

func getGallery(g string) string {
	gallery := strings.Split(g, ",")[0]
	return strings.Trim(strings.Trim(gallery, "{"), "}")
}
