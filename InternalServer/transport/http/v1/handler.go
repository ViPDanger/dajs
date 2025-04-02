package v1

import (
	"io"
	"log"
	"net/http"
	"net/url"
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
	params.Get("user_id")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer r.Body.Close()
	params, err := url.ParseQuery(string(body))
	params.Get("user_id")
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		if err := json.NewEncoder(w).Encode(nil); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_ = json.NewEncoder(w).Encode(resultError)
			return
		}
	*/
	w.WriteHeader(http.StatusOK)
}
