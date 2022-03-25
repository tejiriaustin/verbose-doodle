package server

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tejiriaustin/verbose-doodle/internal/rest/handlers"
	"os"
	"time"
)

var (
	addr = flag.String("addr", ":", os.Getenv("PORT"))
	cert = flag.String("cert", "", ""
	key = flag.String("key", "", "")
	)

func Run() error {
	flag.Parse()

	if *addr == ":"{
		*addr = ":8080"
	}

	app := gin.New()
	app.GET("/", handlers.Welcome)
	app.GET("/room/create", handlers.CreateRoom)
	app.GET("/room/:uuid", handlers.Room)
	app.GET("/room/:uuid/websocket", )
	app.GET("/room/:uuid/chat", handlers.RoomChat)
	app.GET("/room/:uuid/chat/websockets", websocket.New(handlers.RoomChatWebsocket))
	app.GET("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewersocket))
	app.GET("/stream/:ssuid", handlers.Stream)
	app.GET("/stream/:ssuid/websocket", handlers.CreateRoom)
	app.GET("/stream/:ssuid/chat/websocket", handlers.CreateRoom)
	app.GET("/stream/:ssuid/viewer/websocket", handlers.CreateRoom)

}