package neurocapsule

import "stash.lamoda.ru/neurocapsule/neurocapsule/app/processors/neurocapsule"

type Neurocapsule struct {
	processor neurocapsule.Processor
}

func New(processor neurocapsule.Processor) *Neurocapsule {
	return &Neurocapsule{
		processor: processor,
	}
}
