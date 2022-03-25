package model

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Sources struct {
	Conn *websocket.Conn
}

type Participants struct {
	Host bool
	Conn *websocket.Conn
}

type Event struct {
	key   string
	Mutex sync.RWMutex
	P     []Participants
	S     []Sources
}

func MakeNewEvent(key string, P_No int, S_No int) Event {
	return Event{
		key: key,
		P:   make([]Participants, P_No),
		S:   make([]Sources, S_No),
	}
}
