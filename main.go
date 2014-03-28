package main

import (
	"fmt"
	"github.com/ryanuber/pharos/monitor"
	"github.com/ryanuber/pharos/pharos"
)

func main() {
	fmt.Printf("pharos version %s\n", pharos.Version)
	process := monitor.NewMonitoredCommand(
		"true", 1)
	process.Monitor()
}
