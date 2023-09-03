package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Greg")
	hostname, _ := os.Hostname()
	fmt.Println("Hostname: ", hostname)
	ReadTodos()
	os.Exit(0)
}

func ReadTodos() {
	fmt.Println("Read Todos")
	_, err := os.Open("/Users/greg/Documents/my-notes/hello/todo.md")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
}
