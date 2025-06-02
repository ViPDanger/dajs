package request

import (
	"DAJ/Internal/interfaces/api/http/v1/request/models"
	"DAJ/pkg/logger"
	"encoding/json"
	"net/http"
)

func (r *HttpRepository) GetCharacter() error {
	req, _ := http.NewRequest("GET", r.Host+"/protected/character/get", nil)
	req.Header.Set("Authorization", "Bearer "+r.accessString)
	resp, err := r.doProtected(req)
	if err != nil {
		r.log.Logln(logger.Error, err)
		return err
	}
	defer resp.Body.Close()
	var data models.Message
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		_ = r.log.Logln(logger.Error, err)
		return err
	}
	_ = r.log.Logln(logger.Release, data.Message)
	return nil
}

func (r *HttpRepository) doProtected(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+r.accessString)
	resp, err := r.Do(req)
	if err != nil {
		return nil, err
	}
	// Если access token истёк — пытаемся обновить
	if resp.StatusCode == http.StatusUnauthorized {

		_ = r.log.Logln(logger.Debug, "Access token истёк. Обновляем...")
		if err := r.RefreshAccessToken(); err != nil {
			return nil, err
		}
		// Повторяем запрос
		return r.doProtected(req)
	}
	return resp, nil
}
