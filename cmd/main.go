package main

import (
	"context"
	"github.com/giusepperoro/mailer/internals/database"
	"github.com/giusepperoro/mailer/internals/handlers"
	"github.com/giusepperoro/mailer/internals/proccesor"
	"github.com/giusepperoro/mailer/internals/transactionresponse"
	"github.com/giusepperoro/mailer/internals/workerpool"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	db, err := database.New(ctx)
	if err != nil {
		log.Fatal("database connect error")
	}
	sender := transactionresponse.NewSender()
	wr := workerpool.NewWorker(db, sender)
	pr := proccesor.NewProcces(wr)
	http.HandleFunc("/form", handlers.HandleBalanceChanger(pr))
	err = http.ListenAndServe("0.0.0.0:80", nil)
	log.Fatal("err here`:", err)
}
