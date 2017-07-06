package websocket

import (
	"sync"
	"net/http"
	"github.com/google/uuid"
	"github.com/olahol/melody"
)

var (
	wsRoute *melody.Melody
	Clients map[*melody.Session]*ClientInfo
	lock    *sync.Mutex
)

type ClientInfo struct {
	ID uuid.UUID
	IP string
}

func Load(connectHandler func(*melody.Session), msgHandler func(*melody.Session, []byte)) *melody.Melody {
	wsRoute = melody.New()
	wsRoute.Upgrader.CheckOrigin = func(*http.Request) bool { return true }
	Clients = make(map[*melody.Session]*ClientInfo)
	lock = new(sync.Mutex)

	wsRoute.HandleConnect(func(s *melody.Session) {
		lock.Lock()
		if connectHandler != nil {
			connectHandler(s)
		}
		lock.Unlock()
	})

	wsRoute.HandleDisconnect(func(s *melody.Session) {
		lock.Lock()
		delete(Clients, s)
		lock.Unlock()
	})

	wsRoute.HandleMessage(func(s *melody.Session, msg []byte) {
		lock.Lock()
		if msgHandler != nil {
			msgHandler(s, msg)
		}
		lock.Unlock()
	})

	return wsRoute
}

func BroadcastMsg(msg []byte) {
	wsRoute.Broadcast(msg)
}
