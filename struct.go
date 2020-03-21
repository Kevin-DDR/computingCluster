package main

import (
	"time"
)

type Job struct {
	Args     []string        `json: "Args"`
	callback func(JobResult) `json: "callback"`
}

type JobResult struct {
	ExecErr      error         `json: "ExecErr"`
	Stdout       []byte        `json: "Stdout"`
	Stderr       []byte        `json: "Stderr"`
	ExecDuration time.Duration `json: "ExecDuration"`
}

type Message struct {
	IdType int       `json: "idType"` //1 connexion Client Master, 2 = Noeuds master, 3 = deconnexion 4 = envoi d'un job
	Id     int       `json: "id"`
	J      Job       `json: "j"`
	Res    JobResult `json: "res"`
}
