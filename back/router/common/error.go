package common

import (
	"net/http"
	"encoding/json"
)

// SendErr envia un mensaje de error
func SendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}