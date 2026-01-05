package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jebesh67/golang-http/internal/http/handlers"
)

func TestHandleHello(t *testing.T) {
	testPayload := handlers.User{
		Name: "Jebesh",
		Age:  21,
	}

	body, err := json.Marshal(testPayload)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

	handlers.HandleHello(w, r)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad response code, expected %v but got %v\nbody: %s\n",
			desiredCode, w.Code, w.Body.String())
	}
}
