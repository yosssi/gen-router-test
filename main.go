package main

import (
	"fmt"
	"net/http"
)

func topHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "topHandler")
}

func usersIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "usersIndexHandler")
}

func usersShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "usersShowHandler")
}

func usersCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "usersCreateHandler")
}

//go:generate gen-router -i routes.json -o router.go
func main() {
	http.Handle("/", newRouter())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
