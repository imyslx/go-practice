package main

import (
	"bufio"
	"net"
	"strings"

	"log"
)

func CreateListener() net.Listener {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	return ln
}

func ListenWorker(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// output message received
		log.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
func main() {

	// Create Listen port and accept to connection.
	log.Print("Launching server...")
	ln := CreateListener()

	ListenWorker(ln)

}
