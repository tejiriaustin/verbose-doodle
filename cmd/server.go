package server

import "net/http"

func (m *muxrouter) CREATE(uri string, f func(w http.ResponseWriter, r *http.Request)) {

}
func (m *muxrouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {

}

func (m *muxrouter) SERVE(uri string, f func(w http.ResponseWriter, r *http.Request)) {

}
