package handler

import (
	"DAJ/Server/internal/buisness"
	"DAJ/Server/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) PostEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer r.Body.Close()
	params, err := url.ParseQuery(string(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Event := models.Event{User_id: params.Get("user_id"), Date: params.Get("date")}
	if Event.User_id == "" || Event.Date == "" {
		log.Println("Error: Wrong data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := r.URL.String()
	url = url[strings.Index(url, "/")+1:]
	result := buisness.Result(url, Event)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		resultError := models.Error{Err: err}
		_ = json.NewEncoder(w).Encode(resultError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	Event := models.Event{User_id: r.URL.Query().Get("user_id"), Date: r.URL.Query().Get("date")}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if Event.User_id == "" || Event.Date == "" {
		log.Println("Error: Wrong data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := r.URL.String()
	url = url[strings.Index(url, "/")+1 : strings.Index(url, "?")]
	result := buisness.Result(url, Event)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		resultError := models.Error{Err: err}
		_ = json.NewEncoder(w).Encode(resultError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
