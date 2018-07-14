package shared

import (
	"net/http"
	"encoding/json"
)

type BaseController struct {
	JsonParser
}

func (c BaseController) RespondWithJson(w http.ResponseWriter, payload interface{}, code int) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
