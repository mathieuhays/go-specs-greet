package httpserver

import (
	"fmt"
	"github.com/mathieuhays/go-specs-greet/domain/interactions"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if r.URL.Path == "/greet" {
		_, _ = fmt.Fprintf(w, interactions.Greet(name))
	}

	if r.URL.Path == "/curse" {
		_, _ = fmt.Fprintf(w, interactions.Curse(name))
	}
}
