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
func Run(command, address, user, path string) ([]byte, error) {
	if command == "" {
		return nil, errNoCommand
	}
	if address == "" {
		return nil, errNoAddress
	}
	if user == "" {
		return nil, errNoUser
	}
	if path == "" {
		return nil, errNoPath
	}
	client, err := simplessh.ConnectWithKeyFile(address, user, fmt.Sprintf("/home/%s/.ssh/id_rsa", "default"))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	output, err := client.Exec(command)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Result: %s\n", output)
	return output, nil
}
