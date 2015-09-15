package models

import (
	"fmt"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/gorhill/cronexpr"
	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/kballard/go-shellquote"
)

type Job struct {
	Id       int    `json:"id"`
	App      string `json:"app"`
	Schedule string `json:"schedule"`
	Command  string `json:"command"`
}

func (j *Job) Validate() error {

	if _, err := cronexpr.Parse(j.Schedule); err != nil {
		return fmt.Errorf("Schedule is not valid: %s", err.Error())
	}

	if _, err := shellquote.Split(j.Command); err != nil {
		return fmt.Errorf("Command is not valid: %s", err.Error())
	}

	return nil
}

type Jobs []Job

func ListAllJobs() (Jobs, error) {
	rows, err := Db.Query("SELECT id, app, schedule, command FROM jobs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := Jobs{}

	for rows.Next() {
		job := Job{}
		rows.Scan(&job.Id, &job.App, &job.Schedule, &job.Command)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func ListJobs(app string) (Jobs, error) {
	rows, err := Db.Query("SELECT id, schedule, command FROM jobs WHERE app = $1", app)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := Jobs{}

	for rows.Next() {
		job := Job{
			App: app,
		}
		rows.Scan(&job.Id, &job.Schedule, &job.Command)
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func CreateJob(job *Job) error {
	err := job.Validate()

	if err != nil {
		return err
	}

	return Db.QueryRow(`INSERT INTO jobs(app, schedule, command) VALUES ($1, $2, $3) RETURNING id`, job.App, job.Schedule, job.Command).Scan(&job.Id)
}

func DeleteJob(app string, id int) error {
	res, err := Db.Exec("DELETE FROM jobs where app = $1 and id = $2", app, id)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count < 1 {
		return fmt.Errorf("job not found")
	}

	return nil
}
