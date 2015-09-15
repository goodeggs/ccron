package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/goodeggs/ccron/ccron-api/models"
)

func JobsList(rw http.ResponseWriter, r *http.Request) error {
	app := mux.Vars(r)["app"]

	jobs, err := models.ListJobs(app)

	if err != nil {
		return err
	}

	return RenderJson(rw, jobs)
}

func JobsDelete(rw http.ResponseWriter, r *http.Request) error {
	app := mux.Vars(r)["app"]
	job, err := strconv.Atoi(mux.Vars(r)["job"])

	if err != nil {
		return err
	}

	err = models.DeleteJob(app, job)

	if err != nil {
		return err
	}

	return RenderSuccess(rw)
}

func JobsCreate(rw http.ResponseWriter, r *http.Request) error {
	app := mux.Vars(r)["app"]

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	var job models.Job
	err = json.Unmarshal(body, &job)

	if err != nil {
		return err
	}

	if app != job.App {
		return fmt.Errorf("app name does not match")
	}

	err = models.CreateJob(&job)

	if err != nil {
		return err
	}

	return RenderJson(rw, job)
}
