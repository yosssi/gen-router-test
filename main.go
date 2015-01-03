package main

import (
	"fmt"
	"net/http"
)

func topHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "topHandler")
}

//go:generate gen-router -i routes.json -o router.go
func main() {
	http.Handle("/", newRouter())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
