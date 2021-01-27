package sshlib

import "fmt"

// ExecuteCommand executs the command provided, and returns the string output and error if any
func (s *SSH) ExecuteCommand(comamnd string) (string, error) {
	_, err := fmt.Fprintf(s.input, "%s\n", comamnd)
	if err != nil {
		return "", err
	}

	err = s.session.Wait()
	if err != nil {
		return "", err
	}

	output := s.outputBuffer.String()
	s.outputBuffer.Reset()

	return output, nil
}
