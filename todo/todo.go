package todo

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayTodos() {
	fmt.Println("****************")
	fmt.Println("***   TODO   ***")
	fmt.Println("****************")
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
		fileTextLines = append(fileTextLines, removeHyphen(fileScanner.Text()))
	}
	readFile.Close()

	return fileTextLines
}

func removeHyphen(todo string) string {
	return todo[2:]
}
