package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
  "github.com/kieranroneill/new-go-service-template/pkg/logger"
	"net/http"
)

type HealthcheckResponseBody struct {
	Environment string `json:"environment"`
	Name string `json:"name"`
	Version string `json:"version"`
}

type HttpErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	ValidationErrors []string `json:"validationErrors"`
}

func (r *HttpErrorResponse) AppendValidationErrors(errs []validator.FieldError, prx string) *HttpErrorResponse {
	for _, err := range errs {
		if err.Tag() == "required" {
			r.ValidationErrors = append(r.ValidationErrors, fmt.Sprintf("%s '%s' is required", prx, err.Field()))
		}

		if err.Tag() == "email" {
			r.ValidationErrors = append(r.ValidationErrors, fmt.Sprintf("%s '%s' is an invalid email address", prx, err.Field()))
		}
	}

	return r
}

func WriteJsonResponse(w http.ResponseWriter, c int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error.Printf("Failed to create response body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	w.Write(response)
}
