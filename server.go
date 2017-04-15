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
	// CAM group
	cam := r.Group("/cam")

	cam.GET("/command", handler.ParseCommand(), func(c *gin.Context) {
		game.Throw(c.MustGet("command").(*model.CamCommand))

		c.JSON(200, game.GetGame())
	})
	// CAM group end

	// GAME group
	g := r.Group("/game", handler.NoCache())

	g.GET("player", func(c *gin.Context) {
		game.SetPlayer(model.NewPlayer(c.Query("name")))
		game.Restart()

		c.JSON(200, game.GetGame())
	})

	r.LoadHTMLFiles("scoreboards/501.html", "scoreboards/start-game.html", "admin/editthrow.html")
	g.GET("/scoreboard", func(c *gin.Context) {
		if game.GetGame().Status == model.StatusCreate && len(game.GetGame().Players) < 1 {
			c.Redirect(301, "/game/start")
			return
		} else if game.GetGame().Status == model.StatusCreate {
			game.GetGame().Status = model.StatusStarted
		}

		c.HTML(200, "501.html", gin.H{})
	})
	g.GET("/start", func(c *gin.Context) {
		if game.GetGame().Status == model.StatusStarted {
			c.Redirect(301, "/game/scoreboard")
			return
		}

		c.HTML(200, "start-game.html", gin.H{})
	})
	g.GET("/end", func(c *gin.Context) {
		game.GetGame().Reset()
		c.Redirect(301, "/game/start")
	})
	g.Static("/start", "")
	// GAME group end

	// ADMIN group
	adm := r.Group("/admin", handler.NoCache())
	adm.GET("/throws", func(c *gin.Context) {
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
	// ADMIN group end

	// WS websocket
	// https://github.com/olahol/melody
	wsRoute := websocket.Load(func(s *melody.Session) {
		websocket.Clients[s] = &websocket.ClientInfo{ID: uuid.New(), IP: s.Request.RemoteAddr}
		s.Write(game.WebsocketOnConnectMsg())
	}, nil)
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
