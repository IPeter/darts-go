package model

import "github.com/google/uuid"
import "time"

type Throw struct {
	ID       string    `json:"id"`
	Score    int       `json:"score"`
	Modifier int       `json:"modifier"`
	Time     time.Time `json:"time"`
}

func NewThrow(score int, modifier int) *Throw {
	return &Throw{
		ID:       uuid.New().String(),
		Score:    score,
		Modifier: modifier,
		Time:     time.Now().UTC(),
	}
}
