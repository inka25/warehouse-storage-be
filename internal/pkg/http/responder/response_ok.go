package responder

import (
	"encoding/json"
	"net/http"
)

func ResponseOK(rw http.ResponseWriter, body interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	return json.NewEncoder(rw).Encode(body)
}
