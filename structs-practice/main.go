package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title   string
	Content string
	Created time.Time
}

func (note *Note) show() {
	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
	fmt.Println("Created:", note.Created)
}

func main() {
	var note Note
	var err error

	note.Title, err = getUserInput("Enter the note title: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	note.Content, err = getUserInput("Enter the note content: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	note.Created = time.Now()

	note.show()

	err = save(note)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func save(note Note) error {
	valueString, err := json.Marshal(note)
	if err != nil {
		return err
	}

	err = os.WriteFile("note.json", []byte(valueString), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getUserInput(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("input cannot be empty")
	}

	return value, nil
}
