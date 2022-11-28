package workerpool

import (
	"fmt"
	"log"

	"github.com/giusepperoro/requestqueue/internals/database"
	"github.com/giusepperoro/requestqueue/internals/entity"
)

type Worker interface {
	Add(queue entity.Queue)
}

type Work struct {
	db database.DbManager
}

func NewWorker(db database.DbManager) *Work {
	return &Work{
		db: db,
	}
}

func (w *Work) Add(queue entity.Queue) {
	fmt.Println("added...")
	go func(queue entity.Queue) {
		for task := range queue.TaskChan {
			fmt.Println("got some!")
			answer, err := w.db.ChangeBalance(task.Ctx, task.ClientId, task.Amount)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("sending ans back...")
			task.ResultChan <- answer
		}
	}(queue)
}
