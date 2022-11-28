package entity

import "context"

type Task struct {
	Ctx        context.Context
	ResultChan chan bool
	ClientId   int64
	Amount     int64
}
