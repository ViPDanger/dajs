package v1_01

import (
	"DAJ/Internal_Client/models"
	"DAJ/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (r *HttpRepository) Login(email, password string) error {
	data := models.LoginRequest{
		Email:    email,
		Password: password,
	}
	body, _ := json.Marshal(data)

	resp, err := http.Post(r.baseURL+"/auth/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return r.log.Logln(logger.Error, errors.New("Ошибка авторизации"))
	}

	var tokens models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return err
	}

	r.accessToken = tokens.AccessToken
	r.refreshToken = tokens.RefreshToken
	_ = r.log.Logln(logger.Release, "Login успешно. Access и Refresh токены сохранены.")
	return nil
}

func (r *HttpRepository) GetProtectedResource() error {
	req, _ := http.NewRequest("GET", r.baseURL+"/protected", nil)
	req.Header.Set("Authorization", "Bearer "+r.accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Если access token истёк — пытаемся обновить
	if resp.StatusCode == http.StatusUnauthorized {

		_ = r.log.Logln(logger.Release, "Access token истёк. Обновляем...")
		if err := r.refreshAccessToken(); err != nil {
			return err
		}
		// Повторяем запрос
		return r.GetProtectedResource()
	}

	body, _ := io.ReadAll(resp.Body)
	_ = r.log.Logln(logger.Release, "Ответ от /protected:", string(body))
	return nil
}

func (r *HttpRepository) refreshAccessToken() error {
	reqBody := map[string]string{
		"refresh_token": r.refreshToken,
	}
	body, _ := json.Marshal(reqBody)

	resp, err := http.Post(r.baseURL+"/auth/refresh", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка обновления токена: %s", resp.Status)
	}

	var tokens models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
		return err
	}

	r.accessToken = tokens.AccessToken
	fmt.Println("Access token обновлён.")
	return nil
}
