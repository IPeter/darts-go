package main

import (
	"darts-go/game"
	"darts-go/handler"
	"darts-go/model"
	"darts-go/websocket"

	"strconv"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./scoreboards/assets")
	r.Static("/static", "./scoreboards/static")
	// CAM group
	cam := r.Group("/cam")

	cam.GET("/command", handler.ParseCommand(), func(c *gin.Context) {
		command := c.MustGet("command").(*model.CamCommand)

		if command.Modifier == model.HandsOnBoard {
			game.SkipRound()
		} else if command.Modifier == model.GameEndsWithWinner {
			game.FinishGame()
		} else {
			game.Throw(command)
		}

		c.JSON(200, game.GetGame())
	})
	// CAM group end

	// GAME group
	g := r.Group("/game", handler.NoCache())

	g.GET("player", func(c *gin.Context) {
		game.SetPlayer(model.NewPlayer(c.Query("name")))
		game.SendGameDataToClients(game.WebsocketRestartGame)

		c.JSON(200, game.GetGame())
	})

	r.LoadHTMLFiles("scoreboards/awaiting.html",
							    "scoreboards/501.html",
							    "scoreboards/cricket.html",
									"admin/start-game.html",
									"admin/editthrow.html")

	g.GET("/scoreboard", func(c *gin.Context) {
		if game.GetGame().Status == model.StatusCreate {
			c.HTML(200, "awaiting.html", gin.H{})
		} else {
			c.HTML(200, game.GetGame().Name + ".html", gin.H{})
		}
	})
	g.Static("/start", "")
	// GAME group end

	// ADMIN group
	adm := r.Group("/admin", handler.NoCache())
	adm.GET("/throws", func(c *gin.Context) {
		if game.GetGame().Status == model.StatusCreate {
			c.Redirect(301, "/admin/setPlayer")
			return
		}

		c.HTML(200, "editthrow.html", gin.H{})
	})
	adm.GET("/setThrow", func(c *gin.Context) {
		playerID := c.Query("playerId")
		score, _ := strconv.Atoi(c.Query("score"))
		modifier, _ := strconv.Atoi(c.Query("modifier"))
		throwID := c.Query("throwId")

		thr := game.EditThrow(playerID, score, modifier, throwID)

		c.JSON(200, thr)
	})
	adm.GET("/start", func(c *gin.Context) {
		game.GetGame().Status = model.StatusStarted
		game.GetGame().Name = c.Query("gameType")
		game.GetGame().SubType = c.Query("gameSubType")
		game.SendGameDataToClients(game.WebsocketGameStarted)
		c.Redirect(301, "/admin/throws")
	})
	adm.GET("/saveAndCreateNew", func(c *gin.Context) {
		game.SaveAndCreate()

		c.Redirect(301, "/admin/setPlayer")
	})
	adm.GET("/setPlayer", func(c *gin.Context) {
		if game.GetGame().Status != model.StatusCreate {
			c.Redirect(301, "/admin/throws")
			return
		}
		c.HTML(200, "start-game.html", gin.H{})
	})
	// ADMIN group end

	// WS websocket
	// https://github.com/olahol/melody
	connectHandler := func(s *melody.Session) {
		websocket.Clients[s] = &websocket.ClientInfo{ID: uuid.New(), IP: s.Request.RemoteAddr}
		s.Write(game.WebsocketOnConnectMsg())
	}
	msgHandler := func(s *melody.Session, msg []byte) {
		if string(msg) == "__ping__" {
			s.Write([]byte("__pong__"))
		}
	}
	wsRoute := websocket.Load(connectHandler, msgHandler)
	r.GET("/ws", func(c *gin.Context) {
		wsRoute.HandleRequest(c.Writer, c.Request)
	})

	// WS websocket end

	r.Run() // listen and serve on 0.0.0.0:8080
}

/*

JSON

    Url   string `json:"url"`

	func getPages() []Page {
		raw, err := ioutil.ReadFile("./pages.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		var c []Page
		json.Unmarshal(raw, &c)
		return c
	}


*/
