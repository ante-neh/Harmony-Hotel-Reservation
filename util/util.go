package util

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) error{
	response, err := json.Marshal(payload)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("Server Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil 
}


func ResponseWithError(w http.ResponseWriter, code int, message string ){
	ResponseWithJson(w, code, map[string]string{"error":message})
}