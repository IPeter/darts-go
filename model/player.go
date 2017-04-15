package model

import (
	"github.com/google/uuid"
)

type Player struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Rounds       map[int]*Round `json:"rounds"`
	currentRound int
}

func NewPlayer(name string) *Player {
	return &Player{
		ID:           uuid.New().String(),
		Name:         name,
		currentRound: 0,
		Rounds:       make(map[int]*Round, 20),
	}
}

func (p *Player) GetCurrentRoundID() int {
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
	p.currentRound++
}

func (p *Player) SetThrow(t *Throw) {
	pos := len(p.GetCurrentRound().Throws)
	p.Rounds[p.currentRound].Throws[pos] = t
}

func (p *Player) GetThrowById(id string) (round *Round, thr *Throw) {
	for _, round = range p.Rounds {
		for _, thr = range round.Throws {
			if thr.ID == id {
				return
			}
		}
	}

	return nil, nil
}
