package monitor

import (
	"fmt"
	"os/exec"
	"time"
)

// MonitoredCommand represents a command that runs periodically to test
// various conditions on a system.
type MonitoredCommand struct {
	command  string
	interval int
}

// Simple method to create a new monitored command object
func NewMonitoredCommand(command string, interval int) *MonitoredCommand {
	return &MonitoredCommand{
		command:  command,
		interval: interval,
	}
}

// Monitor begins periodically executing a command and reporting success
// or failure. It will probably run in a separate goroutine.
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
