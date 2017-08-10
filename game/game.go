package game

import (
	"darts-go/model"
	"darts-go/websocket"
	"encoding/json"
	"database/sql"
)

var (
	game     *model.Game
	gameView string
)

const (
	WebsocketStartGame    = "start"
	WebsocketGameStarted  = "started"
	WebsocketInsertThrow  = "insert_throw"
	WebsocketInsertDelete = "delete_throw"
	WebsocketEditThrow    = "edit_throw"
	WebsocketFinishGame   = "finish"
	WebsocketRestartGame  = "restart"
)

func GetGame() *model.Game {
	if game == nil {
		game = model.NewGame()
	}

	return game
}

func SendGameDataToClients(command string) {
	g, _ := json.Marshal(struct {
		Command string      `json:"command"`
		Game    *model.Game `json:"game"`
	}{
		Command: command,
		Game:    GetGame(),
	})

	websocket.BroadcastMsg(g)
}

func SaveToDb(db *sql.DB) {
	game := GetGame()
	sql_additem := `
	INSERT INTO games
	( uuid, type, subtype, created)
	values(?, ?, ?, CURRENT_TIMESTAMP)
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err2 := stmt.Exec(
		game.ID,
		game.Name,
		game.SubType)
	if err2 != nil { panic(err2) }

	//save the first round for every player, don't know any better place
	sql_additem2 := `
			INSERT INTO rounds
			( uuid, game_uid, player_uid, created)
			values(?, ?, ?, CURRENT_TIMESTAMP)
			`
	stmt2, err1 := db.Prepare(sql_additem2)
	if err1 != nil { panic(err1) }
	defer stmt2.Close()

	for _, item := range game.Players {
		_, err2 := stmt2.Exec(item.GetCurrentRound().ID, game.ID, item.ID)
		if err2 != nil { panic(err2) }
	}
}

func SaveAndCreate(db *sql.DB) {
	game = model.NewGame()

	SendGameDataToClients(WebsocketRestartGame)
}

func FinishGame() {
	SendGameDataToClients(WebsocketFinishGame)
}

func SetPlayer(player *model.Player, db *sql.DB) {
	sql_additem := `
	INSERT INTO players
	( uuid, name, created)
	values(?, ?, CURRENT_TIMESTAMP)
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err2 := stmt.Exec(
		player.ID,
		player.Name)
	if err2 != nil { panic(err2) }

	GetGame().SetPlayer(player)
}

func SkipRound() {
	GetGame().GetCurrentPlayer().IncRound()
	GetGame().NextPlayer()
}

func Throw(c *model.CamCommand, db *sql.DB) {
	player := GetGame().GetCurrentPlayer()
	if player.HasMoreThrow() == false {
		player.IncRound()
		sql_additem := `
			INSERT INTO rounds
			( uuid, game_uid, player_uid, created)
			values(?, ?, ?, CURRENT_TIMESTAMP)
			`
		stmt, err := db.Prepare(sql_additem)
		if err != nil { panic(err) }
		defer stmt.Close()

		_, err2 := stmt.Exec(
			player.GetCurrentRound().ID,
			GetGame().ID,
			player.ID)
		if err2 != nil { panic(err2) }

		GetGame().NextPlayer()

		player = GetGame().GetCurrentPlayer()
	}

	thr := model.NewThrow(c.Score, c.Modifier, c.X, c.Y, c.Cam1Img, c.Cam2Img)
	player.SetThrow(thr)

	sql_additem := `
	INSERT INTO throws
	( uuid, round_uid, score, modifier, x, y, cam1img, cam2img, cam1x, cam2x, edited_count, created)
	values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err2 := stmt.Exec(
		thr.ID,
		player.GetCurrentRound().ID,
		c.Score,
		c.Modifier,
		c.X,
		c.Y,
		c.Cam1Img,
		c.Cam2Img,
		0,
		0,
		0)
	if err2 != nil { panic(err2) }


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

func EditThrow(playerId string, score int, modifier int, throwId string, db *sql.DB) *model.Throw {
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

	sql_additem := `
	update throws set orig_score = score, orig_modifier = modifier, modified = CURRENT_TIMESTAMP,
	 edited_count = edited_count + 1, score = ?, modifier = ?
	 where uuid = ?
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err2 := stmt.Exec(score, modifier, thr.ID)
	if err2 != nil { panic(err2) }


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
