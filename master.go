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
	Id   int
}

func remove(slice []Noeud, s int) []Noeud {
	return append(slice[:s], slice[s+1:]...)
}

func removeMsg(slice []Message, s int) []Message {
	return append(slice[:s], slice[s+1:]...)
}

var noeuds []Noeud
var listen net.Listener
var clients []Noeud
var file []Message
var compteur int = 0

var sem semaphore = make(semaphore, 1)

func handlerConnexion(port string) {
	//Todo tester si il s'agit d'une connexion serveur ou client
	//Si c'est un client, lancer un handler pour recevoir un job
	//Si c'est un noeud lancer un handler pour retirer un job de la file

	li, _ := net.Listen("tcp", port)
	for {
		tmp, _ := li.Accept()
		message, _ := bufio.NewReader(tmp).ReadString('\n')
		fmt.Println("Connexion recue : ", message)
		var msg Message
		_ = json.Unmarshal([]byte(message), &msg)
		fmt.Println("Message recu ! IdType : ", msg.IdType)
		switch msg.IdType {
		case 1:
			clients = append(clients, Noeud{tmp, 1, compteur})
			go handlerJob(compteur)
			compteur++
		case 2:
			noeuds = append(noeuds, Noeud{tmp, 1, compteur})
			go handlerNoeud(compteur)
			compteur++
		}

	}

}

func handlerJob(Id int) {
	//TODO tester si c'est une demande de deconnexion
	var index int
	for i := 0; i < len(clients); i++ {
		if clients[i].Id == Id {
			index = i
			break
		}
	}
	for {
		message, _ := bufio.NewReader(clients[index].conn).ReadString('\n')
		var msg Message
		_ = json.Unmarshal([]byte(message), &msg)
		switch msg.IdType {
		case 3:

			clients = remove(clients, index)
			return
		default:
			msg.Id = Id
			file = append(file, msg)
			fmt.Print("msg : " + message + "\n")
		}
	}
	//Reception d'un job du client

	//Envoyer le job au noeud

}

func handlerNoeud(Id int) {

	for {
		if len(file) > 0 && len(noeuds) > 0 {
			for i := 0; i < len(noeuds); i++ {
				//Si le noeud est disponible
				Lock(sem)
				if len(file) > 0 && noeuds[i].etat == 1 {
					//On envoi le job au noeud
					message, _ := json.Marshal(file[0])

					file = removeMsg(file, 0)

					noeuds[i].conn.Write(message)
					noeuds[i].conn.Write([]byte("\n"))
					retour, _ := bufio.NewReader(noeuds[i].conn).ReadString('\n')
					var msg Message
					_ = json.Unmarshal([]byte(message), &msg)

					var index int
					for i := 0; i < len(clients); i++ {
						if clients[i].Id == msg.Id {
							index = i
							break
						}
					}
					switch msg.IdType {
					case 3:
						//Deconnexion
						noeuds = remove(noeuds, i)
						return
					default:
						clients[index].conn.Write([]byte(retour))
						clients[index].conn.Write([]byte("\n"))
					}
					Unlock(sem)

					//TODO Attendre une reponse et lar envoyer
					break

				}
				Unlock(sem)

			}
		}

	}

}

func main() {

	fmt.Println("Lancement du serveur")

	// listen on all interfaces

	go handlerConnexion(":9001")
	go handlerConnexion(":9002")
	go handlerConnexion(":9003")
	go handlerConnexion(":9004")
	for {

	}
}
