package main

import (
	"fmt"
	"net/http"
	"time"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	t := time.Now()
	fmt.Fprint(w, t)
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:8080", h)
}
