package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Done      bool      `json:"done"`
	DoneAt    time.Time `json:"done_at,omitempty"`
}

func NewTodo(text string) *Todo {
	return &Todo{
		ID:        uuid.NewString(),
		Text:      text,
		CreatedAt: time.Now().UTC(),
	}
}
