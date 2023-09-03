package main

import (
	"fmt"
	"os"

	"github.com/Gregory-Bach/hello/todo"
)

func main() {

	welcome()

	todo.DisplayTodos()

	goodbye()

	os.Exit(0)
}

func welcome() {
	fmt.Println("Hello Greg")
	hostname, _ := os.Hostname()
	fmt.Println("Hostname: ", hostname)
	fmt.Println("")
}

func goodbye() {
	fmt.Println("")
	fmt.Println("*** Have a good day! ***")
	fmt.Println("")
}
