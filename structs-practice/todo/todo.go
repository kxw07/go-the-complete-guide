package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content cannot be empty")
	}

	return Todo{
		Content: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println("Content:", todo.Content)
}

func (todo Todo) Save() error {
	valueString, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", valueString, 0644)
}
