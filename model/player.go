package model

import (
	"github.com/google/uuid"
)

type Player struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Rounds       map[string]*Round `json:"rounds"`
	currentRound string
}

func NewPlayer(name string) *Player {
	return &Player{
		ID:           uuid.New().String(),
		Name:         name,
		currentRound: "",
		Rounds:       make(map[string]*Round, 20),
	}
}

func (p *Player) GetCurrentRoundID() string {
	return p.currentRound
}

func (p *Player) GetCurrentRound() *Round {
	if p.Rounds[p.currentRound] == nil {
		p.Rounds[p.currentRound] = NewRound()
	}

	return p.Rounds[p.currentRound]
}

func (p *Player) HasMoreThrow() bool {
	if p.GetCurrentRound().max == len(p.GetCurrentRound().Throws) {
		return false
	}

	return true
}

func (p *Player) IncRound() {
	p.currentRound = uuid.New().String()
}

func (p *Player) SetThrow(t *Throw) {
	p.Rounds[p.currentRound].Throws[t.ID] = t
}
