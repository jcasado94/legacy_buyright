package server

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, code int, msg string) {
	WriteJson(w, code, map[string]string{"error": msg})
}

func WriteJson(w http.ResponseWriter, code int, payload interface{}) {
	rsp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(rsp)
}
