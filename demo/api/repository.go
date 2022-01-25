package main

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type Repository interface {
	Get(id int) (*Note, error)
	GetAll() []Note
	Create(note Note) Note
	Update(id int, note Note) error
	Delete(id int) error
}

type repository struct {
	notes  []Note
	mockID int
}

func (r *repository) Get(id int) (*Note, error) {
	for _, note := range r.notes {
		if note.ID == id {
			return &note, nil
		}
	}

	return nil, ErrNotFound
}

func (r *repository) GetAll() []Note {
	return r.notes
}

func (r *repository) Update(id int, note Note) error {
	for i, n := range r.notes {
		if n.ID == id {
			note.CreatedAt = n.CreatedAt
			note.ID = n.ID
			r.notes[i] = note
			return nil
		}
	}

	return ErrNotFound
}

func (r *repository) Create(note Note) Note {
	note.ID = r.mockID
	r.mockID++
	note.CreatedAt = time.Now()

	r.notes = append(r.notes, note)

	return note
}

func (r *repository) Delete(id int) error {
	for i, n := range r.notes {
		if n.ID == id {
			r.notes = append(r.notes[:i], r.notes[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func NewRepository() Repository {
	return &repository{
		notes:  []Note{},
		mockID: 1,
	}
}
