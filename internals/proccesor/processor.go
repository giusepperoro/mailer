package proccesor

import (
	"context"
	"github.com/giusepperoro/mailer/internals/workerpool"
	"net/http"
)

type Processor interface {
	Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64)
	Close()
}

type Pros struct {
	tuskMap  map[int64]chan Task
	isClosed bool
	work     workerpool.Worker
}

type Task struct {
	Ctx      context.Context
	ClientId int64
	Amount   int64
	W        http.ResponseWriter
}
