package main

import (
	"fmt"
	"os"
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

func run(a ...string) {
	/**
	//app := "ls"

	//args := [...]string{"-al", "-l"}

	start := time.Now()

	cmd := exec.Command(string(a[0]), string(a[1]))
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
	**/

}

func main() {

	/**args := [...]string{"ls", "-l"}

	run(args)**/

	argsWithoutProg := os.Args[1:]

	var j Job
	fmt.Println(j.Args)
	fmt.Println(argsWithoutProg)

	sum := 0
	for i := 0; i < len(argsWithoutProg); i++ {
		j.Args[i] :=argsWithoutProg[i]
		fmt.Println(sum)
	}

}
