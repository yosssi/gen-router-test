package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", topHandler).Methods("GET")
	return r
}
