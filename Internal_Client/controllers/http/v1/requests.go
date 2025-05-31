package v1

import (
	"DAJ/pkg/logger"
	"io"
	"net/http"
)

func (r *HttpRepository) GetProtectedResource() error {
	req, _ := http.NewRequest("GET", r.Host+"/protected", nil)
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
		if err := r.RefreshAccessToken(); err != nil {
			return err
		}
		// Повторяем запрос
		return r.GetProtectedResource()
	}

	body, _ := io.ReadAll(resp.Body)
	_ = r.log.Logln(logger.Release, "Ответ от /protected:", string(body))
	return nil
}
