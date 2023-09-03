package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Greg")
	hostname, _ := os.Hostname()
	fmt.Println("Hostname: ", hostname)
	fmt.Println("")
	DisplayTodos()
	os.Exit(0)
}

func DisplayTodos() {
	fmt.Println("*** TODO ***")
	todos := ReadTodos()
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func ReadTodos() []string {
	readFile, err := os.Open("/Users/greg/Documents/my-notes/hello/todo.md")
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	return fileTextLines
}
