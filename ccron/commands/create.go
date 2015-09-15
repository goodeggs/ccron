package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/kballard/go-shellquote"
	"github.com/goodeggs/ccron/ccron/client"
	"github.com/goodeggs/ccron/ccron/stdcli"
)

var CreateCommand = cli.Command{
	Name:        "create",
	Usage:       "create a new scheduled task",
	Description: "ccron create '0 */5 * * * echo hello world'",
	Action:      cmdCronCreate,
	Flags:       []cli.Flag{stdcli.AppFlag, stdcli.ServerFlag},
}

func cmdCronCreate(c *cli.Context) {
	_, app, err := stdcli.DirApp(c, ".")

	if err != nil {
		stdcli.Die(err)
	}

	server := stdcli.Server(c)
	path := fmt.Sprintf("/apps/%s/jobs", app)

	args, err := shellquote.Split(c.Args().First())

	if err != nil {
		stdcli.Die(err)
	}

	if len(args) < 6 {
		stdcli.DieWithUsage(c, fmt.Errorf("argument is not in crontab format: `%s`", c.Args().First()))
	}

	job := stdcli.Job{
		Schedule: strings.Join(args[0:5], " "),
		Command:  strings.Join(args[5:], " "),
		App:      app,
	}

	data, err := json.Marshal(job)

	if err != nil {
		stdcli.Die(err)
	}

	body, err := client.Request("POST", server, path, bytes.NewReader(data))

	if err != nil {
		stdcli.Die(err)
	}

	err = json.Unmarshal(body, &job)

	fmt.Printf("Created #%d\n", job.Id)
	stdcli.PrintJobs(job)
}
