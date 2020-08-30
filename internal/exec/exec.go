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
func Run(command, address, user, privKeys, path string) ([]byte, error) {
	if command == "" {
		return nil, errNoCommand
	}
	if address == "" {
		return nil, errNoAddress
	}
	/*if user == "" {
		return nil, errNoUser
	}*/
	if path == "" {
		return nil, errNoPath
	}
	var (
		client *simplessh.Client
		err    error
	)

	if privKeys != "" {
		client, err = simplessh.ConnectWithKeyFile(address, user, fmt.Sprintf("/home/%s/.ssh/id_rsa", "default"))
	} else {
		client, err = simplessh.ConnectWithPassword(address, "testing", "Nnc8Audm6Ai")
	}
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, fmt.Errorf("unable to init ssh client")
	}
	defer client.Close()

	output, err := client.Exec(command)
	if err != nil {
		return nil, err
	}

	return output, nil
}
