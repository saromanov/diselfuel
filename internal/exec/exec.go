// Package exec provides implementation of execution
// of commands via ssh
package exec

import (
	"fmt"

	"github.com/sfreiberg/simplessh"
)

// Run provides execution of command
func Run(address, user, path string) error {
	client, err := simplessh.ConnectWithKeyFile(address, user, path)
	if err != nil {
		return err
	}
	defer client.Close()

	output, err := client.Exec("uptime")
	if err != nil {
		return err
	}

	fmt.Printf("Uptime: %s\n", output)
	return nil
}
