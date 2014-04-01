package command

import (
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"strings"
)

type StartCommand struct {
	Ui cli.Ui
}

func (c *StartCommand) Help() string {
	helpText := `
Usage: pharos start [options] ...

  Starts the Pharos daemon

Options:
  -bind-addr=0.0.0.0:22322  The bind address for the Pharos daemon
`
	return strings.TrimSpace(helpText)
}

func (c *StartCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("start", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		c.Ui.Output(c.Help())
	}

	var configFile string

	cmdFlags.StringVar(&configFile, "config-file", "", "path to config file")
	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Output(c.Help())
		return 1
	}

	c.Ui.Output(fmt.Sprintf("Config file is %s", configFile))

	return 0
}

func (c *StartCommand) Synopsis() string {
	return "Starts the Pharos daemon"
}
