package controllers

import (
	"net/http"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/gorilla/mux"
)

func HandlerFunc(w http.ResponseWriter, req *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, req)
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/apps/{app}/jobs", api("jobs.retrieve", JobsList)).Methods("GET")
	router.HandleFunc("/apps/{app}/jobs", api("jobs.create", JobsCreate)).Methods("POST")
	router.HandleFunc("/apps/{app}/jobs/{job}", api("jobs.delete", JobsDelete)).Methods("DELETE")

	router.HandleFunc("/crontab", api("crontab.show", CrontabShow)).Methods("GET")

	// utility
	//router.HandleFunc("/boom", UtilityBoom).Methods("GET")
	//router.HandleFunc("/check", UtilityCheck).Methods("GET")

	return
}
