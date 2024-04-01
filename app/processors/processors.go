package processors

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/processors/neurocapsule"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/repositories"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/infra/storage/postgres"
)

type Processors struct {
	Neurocapsule neurocapsule.Processor
}

func New(
	storage *postgres.DB,
	storage2 *postgres.DB,
	repository *repositories.Repository,
) *Processors {
	return &Processors{
		Neurocapsule: neurocapsule.New(storage, storage2, repository),
	}
}
