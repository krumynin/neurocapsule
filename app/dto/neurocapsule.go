package dto

import "stash.lamoda.ru/neurocapsule/neurocapsule/app/models"

type Bottom struct {
	Trousers   *models.Product
	Jeans      *models.Product
	Sweatpants *models.Product
}

type Top1 struct {
	Shirt1  *models.Product
	Shirt2  *models.Product
	TShirt1 *models.Product
	TShirt2 *models.Product
}

type Top2 struct {
	Hoodie  *models.Product
	Sweater *models.Product
	Blazer  *models.Product
}

type Neurocapsule struct {
	Bottom Bottom
	Top1   Top1
	Top2   Top2
}

type GetNeurocapsule struct {
	PriceSegment    string
	ColorScheme     string
	ColorBase       string
	ExcludedSkuList *string
}
