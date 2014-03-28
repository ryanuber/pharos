package monitor

import (
	"fmt"
	"os/exec"
	"time"
)

type MonitoredCommand struct {
	command  string
	interval int
}

func NewMonitoredCommand(command string, interval int) *MonitoredCommand {
	return &MonitoredCommand{
		command:  command,
		interval: interval,
	}
}

func (m *MonitoredCommand) Monitor() error {
	for {
		cmd := exec.Command(m.command)
		if err := cmd.Run(); err != nil {
			fmt.Println(fmt.Sprintf("failure: %s", m.command))
		} else {
			fmt.Println(fmt.Sprintf("success: %s", m.command))
		}
		time.Sleep(time.Duration(m.interval) * time.Second)
	}
}
