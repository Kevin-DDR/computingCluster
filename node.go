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

func run(jj Job) (res JobResult) {

	start := time.Now()
	var l []string = jj.Args[1:]

	var err error
	cmd := exec.Command(jj.Args[0], l...)
	res.Stdout, err = cmd.Output()
	end := time.Now()

	elapsed := end.Sub(start)

	res.ExecDuration = elapsed

	if err != nil {
		ee := string(err.Error())
		//fmt.Println(err.Error())
		res.Stderr = []byte(ee)
		return
	}
	//fmt.Println(elapsed)
	//fmt.Print(string(Stdout))
	return

}

func main() {

	argsWithoutProg := os.Args[1:]

	var j Job
	j.Args = argsWithoutProg

	tmp := run(j)
	fmt.Println(string(tmp.Stdout))
}
