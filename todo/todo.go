package todo

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayTodos(hostname string) {
	fmt.Println("****************")
	fmt.Println("***   TODO   ***")
	fmt.Println("****************")

	fileUrl := ""

	if hostname == "deathstar" {
		fileUrl = "/mnt/c/Users/grego/Documents/my-notes/hello/todo.md"
	} else if hostname == "Gregorys-Air" {
		fileUrl = "/Users/greg/Documents/my-notes/hello/todo.md"
	} else {
		fmt.Println("Unknown or no hostname: ", hostname)
		fmt.Println("Please config correctly...")
		os.Exit(2)
	}

	todos := ReadTodos(fileUrl)
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func ReadTodos(url string) []string {

	readFile, err := os.Open(url)

	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan() {
		if len(fileScanner.Text()) > 1 {
			fileTextLines = append(fileTextLines, removeHyphen(fileScanner.Text()))
		}
	}
	readFile.Close()

	return fileTextLines
}

func removeHyphen(todo string) string {
	if todo[0:2] != "- " {
		return fmt.Sprint("[ ] ", todo)
	}

	return todo[2:]
}
