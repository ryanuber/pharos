package command

import (
	"flag"
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
	cmdFlags := flag.NewFlagSet("start", flag.ContinueOnError)
	cmdFlags.Usage = func() {
		c.Ui.Output(c.Help())
	}

	var bindAddr string

	cmdFlags.StringVar(&bindAddr, "bind-addr", "", "bind address for Pharos")
	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Output(c.Help())
		return 1
	}

	config := pharos.DefaultConfig()
	if bindAddr != "" {
		config.BindAddr = bindAddr
	}

	p := pharos.NewPharos(config)
	c.Ui.Output(fmt.Sprintf("%#v", p))

	return 0
}

func (c *StartCommand) Synopsis() string {
	return "Starts the Pharos daemon"
}
