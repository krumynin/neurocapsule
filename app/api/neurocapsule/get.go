package neurocapsule

import (
	"context"
	neurocapsuleParams "stash.lamoda.ru/neurocapsule/neurocapsule/app/builders/params"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/builders/responses"

	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/generated/server/models"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/generated/server/parameters"
)

func (n *Neurocapsule) Get(ctx context.Context, params parameters.NeurocapsuleGetParams) (*models.NeurocapsuleResponse, error) {
	neurocapsule, err := n.processor.GetNeurocapsule(
		ctx,
		neurocapsuleParams.BuildFromNeurocapsuleGetParams(params),
	)
	if err != nil {
		return nil, err
	}
	return responses.BuildNeurocapsuleGetResponse(neurocapsule), nil
}
