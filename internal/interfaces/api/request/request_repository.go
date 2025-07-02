package request

import (
	"DAJ/internal/domain/entity"
	logger "DAJ/pkg/logger/v3"
	"net/http"
)

type Client struct {
	http.Client
	Host         string
	accessToken  entity.AccessToken
	refreshToken entity.RefreshToken
}

func NewHttpRepository(logger logger.Ilogger, host string) *Client {
	return &Client{
		Host: host,
	}
}
