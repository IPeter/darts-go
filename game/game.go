package game

import (
	"darts-go/model"
	"darts-go/websocket"
	"encoding/json"

	"github.com/google/uuid"
)

var (
	game     *model.Game
	gameView string
)

const (
	WebsocketStartGame    = "start"
	WebsocketInsertThrow  = "insert_throw"
	WebsocketInsertDelete = "delete_throw"
	WebsocketEditThrow    = "edit_throw"
	WebsocketRestartGame  = "restart"
)

func GetGame() *model.Game {
	if game == nil {
		game = model.NewGame()
	}

	return game
}

func SetPlayer(player *model.Player) {
	GetGame().SetPlayer(player)
}

func Throw(c *model.CamCommand) {
	player := GetGame().GetCurrentPlayer()
	if c.Modifier == -1 {
		player.IncRound()
		GetGame().NextPlayer()
		return
	}
	if player.HasMoreThrow() == false {
		player.IncRound()

		GetGame().NextPlayer()

		player = GetGame().GetCurrentPlayer()
	}

	thr := &model.Throw{
		ID:       uuid.New().String(),
		Score:    c.Score,
		Modifier: c.Modifier,
	}
	player.SetThrow(thr)

	jsonThr, _ := json.Marshal(struct {
		Command  string       `json:"command"`
		ID       string       `json:"id"`
		PlayerID string       `json:"playerId"`
		RoundID  string       `json:"roundId"`
		Thr      *model.Throw `json:"throw"`
	}{
		Command:  WebsocketInsertThrow,
		ID:       thr.ID,
		PlayerID: player.ID,
		RoundID:  player.GetCurrentRound().ID,
		Thr:      thr,
	})

	websocket.BroadcastMsg(jsonThr)
}

func Restart() {
	g, _ := json.Marshal(struct {
		Command string      `json:"command"`
		Game    *model.Game `json:"game"`
	}{
		Command: WebsocketRestartGame,
		Game:    GetGame(),
	})

	websocket.BroadcastMsg(g)
}

func WebsocketOnConnectMsg() []byte {
	g, _ := json.Marshal(struct {
		Command string      `json:"command"`
		Game    *model.Game `json:"game"`
	}{
		Command: WebsocketStartGame,
		Game:    GetGame(),
	})

	return g
}

func EditThrow(playerId string, score int, modifier int, throwId string) *model.Throw {
	player := GetGame().GetPlayerById(playerId)

	if player == nil {
		return &model.Throw{}
	}

	round, thr := player.GetThrowById(throwId)

	if round == nil || thr == nil {
		return &model.Throw{}
	}

	thr.Score = score
	thr.Modifier = modifier

	jsonThr, _ := json.Marshal(struct {
		Command  string       `json:"command"`
		ID       string       `json:"id"`
		PlayerID string       `json:"playerId"`
		RoundID  string       `json:"roundId"`
		Thr      *model.Throw `json:"throw"`
	}{
		Command:  WebsocketEditThrow,
		ID:       thr.ID,
		PlayerID: player.ID,
		RoundID:  round.ID,
		Thr:      thr,
	})

	websocket.BroadcastMsg(jsonThr)

	return thr
}
