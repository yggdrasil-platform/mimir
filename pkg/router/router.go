package router

import (
  "github.com/gorilla/mux"
	"github.com/kieranroneill/new-go-service-template/pkg/application"
	"github.com/kieranroneill/new-go-service-template/pkg/handler"
	"github.com/kieranroneill/new-go-service-template/pkg/middleware"
  "net/http"
)

func New(app *application.Application) *mux.Router {
	r := mux.NewRouter()

	hlcksr := r.PathPrefix("/healthcheck").Subrouter()
  hlcksr.
		HandleFunc(
			"",
			middleware.ApplyMiddleware(
				handler.CreateHealthcheckHandler(app),
				middleware.LogRequest(),
			),
		).
		Methods(http.MethodGet)

	return r
}
