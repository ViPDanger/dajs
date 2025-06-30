package request

import (
	"DAJ/internal/interfaces/api/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type Fetcher[T any] interface {
	Get(id string) (object T, err error)
	New(object T) (id string, err error)
	All() (objects []T, err error)
	Set(object T) (id string, err error)
	Delete(id string) (object T, err error)
}

type defaultFetcher[T any, Tdto any] struct {
	ToDTO      func(T) Tdto
	ToEntity   func(Tdto) T
	Client     *Client
	GetPath    string
	NewPath    string
	AllPath    string
	SetPath    string
	DeletePath string
}

func NewDefaultFetcher[T any, Tdto any](
	ToDTO func(T) Tdto,
	ToEntity func(Tdto) T,
	Client *Client,
	GetPath string,
	NewPath string,
	AllPath string,
	SetPath string,
	DeletePath string) Fetcher[T] {
	return &defaultFetcher[T, Tdto]{
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

// GET
func (f *defaultFetcher[T, Tdto]) Get(id string) (object T, err error) {
	req, _ := http.NewRequest("GET", f.Client.Host+f.GetPath, nil)
	req.Header.Set("id", id)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return object, fmt.Errorf("DefaultFetcher[%s].Get()/%w", TName, err)
	}
	defer resp.Body.Close()
	var DTO Tdto
	if err = json.NewDecoder(resp.Body).Decode(&DTO); err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return object, fmt.Errorf("DefaultFetcher[%s].Get()/%w", TName, err)
	}
	object = f.ToEntity(DTO)
	return
}

// POST
func (f *defaultFetcher[T, Tdto]) New(object T) (id string, err error) {
	TName := reflect.TypeOf(object).Name()
	body, err := json.Marshal(f.ToDTO(object))
	if err != nil {
		return "", fmt.Errorf("DefaultFetcher[%s].New()/%w", TName, err)
	}
	req, _ := http.NewRequest("POST", f.Client.Host+f.NewPath, bytes.NewBuffer(body))
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return "", fmt.Errorf("DefaultFetcher[%s].New()/%w", TName, err)

	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return "", fmt.Errorf("DefaultFetcher[%s].New()/%w", TName, err)
	}
	return message.Message, nil
}

// GET all
func (f *defaultFetcher[T, Tdto]) All() (objects []T, err error) {

	TName := reflect.TypeOf(objects).Name()
	req, _ := http.NewRequest("GET", f.Client.Host+f.AllPath, nil)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultFetcher[%s].New()/%w", TName, err)
	}
	defer resp.Body.Close()
	DTOs := make([]Tdto, 0)
	if err = json.NewDecoder(resp.Body).Decode(&DTOs); err != nil {
		return nil, fmt.Errorf("DefaultFetcher[%s].New()/%w", TName, err)
	}
	objects = make([]T, len(DTOs))
	for i := range DTOs {
		objects[i] = f.ToEntity(DTOs[i])
	}
	return
}

// PUT
func (f *defaultFetcher[T, Tdto]) Set(object T) (id string, err error) {
	body, err := json.Marshal(f.ToDTO(object))
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return "", fmt.Errorf("DefaultFetcher[%s].Set()/%w", TName, err)
	}
	req, _ := http.NewRequest("PUT", f.Client.Host+f.SetPath, bytes.NewBuffer(body))
	resp, err := f.Client.doProtected(req)
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return "", fmt.Errorf("DefaultFetcher[%s].Set()/%w", TName, err)
	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return "", fmt.Errorf("DefaultFetcher[%s].Set()/%w", TName, err)
	}
	return message.Message, nil
}

// DELETE Glossary
func (f *defaultFetcher[T, Tdto]) Delete(id string) (object T, err error) {
	req, _ := http.NewRequest("DELETE", f.Client.Host+f.DeletePath, nil)
	req.Header.Set("id", id)
	resp, err := f.Client.doProtected(req)
	if err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return object, fmt.Errorf("DefaultFetcher[%s].Set()/%w", TName, err)
	}
	defer resp.Body.Close()
	var DTO Tdto
	if err = json.NewDecoder(resp.Body).Decode(&DTO); err != nil {
		TName := reflect.TypeOf(new(T)).Name()[1:]
		return object, fmt.Errorf("DefaultFetcher[%s].Set()/%w", TName, err)
	}
	return f.ToEntity(DTO), nil
}
