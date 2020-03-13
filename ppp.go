package main

import (
    "fmt"
)
func processServer(s string) {
   fmt.Println(s, " : started processing")
}

func main() {
serverList := [...]string{"server1","server2","server3"}

for _,server := range serverList {
   fmt.Println("sending ", server)
   processServer(server)
}
}
