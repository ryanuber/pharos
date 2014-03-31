package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/ryanuber/pharos/pharos"
)

type VersionCommand struct {
	Ui        cli.Ui
	GitCommit string
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintf("Pharos v%s", pharos.Version))
	c.Ui.Output(fmt.Sprintf("Git revision %s", c.GitCommit))
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the Pharos version number"
}
