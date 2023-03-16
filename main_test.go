package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTimezoneRoute(t *testing.T) {
	router := setupRouter()

	t.Run("Request defalt timezone", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assertCorrectMessage(t, w.Code, 200)
	})

	t.Run("Request invalid timezone", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/?zone=local", nil)
		router.ServeHTTP(w, req)

		assertCorrectMessage(t, w.Code, 400)
	})

	t.Run("Request America/Sao_Paulo timezone", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/?zone=America/Sao_Paulo", nil)
		router.ServeHTTP(w, req)

		assertCorrectMessage(t, w.Code, 200)
	})
}
