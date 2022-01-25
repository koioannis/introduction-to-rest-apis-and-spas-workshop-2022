package main

import (
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
}
