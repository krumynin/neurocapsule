package api

import (
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/api/neurocapsule"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/processors"
	"stash.lamoda.ru/neurocapsule/neurocapsule/app/scratch/server"
)

func New(s *server.Server, p *processors.Processors) {
	neurocapsuleApi := neurocapsule.New(p.Neurocapsule)

	s.Get("/neurocapsule.get", neurocapsuleApi.Get)
}
