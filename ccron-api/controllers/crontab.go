package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/kballard/go-shellquote"
	"github.com/goodeggs/ccron/ccron-api/models"
)

type RedisConfig struct {
	Host     string
	Port     string
	Database string
	Auth     string
}

func getRedis() (*RedisConfig, error) {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "redis://localhost/"
	}

	parts, err := url.Parse(redisUrl)

	if err != nil {
		return nil, err
	}

	redis := &RedisConfig{}

	hostParts := strings.Split(parts.Host, ":")
	redis.Host = hostParts[0]
	if len(hostParts) == 2 {
		redis.Port = hostParts[1]
	} else {
		redis.Port = "6379"
	}

	if parts.Path == "" || parts.Path == "/" {
		redis.Database = "0"
	} else {
		redis.Database = strings.TrimLeft(parts.Path, "/")
	}

	if parts.User != nil {
		redis.Auth, _ = parts.User.Password()
	}

	return redis, nil
}

func CrontabShow(rw http.ResponseWriter, r *http.Request) error {
	redis, err := getRedis()

	if err != nil {
		return err
	}

	jobs, err := models.ListAllJobs()

	if err != nil {
		return err
	}

	var crontab []byte
	cw := bytes.NewBuffer(crontab)
	fmt.Fprintln(cw, "# autogenerated by ccron at", time.Now().UTC())
	fmt.Fprintf(cw, "CRONLOCK_HOST=%s\n", redis.Host)
	fmt.Fprintf(cw, "CRONLOCK_PORT=%s\n", redis.Port)
	fmt.Fprintf(cw, "CRONLOCK_AUTH=%s\n", redis.Auth)
	fmt.Fprintf(cw, "CRONLOCK_DB=%s\n", redis.Database)
	for _, job := range jobs {
		command := fmt.Sprintf("CRONLOCK_KEY=%s.%d %s", job.App, job.Id, shellquote.Join("cronlock", job.Command))
		fmt.Fprintln(cw, job.Schedule, command)
	}

	return RenderText(rw, cw.String())
}
