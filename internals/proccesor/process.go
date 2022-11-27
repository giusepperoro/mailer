package proccesor

import (
	"context"
	"github.com/giusepperoro/mailer/internals/entity"
	"github.com/giusepperoro/mailer/internals/workerpool"
	"net/http"
)

const defaultSize = 10

func NewProcessor(work workerpool.Worker) *Pros {
	return &Pros{
		taskMap: make(map[int64]chan entity.Task),
		work:    work,
	}
}

func (p *Pros) Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64) {
	if p.isClosed {
		return
	}
	ch, exists := p.taskMap[clientId]
	if !exists {
		ch = make(chan entity.Task, defaultSize)
		p.taskMap[clientId] = ch
		p.work.Add(p.taskMap[clientId])
	}
	if p.isClosed {
		return
	}
	ch <- entity.Task{
		Ctx:      ctx,
		ClientId: clientId,
		Amount:   amount,
		W:        w,
	}
}

func (p *Pros) Close() {
	p.isClosed = true
	for _, ch := range p.taskMap {
		close(ch)
	}
}
