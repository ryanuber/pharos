package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	cli := &cli.CLI{
		Args:     os.Args[1:],
		Commands: Commands,
	}
	retval, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s", err)
		return 1
	}
	return retval
}
