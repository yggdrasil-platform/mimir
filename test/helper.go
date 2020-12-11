package test

import (
	"net/http"
)

type Spy struct {
	Called bool
	NumOfTimes int
}

func CreateNextHandler(s *Spy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Called = true
		s.NumOfTimes++
	}
}
