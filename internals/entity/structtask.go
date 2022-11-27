package entity

import (
	"context"
	"net/http"
)

type Task struct {
	Ctx      context.Context
	ClientId int64
	Amount   int64
	W        http.ResponseWriter
}
