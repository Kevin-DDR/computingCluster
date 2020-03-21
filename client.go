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
	msg := Message{IdType: 1, Id: 0, J: Job{
		Args:     nil,
		callback: nil,
	}, Res: JobResult{
		ExecErr:      nil,
		Stdout:       nil,
		Stderr:       nil,
		ExecDuration: 0,
	},
	}

	//Envoi de la demande de connexion
	fmt.Println("Envoi de la demande de conenxion")

	messageJSON, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
	conn.Write(messageJSON)
	conn.Write([]byte("\n"))

	reader := bufio.NewReader(os.Stdin)

	for {
		msg.J.Args = nil
		msg.Res.Stdout = nil
		msg.Res.Stderr = nil
		// read in input from stdin

		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		// send to socket
		msg.J.Args = append(msg.J.Args, text)
		msg.IdType = 4

		messageJSON, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("error:", err)
		}
		conn.Write(messageJSON)
		conn.Write([]byte("\n"))

		// listen for reply
		retour, _ := bufio.NewReader(conn).ReadString('\n')
		_ = json.Unmarshal([]byte(retour), &msg)

		if msg.Res.Stderr != nil {
			fmt.Println("erreur de commande ---> " + string(msg.Res.Stderr))

		} else {
			fmt.Print("Message from server: " + string(msg.Res.Stdout))
			//fmt.Print("Time Execution: " + msg.Res.ExecDuration)

		}

	}
}
