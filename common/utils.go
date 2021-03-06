package common

import "net/http"
import "encoding/json"
import "log"

type appErr struct {
	Error string "json:'error'"
	Status int "json:'status'"
}

func JsonError(w http.ResponseWriter, handleErr error, code int) {
	log.Printf("Error: %s\n", handleErr)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if res, err := json.Marshal(&appErr{handleErr.Error(), code}); err == nil {
		w.Write(res)
	}
}

func JsonValid(w http.ResponseWriter, res []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(res)
}

func JsonStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}