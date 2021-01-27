package sshlib

import (
	"fmt"
	"time"
)

// ExecuteCommand executs the command provided, and returns the string output and error if any
func (s *SSH) ExecuteCommand(comamnd string, timeout time.Duration) (string, error) {
	_, err := fmt.Fprintf(s.input, "%s\n", comamnd)
	if err != nil {
		return "", err
	}

	output := s.outputBuffer.String()
	if len(output) == 0 {
		time.Sleep(timeout * time.Second)

		output = s.outputBuffer.String()
	}
	s.outputBuffer.Reset()

	return output, nil
}
