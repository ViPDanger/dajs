package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ViPDanger/dajs/internal/interfaces/api/dto"
	"github.com/ViPDanger/dajs/internal/interfaces/api/mapper"
	logger "github.com/ViPDanger/dajs/pkg/logger/v3"
)

const (
	registerPath = "/auth/register"
	loginPath    = "/auth/login"
	refreshPath  = "/auth/refresh"
)

func (r *Client) Register(user, password string) error {

	body, err := json.Marshal(dto.UserDTO{
		Username: user,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("Client.Register()/%w", err)
	}
	req, _ := http.NewRequest("POST", r.Host+registerPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.Do(req)
	if err != nil {
		return fmt.Errorf("Client.Register()/%w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Client.Register()/%s", resp.Status)
	}
	var data dto.Message
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("Client.Register()/%w", err)
	}
	logger.Logln(logger.Debug, data)
	return nil

}

func (r *Client) Login(user, password string) error {
	body, err := json.Marshal(dto.UserDTO{
		Username: user,
		Password: password,
	})

	if err != nil {
		return fmt.Errorf("Client.Login()/%w", err)

	}
	req, _ := http.NewRequest("POST", r.Host+loginPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.Do(req)
	if err != nil {
		return fmt.Errorf("Client.Login()/%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Client.Login()/%s", resp.Status)
	}
	var tokensDTO dto.TokensDTO
	if err := json.NewDecoder(resp.Body).Decode(&tokensDTO); err != nil {
		return fmt.Errorf("Client.Login()/%w", err)
	}
	r.accessToken = mapper.ToAccessTokenEntity(tokensDTO.Access)
	r.refreshToken = mapper.ToRefreshTokenEntity(tokensDTO.Refresh)
	logger.Logln(logger.Debug, "Successeful Authorization. Access и Refresh tokens saved.")
	return nil
}

func (r *Client) RefreshAccessToken() error {

	body, err := json.Marshal(mapper.ToRefreshTokenDTO(r.refreshToken))
	if err != nil {
		return fmt.Errorf("Client.RefreshAccessToken()/%w", err)
	}
	resp, err := http.Post(r.Host+refreshPath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("Client.RefreshAccessToken()/%w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Client.Login()/%s", resp.Status)
	}
	var accessTokenDTO dto.AccessTokenDTO
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenDTO); err != nil {
		return fmt.Errorf("Client.RefreshAccessToken()/%w", err)
	}

	r.accessToken = mapper.ToAccessTokenEntity(accessTokenDTO)
	logger.Logln(logger.Debug, "Acces Token is refreshed")
	return nil
}

// реквест на middleware проверку protected источников
func (r *Client) doProtected(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+r.accessToken.Str)
	resp, err := r.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.doProtected()/%w", err)
	}
	// Если access token истёк — пытаемся обновить
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Logln(logger.Debug, "Access token is expired. Refreshing...")
		if err := r.RefreshAccessToken(); err != nil {
			return nil, fmt.Errorf("Client.doProtected()/%w", err)
		}
		// Повторяем запрос
		return r.doProtected(req)
	}
	return resp, nil
}
