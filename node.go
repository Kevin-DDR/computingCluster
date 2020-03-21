package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:9001")
	msg := Message{IdType: 2, Id: 0, J: Job{
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

	for {

		// listen for reply
		retour, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + retour)
		var msg Message
		_ = json.Unmarshal([]byte(retour), &msg)

		fmt.Println("hhhhhhhhh", msg.J.Args)
		fmt.Println("hhhhhhhhh", len(msg.J.Args))

		words := strings.Fields(msg.J.Args[0])

		fmt.Println(words, len(words))

		words2 := words[1:]

		fmt.Println(words2)

		start := time.Now()

		cmd := exec.Command(words[0], words2...)

		msg.Res.Stdout, err = cmd.Output()
		fmt.Println(msg.Res.Stdout)
		fmt.Println(string(msg.Res.Stdout))
		end := time.Now()

		elapsed := end.Sub(start)

		msg.Res.ExecDuration = elapsed

		if err != nil {
			ee := string(err.Error())
			//fmt.Println(err.Error())
			msg.Res.Stderr = []byte(ee)
		}

		fmt.Println("Envoi de la réponse")
		msg.IdType = 5
		messageJSON, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("error:", err)
		}
		conn.Write(messageJSON)
		conn.Write([]byte("\n"))

	}

}
