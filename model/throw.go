package model

import "github.com/google/uuid"

type Throw struct {
	ID       string `json:"id"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}

func NewThrow(score int, modifier int) *Throw {
	return &Throw{
		ID:       uuid.New().String(),
		Score:    score,
		Modifier: modifier,
	}
}
