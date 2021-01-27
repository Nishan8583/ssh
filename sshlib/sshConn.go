package sshlib

import (
	"bytes"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
)

// SSH holds the neccessary data for SSH connection, command execution input, and output storage
type SSH struct {
	username     string
	password     string
	host         string
	port         string
	client       *ssh.Client
	session      *ssh.Session
	output       []byte
	outputBuffer *bytes.Buffer
	input        io.WriteCloser
}

// New creates a new SSH type, instantiates it, establishes a connection, then a session, and returns it
func New(username, password, hostname, port string) (SSH, error) {

	sshConn := SSH{
		username: username,
		password: password,
		host:     hostname,
		port:     port,
	}

	// creating a SSH configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to host
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
	if err != nil {
		return sshConn, err
	}
	log.Printf("[+] Success Connection with %s established\n", hostname)

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		return sshConn, err
	}
	log.Printf("[+] Success Session with %s established\n", hostname)

	sshConn.session = sess

	// Setting up standard input for commands
	stdin, err := sshConn.session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	sshConn.input = stdin

	// storing the output
	newBuffer := bytes.NewBuffer(sshConn.output)
	sshConn.outputBuffer = newBuffer
	sshConn.session.Stdout = sshConn.outputBuffer
	sshConn.session.Stderr = sshConn.outputBuffer

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("SHELL gotten")
	return sshConn, nil

}

// Close closes the session, client connection and resets the buffer output respectively
func (s *SSH) Close() {
	s.session.Close()
	s.client.Close()
	s.outputBuffer.Reset()
}
