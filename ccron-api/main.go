package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/ddollar/logger"
	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/ddollar/nlogger"
	"github.com/goodeggs/ccron/ccron-api/controllers"
	"github.com/goodeggs/ccron/ccron-api/models"
)

func recoverWith(f func(err error)) {
	if r := recover(); r != nil {
		// coerce r to error type
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}

		f(err)
	}
}

func recovery(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer recoverWith(func(err error) {
		log := logger.New("ns=ccron-api").At("panic")
		log.Error(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	})

	next(rw, r)
}

var port string = "5000"

func main() {
	err := models.Connect()

	if err != nil {
		panic(err)
	}

	defer models.Disconnect()

	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	n := negroni.New()

	n.Use(negroni.HandlerFunc(recovery))
	n.Use(nlogger.New("ns=ccron-api", nil))

	n.UseHandler(controllers.NewRouter())

	n.Run(fmt.Sprintf(":%s", port))
}
