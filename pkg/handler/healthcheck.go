package handler

import (
	"github.com/kieranroneill/mimir/pkg/application"
	_error "github.com/kieranroneill/mimir/pkg/error"
	"github.com/kieranroneill/mimir/pkg/server"
  "net/http"
)

func healthcheck(app *application.Application, w http.ResponseWriter) {
  var connected = false
  db, err := app.Database.DB()
  if err == nil {
    if err = db.Ping(); err == nil {
      connected = true
    }
  }

	server.WriteJsonResponse(w, http.StatusOK, server.HealthcheckResponseBody{
		Environment: app.Config.Environment,
		IsDatabaseConnected: connected,
		Name: app.Config.ServiceName,
		Version: app.Config.Version,
	})
}

func CreateHealthcheckHandler(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			healthcheck(app, w)
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
