package main

import "time"

type Job struct {
	Args     []string
	callback func(JobResult)
}

type JobResult struct {
	ExecErr      error
	Stdout       []byte
	Stderr       []byte
	ExecDuration time.Duration
}

type Message struct {
	idType int
	id     int
	j      Job
	res    JobResult
}
