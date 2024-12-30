package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kxw07/structs-practice/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote := *note.New(title, content)
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

	err = os.WriteFile("userNote.json", valueString, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the note title: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Enter the note content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&value)
	if err != nil {
		return "", err
	}

	if value == "" {
		return "", errors.New("input cannot be empty")
	}

	return value, nil
}
