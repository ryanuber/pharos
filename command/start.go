package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/ryanuber/pharos/pharos"
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
	c.Ui.Output(fmt.Sprintf("Pharos version %s", pharos.Version))
	return 0
}

func (c *StartCommand) Synopsis() string {
	return "Starts the Pharos daemon"
}
