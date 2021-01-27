package main

import (
	"log"

	"github.com/Nishan8583/ssh/sshlib"
)

func main() {
	sshConn, err := sshlib.New("username", "password", "127.0.0.1", "22")
	if err != nil {
		log.Fatal("Could not create SSH connection due to error", err)
	}
	defer sshConn.Close()

	output, err := sshConn.ExecuteCommand("config firewall address")
	log.Println("First command output", output, err)

	output, err = sshConn.ExecuteCommand("show")
	log.Println(output, err)
}
