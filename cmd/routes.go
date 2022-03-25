package server

import "net/http"

type muxrouter struct {
}

type Routes interface {
	CREATE(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(uri string, f func(w http.ResponseWriter, r *http.Request))
}

func NewMuxRouter() Routes {
	return &muxrouter{}
}
