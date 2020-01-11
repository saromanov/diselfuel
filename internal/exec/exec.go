// Package exec provides implementation of execution
// of commands via ssh
package exec

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Execute struct {
}

func (e *Execute) Do() error {
	_, err := executeCommand()
	return err
}
func executeCommand(command, hostname string, port string, config *ssh.ClientConfig) (string, error) {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	if err != nil {
		return "", fmt.Errorf("unable to connect: %v", err)
	}
	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("unable to make session: %v", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("unable to run command: %v", err)
	}

	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String()), nil
}
