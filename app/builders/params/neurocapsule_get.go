package params

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/dto"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/generated/server/parameters"
)

func BuildFromNeurocapsuleGetParams(params parameters.NeurocapsuleGetParams) *dto.GetNeurocapsule {
	return &dto.GetNeurocapsule{
		PriceSegment:    params.PriceSegment,
		ColorScheme:     params.ColorScheme,
		ColorBase:       params.ColorBase,
		ExcludedSkuList: params.ExcludedSkuList,
	}
}
