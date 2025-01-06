package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kxw07/structs-practice/note"
	"github.com/kxw07/structs-practice/todo"
)

type saver interface {
	Save() error
}

func save(s saver) error {
	return s.Save()
}

func main() {
	title, content := getNoteData()
	noteObj, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	noteObj.Show()

	todoContent := getTodoData()
	todoObj, err := todo.New(todoContent)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	todoObj.Show()

	err = save(noteObj)
	if err != nil {
		fmt.Println("Note save error: ", err)
		return
	}

	err = save(todoObj)
	if err != nil {
		fmt.Println("Todo save error: ", err)
		return
	}
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
