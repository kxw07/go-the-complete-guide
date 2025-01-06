package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Show() {
	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
	fmt.Println("CreatedAt:", note.CreatedAt)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
