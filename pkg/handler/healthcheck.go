package handler

import (
	"github.com/kieranroneill/new-go-service-template/pkg/application"
	_error "github.com/kieranroneill/new-go-service-template/pkg/error"
	"github.com/kieranroneill/new-go-service-template/pkg/server"
  "net/http"
)

func healthcheck(app *application.Application, w http.ResponseWriter, r *http.Request) {
	server.WriteJsonResponse(w, http.StatusOK, server.HealthcheckResponseBody{
		Environment: app.Config.Environment,
		Name: app.Config.ServiceName,
		Version: app.Config.Version,
	})
}

func CreateHealthcheckHandler(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			healthcheck(app, w, r)
			break
		default:
			server.WriteJsonResponse(w, http.StatusMethodNotAllowed, server.HttpErrorResponse{
				Code: _error.MethodNotAllowed,
				Message: _error.GetErrMessage(_error.MethodNotAllowed),
			})
			break
		}
	}
}
