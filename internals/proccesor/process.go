package proccesor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/giusepperoro/requestqueue/internals/entity"
	"github.com/giusepperoro/requestqueue/internals/response"
	"github.com/giusepperoro/requestqueue/internals/workerpool"
)

const defaultSize = 10

func NewProcessor(work workerpool.Worker, sender response.ResponseSender) *Pros {
	return &Pros{
		taskMap: make(map[int64]entity.Queue),
		work:    work,
		sender:  sender,
	}
}

func (p *Pros) Process(ctx context.Context, w http.ResponseWriter, clientId, amount int64) {
	if p.isClosed {
		return
	}
	queue, exists := p.taskMap[clientId]
	if !exists {
		taskChan := make(chan entity.Task, defaultSize)
		queue = entity.Queue{
			TaskChan: taskChan,
		}
		p.taskMap[clientId] = queue
		p.work.Add(p.taskMap[clientId])
	}
	if p.isClosed {
		return
	}
	fmt.Println("sending data...")
	resultChan := make(chan bool)
	queue.TaskChan <- entity.Task{
		ResultChan: resultChan,
		Ctx:        ctx,
		ClientId:   clientId,
		Amount:     amount,
	}
	fmt.Println("ans recieved!")
	res := <-resultChan
	close(resultChan)
	p.sender.SendResponse(w, res)
}

func (p *Pros) Close() {
	p.isClosed = true
	for _, queue := range p.taskMap {
		close(queue.TaskChan)
	}
}
