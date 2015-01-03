package main

import (
	"fmt"
	"net/http"
)

func topHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "topHander")
}

func main() {
	http.Handle("/", newRouter())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
