package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kxw07/structs-practice/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote.Show()

	err = save(userNote)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func save(userNote note.Note) error {
	valueString, err := json.Marshal(userNote)
	if err != nil {
		return err
	}

	fileName := strings.ReplaceAll(userNote.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	return os.WriteFile(fileName, valueString, 0644)
}

func getNoteData() (string, string, error) {
	title := getUserInput("Enter the note title: ")
	content := getUserInput("Enter the note content: ")

	return title, content, nil
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
