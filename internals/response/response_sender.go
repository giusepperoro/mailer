package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseSender interface {
	SendResponse(w http.ResponseWriter, approve bool)
}

type SenderInterface struct {
}

type ResponseSend struct {
	Approved bool `json:"approved"`
	//Err      string `json:"error,omitempty"`
}

func NewSender() *SenderInterface {
	return &SenderInterface{}
}

func (rs *SenderInterface) SendResponse(w http.ResponseWriter, approve bool) {
	response := ResponseSend{
		Approved: approve,
	}

	rawData, err := json.Marshal(response)
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(rawData)
	if err != nil {
		log.Printf("unable to write data: %v", err)
	}
}
