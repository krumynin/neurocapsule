package neurocapsule

import (
	"context"

	"stash.lamoda.ru/neurocapsule/neurocapsule/app/dto"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/repositories"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/infra/storage/postgres"
)

type Processor interface {
	GetNeurocapsule(ctx context.Context, params *dto.GetNeurocapsule) (*dto.Neurocapsule, error)
}

type neurocapsuleProcessor struct {
	storage    *postgres.DB
	storage2   *postgres.DB
	repository *repositories.Repository
}

func New(
	storage *postgres.DB,
	storage2 *postgres.DB,
	repository *repositories.Repository,
) Processor {
	return &neurocapsuleProcessor{
		storage:    storage,
		storage2:   storage2,
		repository: repository,
	}
}
