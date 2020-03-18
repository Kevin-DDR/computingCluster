package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9001")
	var j = Job{}
	var jr JobResult
	var msg Message = Message{1, 0, j, jr}

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		msg.j.Args = append(msg.j.Args, text)

		message, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(message)
		conn.Write(message)
		conn.Write([]byte("\n"))

		fmt.Println(text)
		// listen for reply
		retour, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + retour)
	}
}
