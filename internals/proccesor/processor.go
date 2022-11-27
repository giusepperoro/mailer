package proccesor

import (
	"context"
	"github.com/giusepperoro/mailer/internals/entity"
	"github.com/giusepperoro/mailer/internals/workerpool"
	"net/http"
)

type Processor interface {
	Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64)
	Close()
}

type Pros struct {
	taskMap  map[int64]chan entity.Task
	isClosed bool
	work     workerpool.Worker
}
