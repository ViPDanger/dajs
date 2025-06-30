package request

import (
	"DAJ/internal/interfaces/api/dto"
	"DAJ/internal/interfaces/api/mapper"
	logger "DAJ/pkg/logger/v3"
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

func (r *RequestRepository) Register(user, password string) error {

	body, err := json.Marshal(dto.UserDTO{
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
	var data dto.Message
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	r.Log.Logln(logger.Debug, data)
	return nil

}

func (r *RequestRepository) Login(user, password string) error {
	body, err := json.Marshal(dto.UserDTO{
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
	var tokensDTO dto.TokensDTO
	if err := json.NewDecoder(resp.Body).Decode(&tokensDTO); err != nil {
		return err
	}
	r.accessToken = mapper.ToAccessTokenEntity(tokensDTO.Access)
	r.refreshToken = mapper.ToRefreshTokenEntity(tokensDTO.Refresh)
	r.Log.Logln(logger.Debug, "Успешная Авторизация. Access и Refresh токены сохранены.")
	return nil

}

func (r *RequestRepository) RefreshAccessToken() error {

	body, err := json.Marshal(mapper.ToRefreshTokenDTO(r.refreshToken))
	if err != nil {
		return err
	}
	resp, err := http.Post(r.Host+refreshPath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	var accessTokenDTO dto.AccessTokenDTO
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenDTO); err != nil {
		return err
	}

	r.accessToken = mapper.ToAccessTokenEntity(accessTokenDTO)
	r.Log.Logln(logger.Debug, "Acces Token is refreshed")
	return nil
}

func (r *RequestRepository) doProtected(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+r.accessToken.Str)
	resp, err := r.Do(req)
	if err != nil {
		return nil, err
	}
	// Если access token истёк — пытаемся обновить
	if resp.StatusCode == http.StatusUnauthorized {

		r.Log.Logln(logger.Debug, "Access token истёк. Обновляем...")
		if err := r.RefreshAccessToken(); err != nil {
			return nil, err
		}
		// Повторяем запрос
		return r.doProtected(req)
	}
	return resp, nil
}
