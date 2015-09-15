package commands

import (
	"fmt"
	"strconv"

	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/goodeggs/ccron/ccron/client"
	"github.com/goodeggs/ccron/ccron/stdcli"
)

var DeleteCommand = cli.Command{
	Name:        "delete",
	Usage:       "delete a scheduled task",
	Description: "ccron delete 1",
	Action:      cmdCronDelete,
	Flags:       []cli.Flag{stdcli.AppFlag, stdcli.ServerFlag},
}

func cmdCronDelete(c *cli.Context) {
	_, app, err := stdcli.DirApp(c, ".")

	if err != nil {
		stdcli.Die(err)
	}

	server := stdcli.Server(c)

	id, err := strconv.Atoi(c.Args().First())

	if err != nil {
		stdcli.Die(err)
	}

	path := fmt.Sprintf("/apps/%s/jobs/%d", app, id)

	_, err = client.Request("DELETE", server, path, nil)

	if err != nil {
		stdcli.Die(err)
	}

	fmt.Printf("Deleted #%d\n", id)
}
