package main

import (
	"fmt"
	"os"

	"github.com/Gregory-Bach/hello/todo"
)

func main() {

	args := os.Args[1:]

	hostname, _ := os.Hostname()

	if len(args) == 0 {
		welcome()
	}

	if len(args) > 0 && args[0] != "--todo" {
		welcome()
	}

	todo.DisplayTodos(hostname)

	goodbye()

	os.Exit(0)
}

func welcome() {
	fmt.Println("Hello Greg")
	fmt.Println("")
}

func goodbye() {
	fmt.Println("")
	fmt.Println("*** Have a good day! ***")
	fmt.Println("")
}
