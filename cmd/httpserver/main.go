package main

import (
	"github.com/mathieuhays/go-specs-greet/adapters/httpserver"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	_ = http.ListenAndServe(":8080", handler)
}
