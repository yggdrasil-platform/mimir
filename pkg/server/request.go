package server

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_error "github.com/kieranroneill/mimir/pkg/error"
	"github.com/kieranroneill/mimir/pkg/logger"
  "net/http"
  "reflect"
  "strings"
)

func ParseBody(r *http.Request, b interface{}) *HttpErrorResponse {
	// Make sure te body is not empty.
	if r.ContentLength <= 0 {
		return &HttpErrorResponse{
			Code: _error.RequestBodyEmpty,
			Message: _error.GetErrMessage(_error.RequestBodyEmpty),
		}
	}

	// Parse the body.
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		logger.Error.Print(err)
		return &HttpErrorResponse{
			Code: _error.Unknown,
			Message: err.Error(),
		}
	}

	return nil
}

func ValidateBody(b interface{}) *HttpErrorResponse {
	// Create a validator and convert the field name to the json tag name.
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	})

	// Validate the body.
	if err := validate.Struct(b); err != nil {
		result := HttpErrorResponse{
			Code: _error.RequestBodyInvalid,
			Message: _error.GetErrMessage(_error.RequestBodyInvalid),
		}

		return result.AppendValidationErrors(err.(validator.ValidationErrors), "body:")
	}

	return nil
}

func ValidateQuery(q interface{}) *HttpErrorResponse {
	// Create a validator and convert the field name to the json tag name.
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	})

	// Validate the query.
	if err := validate.Struct(q); err != nil {
		result := HttpErrorResponse{
			Code: _error.RequestQueryInvalid,
			Message: _error.GetErrMessage(_error.RequestQueryInvalid),
		}

		return result.AppendValidationErrors(err.(validator.ValidationErrors), "query:")
	}

	return nil
}
