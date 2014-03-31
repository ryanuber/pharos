package main

import (
	"fmt"
	"github.com/ryanuber/pharos/monitor"
	"github.com/ryanuber/pharos/pharos"
)

func main() {
	fmt.Printf("pharos version %s\n", pharos.Version)

	fmt.Println("starting to monitor 'true' (1s)")
	go monitor.NewMonitoredCommand("true", 1).Monitor()

	fmt.Println("starting to monitor 'false' (3s)")
	monitor.NewMonitoredCommand("false", 3).Monitor()
}
