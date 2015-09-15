package commands

import (
	"fmt"

	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/goodeggs/ccron/ccron/client"
	"github.com/goodeggs/ccron/ccron/stdcli"
)

var CrontabCommand = cli.Command{
	Name:   "crontab",
	Usage:  "Prints the current global crontab.",
	Action: cmdCrontab,
	Flags:  []cli.Flag{stdcli.AppFlag, stdcli.ServerFlag},
}

func cmdCrontab(c *cli.Context) {
	server := stdcli.Server(c)
	path := fmt.Sprintf("/crontab")

	body, err := client.Request("GET", server, path, nil)

	if err != nil {
		stdcli.Die(err)
	}

	fmt.Printf(string(body))
}
