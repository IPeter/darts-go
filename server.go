package main

import (
	"darts-go/game"
	"darts-go/handler"
	"darts-go/model"
	"darts-go/websocket" 

	"github.com/google/uuid"
	"github.com/olahol/melody"

	"gopkg.in/gin-gonic/gin.v1"
)

type GopherInfo struct {
	ID, X, Y string
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./scoreboards/assets")
	// CAM group
	cam := r.Group("/cam")

	cam.GET("/command", handler.ParseCommand(), func(c *gin.Context) {
		game.Throw(c.MustGet("command").(*model.CamCommand))

		c.JSON(200, game.GetGame())
	})
	// CAM group

	// GAME group
	g := r.Group("/game")

	g.GET("player", func(c *gin.Context) {
		game.SetPlayer(model.NewPlayer(c.Query("name")))

		c.JSON(200, game.GetGame())
	})

	r.LoadHTMLFiles("scoreboards/501.html", "scoreboards/start-game.html")
	g.GET("/scoreboard", func(c *gin.Context) {
		c.HTML(200, "501.html", gin.H{})
	})
	g.GET("/start", func(c *gin.Context) {
		c.HTML(200, "start-game.html", gin.H{})
	})
	g.Static("/start", "")
	// GAME group

	// WS websocket
	// https://github.com/olahol/melody
	wsRoute := websocket.Load(func(s *melody.Session) {
		websocket.Clients[s] = &websocket.ClientInfo{ID: uuid.New(), IP: s.Request.RemoteAddr}
		s.Write(game.WebsocketOnConnectMsg())
	}, nil)
	r.GET("/ws", func(c *gin.Context) {
		wsRoute.HandleRequest(c.Writer, c.Request)
	})

	// WS websocket

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
