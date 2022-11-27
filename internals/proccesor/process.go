package proccesor

import (
	"context"
	"github.com/giusepperoro/mailer/internals/workerpool"
	"net/http"
)

const defaultSize = 10

func NewProcces(work workerpool.Worker) *Pros {
	return &Pros{
		tuskMap: make(map[int64]chan Task),
		work:    work,
	}
}

func (p *Pros) Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64) {
	if p.isClosed {
		return
	}
	ch, exists := p.tuskMap[clientId]
	if !exists {
		ch = make(chan Task, defaultSize)
		p.tuskMap[clientId] = ch
		p.work.Add(p.tuskMap[clientId])
	}
	if p.isClosed {
		return
	}
	ch <- Task{
		Ctx:      ctx,
		ClientId: clientId,
		Amount:   amount,
		W:        w,
	}
}

func (p *Pros) Close() {
	p.isClosed = true
	for _, ch := range p.tuskMap {
		close(ch)
	}
}
