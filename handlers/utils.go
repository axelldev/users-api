package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondJSON creates a new Encoder withe the passed writter w
// and send the payload to response.
func RespondJSON(w http.ResponseWriter, payload any) error {
	return json.NewEncoder(w).Encode(payload)
}
