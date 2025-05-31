package v1

import (
	"DAJ/Internal_Client/models"
	"DAJ/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	registerPath = "/auth/register/"
	loginPath    = "/auth/login/"
	refreshPath  = "/auth/refresh/"
)

func (r *HttpRepository) Register(user, password string) error {

	body, err := json.Marshal(models.User{
		Username: user,
		Password: password,
	})
	if err != nil {
		return r.log.Logln(logger.Error, err)
	}
	path := r.Host + registerPath
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return r.log.Logln(logger.Error, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return r.log.Logln(logger.Error, resp.Status)
	}
	var data models.RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return r.log.Logln(logger.Error, data)
	}
	return r.log.Logln(logger.Debug, data)

}

func (r *HttpRepository) Login(user, password string) error {
	body, err := json.Marshal(models.User{
		Username: user,
		Password: password,
	})

	if err != nil {
		return r.log.Logln(logger.Error, err)

	}
	path := r.Host + loginPath
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(body))
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

	r.accessToken = tokens.AccessToken
	r.refreshToken = tokens.RefreshToken
	_ = r.log.Logln(logger.Debug, "Успешная Авторизация. Access и Refresh токены сохранены.")
	return nil

}

func (r *HttpRepository) RefreshAccessToken() error {

	body, err := json.Marshal(map[string]string{
		"refresh_token": r.refreshToken,
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

	r.accessToken = tokens.AccessToken
	_ = r.log.Logln(logger.Debug, "Acces Token is refreshed")
	return nil
}
