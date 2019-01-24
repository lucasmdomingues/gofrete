package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Log struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (log *Log) Info(w http.ResponseWriter, data interface{}, message string) {

	log.Data = data
	log.Message = message

	json, err := json.Marshal(log)
	if err != nil {
		w.Write([]byte(json))
		return
	}

	w.Write([]byte(json))
}

func (log *Log) Error(w http.ResponseWriter, err error) {
	fmt.Printf("Error: %+v", err)
	http.Error(w, http.StatusText(500), 500)
}
