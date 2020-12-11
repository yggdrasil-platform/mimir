package handler

import (
	"encoding/json"
  "github.com/kieranroneill/new-go-service-template/pkg/application"
  "github.com/kieranroneill/new-go-service-template/pkg/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

)

func Test_healthcheck(t *testing.T) {
	const url = "/healthcheck"
	app, err := application.New()
	if err != nil {
		t.Error(err.Error())
	}

	t.Run("should return 200", func(t *testing.T) {
		var response server.HealthcheckResponseBody
		r, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Errorf("error creating request: %v", err)
		}

		w := httptest.NewRecorder()

		CreateHealthcheckHandler(app).ServeHTTP(w, r)

		if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
			t.Errorf("error encoding request body: %v", err)
		}

		assert.Equal(t, http.StatusOK, w.Code, "expected status code: %v, got status code: %v", http.StatusOK, w.Code)
	})
}
