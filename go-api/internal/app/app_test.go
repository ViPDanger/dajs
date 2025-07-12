package app_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ViPDanger/dajs/go-api/internal/app"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
	log "github.com/ViPDanger/dajs/go-api/pkg/logger/v3"
)

func TestRun_AllEndpoints(t *testing.T) {
	srv, cancel := app.Run(log.Initialization("", ""), app.APIConfig{
		Addres: "127.0.0.1",
		Port:   "8090",
	})
	defer cancel()

	time.Sleep(500 * time.Millisecond) // ждём старта сервера

	type testCase struct {
		method string
		path   string
		body   string
		expect int
	}
	charData, _ := json.Marshal(dto.CharacterDTO{ID: "01"})

	cases := []testCase{
		{"GET", "/", "", http.StatusOK},
		{"POST", "/character", string(charData), http.StatusOK},
		{"GET", "/character", `01`, http.StatusOK},
		{"GET", "/character/all", ``, http.StatusOK},
		{"PUT", "/character", string(charData), http.StatusOK},
		{"DELETE", "/character", `02`, http.StatusOK},

		{"POST", "/character", ``, http.StatusBadRequest},
		{"GET", "/character", ``, http.StatusInternalServerError},
		{"PUT", "/character", ``, http.StatusBadRequest},
		{"DELETE", "/character", ``, http.StatusInternalServerError},
	}
	tokens := dto.TokensDTO{}
	dataUser, _ := json.Marshal(dto.UserDTO{Username: "виталий"})
	w := httptest.NewRecorder()
	// register
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8090/register", bytes.NewBuffer(dataUser))
	req.Header.Set("Content-Type", "application/json")

	srv.Handler.ServeHTTP(w, req)
	body, _ := io.ReadAll(w.Body)
	if w.Code != http.StatusCreated {
		t.Errorf("%s %s: ожидался %d, получен %d. Ответ: %s", "POST", "/register", http.StatusCreated, w.Code, string(body))
	}
	// login
	req, _ = http.NewRequest("POST", "http://127.0.0.1:8090/login", bytes.NewBuffer(dataUser))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	srv.Handler.ServeHTTP(w, req)
	body, _ = io.ReadAll(w.Body)
	if w.Code != http.StatusOK {
		t.Errorf("%s %s: ожидался %d, получен %d. Ответ: %s", "POST", "/login", http.StatusOK, w.Code, string(body))
	}
	_ = json.Unmarshal(body, &tokens)
	refreshToken, _ := json.Marshal(tokens.Refresh)
	// refresh
	req, _ = http.NewRequest("POST", "http://127.0.0.1:8090/refresh", bytes.NewBuffer(refreshToken))
	req.Header.Set("Content-Type", "application/json")
	srv.Handler.ServeHTTP(w, req)
	body, _ = io.ReadAll(w.Body)
	if w.Code != http.StatusOK {
		t.Errorf("%s %s: ожидался %d, получен %d. Ответ: %s", "POST", "/refresh", http.StatusOK, w.Code, string(body))
	}
	for _, tc := range cases {
		req, err := http.NewRequest(tc.method, "http://127.0.0.1:8090"+tc.path, bytes.NewBufferString(tc.body))
		if err != nil {
			t.Errorf("ошибка создания запроса: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", tokens.Access.Str)
		req.Header.Set("id", "01")
		w = httptest.NewRecorder()
		srv.Handler.ServeHTTP(w, req)
		body, _ := io.ReadAll(w.Body)
		if w.Code != tc.expect {
			t.Errorf("%s %s: ожидался %d, получен %d. Ответ: %s", tc.method, tc.path, tc.expect, w.Code, string(body))
		}
	}
}
