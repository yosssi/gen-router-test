package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", topHandle).Methods("GET")
	return r
}
