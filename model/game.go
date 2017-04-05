package model

import (
	"dartsvader-go/helper"

	"github.com/google/uuid"
)

const (
	StatusCreate  int = 0
	StatusStarted int = 1
)

type Game struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Players       map[string]*Player `json:"players"`
	Status        int                `json:"status"`
	currentPlayer string
	maxPlayer     int
}

func NewGame() *Game {
	return &Game{
		ID:            uuid.New(),
		Status:        StatusCreate,
		currentPlayer: "",
		maxPlayer:     2,
		Name:          "501",
		Players:       make(map[string]*Player, 2),
	}
}

func (g *Game) SetPlayer(player *Player) {
	if len(g.Players) < 1 {
		g.currentPlayer = player.ID
	}

	g.Players[player.ID] = player
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.Players[g.currentPlayer]
}

func (g *Game) GetCurrentPlayerId() string {
	return g.currentPlayer
}

func (g *Game) NextPlayer() {
	_, _, next := helper.GetMapPosition(g.Players, g.currentPlayer)
	g.currentPlayer = next
}
