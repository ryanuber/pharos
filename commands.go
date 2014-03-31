package main

import (
	"github.com/mitchellh/cli"
	"github.com/ryanuber/pharos/command"
	"os"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{
		Writer: os.Stdout,
	}
	Commands = map[string]cli.CommandFactory{
		"start": func() (cli.Command, error) {
			return &command.StartCommand{
				Ui: ui,
			}, nil
		},
	}
}
