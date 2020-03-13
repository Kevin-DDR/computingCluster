package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

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

func run(jj Job) {

	start := time.Now()
	var l []string = jj.Args[1:]

	cmd := exec.Command(jj.Args[0], l...)
	stdout, err := cmd.Output()
	end := time.Now()
	elapsed := end.Sub(start)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(elapsed)
	fmt.Print(string(stdout))
	return

}

func main() {

	argsWithoutProg := os.Args[1:]

	var j Job
	j.Args = argsWithoutProg

	run(j)
}
