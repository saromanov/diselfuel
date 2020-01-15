// Package exec provides implementation of execution
// of commands via ssh
package exec

import (
	"errors"
	"fmt"

	"github.com/sfreiberg/simplessh"
)

var (
	errNoCommand = errors.New("command is not defined")
	errNoAddress = errors.New("address is not defined")
	errNoUser    = errors.New("user is not defined")
	errNoPath    = errors.New("path is not defined")
)

// Run provides execution of command
func Run(command, address, user, path string) error {
	if command == "" {
		return errNoCommand
	}
	if address == "" {
		return errNoAddress
	}
	if user == "" {
		return errNoUser
	}
	if path == "" {
		return errNoPath
	}
	client, err := simplessh.ConnectWithKeyFile(address, user, path)
	if err != nil {
		return err
	}
	defer client.Close()

	output, err := client.Exec(command)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %s\n", output)
	return nil
}
