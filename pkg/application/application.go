package application

import (
	"github.com/kieranroneill/new-go-service-template/pkg/config"
)

type Application struct {
	Config *config.Config
}

func New() (*Application, error) {
	return &Application{
		Config: config.New(),
	}, nil
}
