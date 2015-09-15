package stdcli

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"text/tabwriter"

	"github.com/goodeggs/ccron/ccron/Godeps/_workspace/src/github.com/codegangsta/cli"
)

var AppFlag = cli.StringFlag{
	Name:  "app",
	Usage: "App name. Inferred from current directory if not specified.",
}

var ServerFlag = cli.StringFlag{
	Name:  "server",
	Usage: "ccron server URL.",
}

// If user specifies the app's name from command line, then use it;
// otherwise use the current working directory's name
// from https://github.com/convox/cli/blob/0c3efc6892cb9a58aec60bd912e060c602f9ab61/stdcli/stdcli.go#L86
func DirApp(c *cli.Context, wd string) (string, string, error) {
	abs, err := filepath.Abs(wd)

	if err != nil {
		return "", "", err
	}

	app := c.String("app")
	if app == "" {
		app = c.GlobalString("app")
	}

	if app == "" {
		app = path.Base(abs)
	}

	return abs, app, nil
}

func Die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}

func DieWithUsage(c *cli.Context, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
	cli.ShowSubcommandHelp(c)
	os.Exit(1)
}

func PrintJobs(jobs ...Job) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Id\tSchedule\tCommand")
	for _, job := range jobs {
		fmt.Fprintf(w, "%d\t%s\t%s\n", job.Id, job.Schedule, job.Command)
	}
	fmt.Fprintln(w)
	w.Flush()
}

type Job struct {
	Id       int    `json:"id"`
	App      string `json:"app"`
	Schedule string `json:"schedule"`
	Command  string `json:"command"`
}

type Jobs []Job

func Server(c *cli.Context) string {
	val := c.String("server")
	if val == "" {
		val = c.GlobalString("server")
	}
	return val
}
