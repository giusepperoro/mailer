package workerpool

import (
	"github.com/giusepperoro/mailer/internals/database"
	"github.com/giusepperoro/mailer/internals/proccesor"
	"github.com/giusepperoro/mailer/internals/transactionresponse"
	"log"
)

type Worker interface {
	Add(ch chan proccesor.Task)
}

type Work struct {
	db database.DbManager
	sr transactionresponse.ResponseSender
}

func NewWorker(db database.DbManager, sr transactionresponse.ResponseSender) *Work {
	return &Work{
		db: db,
		sr: sr,
	}
}

func (w *Work) Add(ch chan proccesor.Task) {
	go func() {
		for task := range ch {
			answer, err := w.db.ChangeBalance(task.Ctx, task.ClientId, task.Amount)
			w.sr.SendResponse(task.W, answer)
			log.Println(err)
		}
	}()
}
