package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	var user User

	wc, err := w.Write([]byte("Hello world!\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&user)
	if err != nil {
		slog.Error("error wirting response", "err", err)
		return
	}

	fmt.Println(user)

	fmt.Println("bytes written", wc)
}
