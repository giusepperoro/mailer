package handlers

import (
	"encoding/json"
	"github.com/giusepperoro/requestqueue/internals/proccesor"
	"io"
	"net/http"
)

type BalanceRequest struct {
	ClientId int64 `json:"client_id"`
	Amount   int64 `json:"amount"`
}

type BalanceResponse struct {
	Approved bool   `json:"approved"`
	Err      string `json:"error,omitempty"`
}

func HandleBalanceChanger(manager proccesor.Processor) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		req := new(BalanceRequest)

		if request.Method != "POST" {
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, req)
		if err != nil {
			return
		}
		ctx := request.Context()
		manager.Process(ctx, writer, req.ClientId, req.Amount)
	}
}
