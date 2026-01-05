package http

import (
	"net/http"

	"github.com/jebesh67/golang-http/internal/http/handlers"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HandleHello)
}
