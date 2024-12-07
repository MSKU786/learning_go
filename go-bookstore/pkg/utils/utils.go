package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // Return the error to the caller
	}
	
	if err := json.Unmarshal(body, x); err != nil {
		return err // Return the error to the caller
	}
	
	return nil
}
