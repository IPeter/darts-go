package model

import "github.com/google/uuid"

type Round struct {
	ID      string         `json:"id"`
	Throws  map[int]*Throw `json:"throws"`
	current int
	max     int
}

func NewRound() *Round {
	return &Round{
		ID:      uuid.New().String(),
		Throws:  make(map[int]*Throw, 3),
		current: 0,
		max:     3,
	}
}
