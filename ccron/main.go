package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/goodeggs/ccron/ccron/client"
	"github.com/goodeggs/ccron/ccron/commands"
	"github.com/goodeggs/ccron/ccron/stdcli"
)

var globalServerFlag = cli.StringFlag{
	Name:   "server",
	Usage:  "ccron server URL.",
	Value:  "http://localhost:5000",
	EnvVar: "CCRON_SERVER",
}

func main() {
	app := cli.NewApp()
	app.Name = "ccron"
	app.Usage = "manage an app's scheduled tasks"
	app.Action = cmdCron
	app.Flags = []cli.Flag{stdcli.AppFlag, globalServerFlag}
	app.Commands = []cli.Command{
		commands.CreateCommand,
		commands.DeleteCommand,
		commands.CrontabCommand,
	}
	app.Run(os.Args)
}

func cmdCron(c *cli.Context) {
	_, app, err := stdcli.DirApp(c, ".")

	if err != nil {
		stdcli.Die(err)
	}

	server := stdcli.Server(c)
	path := fmt.Sprintf("/apps/%s/jobs", app)

	body, err := client.Request("GET", server, path, nil)

	if err != nil {
		stdcli.Die(err)
	}

	var jobs stdcli.Jobs
	err = json.Unmarshal(body, &jobs)

	if err != nil {
		stdcli.Die(err)
	}

	stdcli.PrintJobs(jobs...)
}
