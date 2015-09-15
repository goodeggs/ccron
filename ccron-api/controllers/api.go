package controllers

import (
	"net/http"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/ddollar/logger"
)

type ApiHandlerFunc func(http.ResponseWriter, *http.Request) error

func api(at string, handler ApiHandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log := logger.New("ns=ccron-api").At(at).Start()

		err := handler(rw, r)

		if err != nil {
			log.Error(err)
			RenderError(rw, err)
			return
		}

		log.Log("state=success")
	}
}
