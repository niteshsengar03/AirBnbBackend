package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// init() function is a special-purpose function used for package initialization.
// It is automatically called by the Go runtime before the main() function of the program is executed.
// and regardless of how many time it is called it will run only for one time
var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func ReadJsonBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // prevent unknow fields
	return decoder.Decode(result)
}


func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json") // set content type
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) // encode data as json
}



