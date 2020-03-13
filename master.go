package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

//TODO Rajouter une file d'attente des jobs
//TODO Creer une structure contenant un job et un client
//TODO Rajouter une go routine qui insere le job dans la file d'attente en ecoutant les clients
//TODO Rajotuer une go routine qui extrait un job de la file

type Noeud struct {
	conn net.Conn
	etat int
}

var noeuds [4]Noeud
var li net.Listener
var client net.Conn
var file []Message

func handlerConnexion(id int) {
	//Todo tester si il s'agit d'une connexion serveur ou client
	//Si c'est un client, lancer un handler pour recevoir un job
	//Si c'est un noeud lancer un handler pour retirer un job de la file

	noeuds[id].conn, _ = li.Accept()
	noeuds[id].etat = 1
}

func handlerJob() {

	//Reception d'un job du client
	message, _ := bufio.NewReader(noeuds[0].conn).ReadString('\n')
	var msg Message
	err := json.Unmarshal([]byte(message), &msg)
	file = append(file, msg)
	fmt.Print("msg : " + message + "\n")
	//Envoyer le job au noeud
	for i := 0; i < 4; i++ {
		//Si le noeud est disponible
		if noeuds[i].etat == 1 {
			//On envoi le job au noeud
			noeuds[i].conn.Write([]byte(message + "\n"))
			break

		}
	}

}

func main() {

	fmt.Println("Lancement du serveur")

	// listen on all interfaces
	li, _ = net.Listen("tcp", ":9001")
	for i := 0; i < 1; i++ {
		handlerConnexion(i)
	}
	handlerJob()
	for {

	}
}
