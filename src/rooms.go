package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Participants struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participants
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participants, 0)
}

func (r *RoomMap) Get(RoomID string) []Participants {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[RoomID]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	roomID := string(b)
	fmt.Println(roomID)
	
	r.Map[roomID] = []Participants{}

	return roomID
}

func (r *RoomMap) InsertIntoRoom(RoomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participants{host, conn}

	log.Println("Joining room: ", RoomID)
	r.Map[RoomID] = append(r.Map[RoomID], p)
}
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
