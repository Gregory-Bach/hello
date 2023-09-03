package main

import (
	"fmt"
	"os"

	"github.com/Gregory-Bach/hello/todo"
)

func main() {

	hostname := welcome()

	todo.DisplayTodos(hostname)

	goodbye()

	os.Exit(0)
}

func welcome() string {
	fmt.Println("Hello Greg")
	hostname, _ := os.Hostname()
	fmt.Println("Hostname: ", hostname)
	fmt.Println("")

	return hostname
}

func goodbye() {
	fmt.Println("")
	fmt.Println("*** Have a good day! ***")
	fmt.Println("")
}
