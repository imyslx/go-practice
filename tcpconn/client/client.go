package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func MakeConnection(server, port string) net.Conn {
	conn, err := net.Dial("tcp", server+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func GetInput() string {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input text: ")
		text, _ := reader.ReadString('\n')

		if len(text) > 1 {
			return text
		}
	}
}

func PublishMessage(conn net.Conn) {

Loop:
	for {
		// Read in input from stdin
		text := GetInput()

		// If the input is "exit", we will stop the loop and end client.
		switch {
		case text == "exit\n":
			log.Print("Entering exit command. EXIT.")
			break Loop
		}
		log.Print("Will send text: " + text)

		// Send to socket
		conn.Write([]byte(text + "\n"))

		// Listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Print("Message from server: " + message)
	}
}

func main() {

	// connect to this socket
	conn := MakeConnection("127.0.0.1", "8081")
	defer conn.Close()

	log.Print("Connected.")

	PublishMessage(conn)
}
