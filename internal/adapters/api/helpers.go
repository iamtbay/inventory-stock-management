package api

import (
	"encoding/json"
	"net/http"
)

func (h *HTTPHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (h *HTTPHandler) writeError(w http.ResponseWriter, status int, message string) {
	type envelope struct {
		Error string `json:"error"`
	}
	h.writeJSON(w, status, &envelope{Error: message})
}

func (h *HTTPHandler) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes :=1048576
	r.Body=http.MaxBytesReader(w,r.Body,int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	if err:=dec.Decode(data);err!=nil{
		return err
	}
	return nil
}