package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/giusepperoro/requestqueue/internals/database"
	"github.com/giusepperoro/requestqueue/internals/handlers"
	"github.com/giusepperoro/requestqueue/internals/proccesor"
	"github.com/giusepperoro/requestqueue/internals/response"
	"github.com/giusepperoro/requestqueue/internals/workerpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db, err := database.New(ctx)
	if err != nil {
		log.Println(err)
		log.Fatal("database connect error")
	}
	sender := response.NewSender()
	wr := workerpool.NewWorker(db)
	pr := proccesor.NewProcessor(wr, sender)
	http.HandleFunc("/form", handlers.HandleBalanceChanger(pr))
	err = http.ListenAndServe("0.0.0.0:80", nil)
	log.Fatal("err here`:", err)
}
