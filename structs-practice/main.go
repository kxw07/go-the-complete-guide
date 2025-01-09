package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kxw07/structs-practice/note"
	"github.com/kxw07/structs-practice/todo"
)

type outputtable interface {
	Save() error
	Display()
}

func output(o outputtable) error {
	o.Display()
	return o.Save()
}

func main() {
	printAny(aPlusB(1, 5))
	printAny(aPlusB(1.5, 5.5))
	printAny(aPlusB("Hello, ", "World!"))

	title, content := getNoteData()
	noteObj, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printAny(noteObj)

	todoContent := getTodoData()
	todoObj, err := todo.New(todoContent)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = output(noteObj)
	if err != nil {
		fmt.Println("Note save error: ", err)
		return
	}

	err = output(todoObj)
	if err != nil {
		fmt.Println("Todo save error: ", err)
		return
	}
}

func printAny(value any) {
	fmt.Println(value)
}

func aPlusB[T int | float64 | string](a, b T) T {
	return a + b
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the note title: ")
	content := getUserInput("Enter the note content: ")

	return title, content
}

func getTodoData() string {
	content := getUserInput("Enter the todo content: ")

	return content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
