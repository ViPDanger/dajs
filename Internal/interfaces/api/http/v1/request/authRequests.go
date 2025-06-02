package request

import (
	"DAJ/Internal/interfaces/api/http/v1/request/models"
	"DAJ/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	registerPath = "/auth/register"
	loginPath    = "/auth/login"
	refreshPath  = "/auth/refresh"
)

func (r *HttpRepository) Register(user, password string) error {

	body, err := json.Marshal(models.User{
		Username: user,
		Password: password,
	})
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", r.Host+registerPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}
	var data models.RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	return r.log.Logln(logger.Debug, data)

}

func (r *HttpRepository) Login(user, password string) error {
	body, err := json.Marshal(models.User{
		Username: user,
		Password: password,
	})

	if err != nil {
		return err

	}
	req, _ := http.NewRequest("POST", r.Host+loginPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	var tokens models.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return err
	}

	r.accessString = tokens.AccessString
	r.refreshString = tokens.RefreshString
	_ = r.log.Logln(logger.Debug, "Успешная Авторизация. Access и Refresh токены сохранены.")
	return nil

}

func (r *HttpRepository) RefreshAccessToken() error {

	body, err := json.Marshal(models.RefreshToken{
		RefreshString: r.refreshString,
	})
	if err != nil {
		return r.log.Logln(logger.Error, err)
	}
	resp, err := http.Post(r.Host+refreshPath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	var tokens models.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return err
	}

	r.accessString = tokens.AccessString
	_ = r.log.Logln(logger.Debug, "Acces Token is refreshed")
	return nil
}
