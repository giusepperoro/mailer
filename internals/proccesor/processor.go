package proccesor

import (
	"context"
	"net/http"

	"github.com/giusepperoro/requestqueue/internals/entity"
	"github.com/giusepperoro/requestqueue/internals/response"
	"github.com/giusepperoro/requestqueue/internals/workerpool"
)

type Processor interface {
	Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64)
	Close()
}

type Pros struct {
	taskMap  map[int64]entity.Queue
	isClosed bool
	work     workerpool.Worker
	sender   response.ResponseSender
}
