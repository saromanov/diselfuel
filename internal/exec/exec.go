// Package exec provides implementation of execution
// of commands via ssh
package exec

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)
func executeCommand(command, hostname string, port string, config *ssh.ClientConfig) string {
    conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
    session, _ := conn.NewSession()
    defer session.Close()

    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    session.Run(command)

    return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
}