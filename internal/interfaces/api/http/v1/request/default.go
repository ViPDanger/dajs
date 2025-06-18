package request

import (
	"DAJ/internal/interfaces/api/dto"
	"bytes"
	"encoding/json"
	"net/http"
)

// GET Glossary

type Fetcher[T any] interface {
	Get(id string) (object T, err error)
	New(object T) (id string, err error)
	All() (objects []T, err error)
	Set(object T) (id string, err error)
	Delete(id string) (object T, err error)
}

type DefaultFetcher[T any, Tdto any] struct {
	ToDTO      func(T) Tdto
	ToEntity   func(Tdto) T
	Client     *HttpRepository
	GetPath    string
	NewPath    string
	AllPath    string
	SetPath    string
	DeletePath string
}

func NewDefaultFetcher[T any, Tdto any](
	ToDTO func(T) Tdto,
	ToEntity func(Tdto) T,
	Client *HttpRepository,
	GetPath string,
	NewPath string,
	AllPath string,
	SetPath string,
	DeletePath string) Fetcher[T] {
	return &DefaultFetcher[T, Tdto]{
		ToDTO:      ToDTO,
		ToEntity:   ToEntity,
		Client:     Client,
		GetPath:    GetPath,
		NewPath:    NewPath,
		AllPath:    AllPath,
		SetPath:    SetPath,
		DeletePath: DeletePath,
	}
}

func (f *DefaultFetcher[T, Tdto]) Get(id string) (object T, err error) {
	req, _ := http.NewRequest("GET", f.Client.Host+f.GetPath, nil)
	req.Header.Set("id", id)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var DTO Tdto
	if err = json.NewDecoder(resp.Body).Decode(&DTO); err != nil {
		return
	}
	object = f.ToEntity(DTO)
	return
}

// POST Glossary
func (f *DefaultFetcher[T, Tdto]) New(object T) (id string, err error) {
	body, err := json.Marshal(f.ToDTO(object))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", f.Client.Host+f.NewPath, bytes.NewBuffer(body))
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return
	}
	return message.Message, nil
}

// GET ALL Glossary
func (f *DefaultFetcher[T, Tdto]) All() (objects []T, err error) {
	req, _ := http.NewRequest("GET", f.Client.Host+f.AllPath, nil)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	DTOs := make([]Tdto, 0)
	if err = json.NewDecoder(resp.Body).Decode(&DTOs); err != nil {
		return
	}
	objects = make([]T, len(DTOs))
	for i := range DTOs {
		objects[i] = f.ToEntity(DTOs[i])
	}
	return
}

// PUT Glossary
func (f *DefaultFetcher[T, Tdto]) Set(object T) (id string, err error) {
	body, err := json.Marshal(f.ToDTO(object))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("PUT", f.Client.Host+f.SetPath, bytes.NewBuffer(body))
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return
	}
	return message.Message, nil
}

// DELETE Glossary
func (f *DefaultFetcher[T, Tdto]) Delete(id string) (object T, err error) {
	req, _ := http.NewRequest("DELETE", f.Client.Host+f.DeletePath, nil)
	req.Header.Set("id", id)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var DTO Tdto
	if err = json.NewDecoder(resp.Body).Decode(&DTO); err != nil {
		return
	}
	return f.ToEntity(DTO), nil
}
