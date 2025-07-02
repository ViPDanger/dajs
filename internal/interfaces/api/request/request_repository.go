package request

import (
	"net/http"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	logger "github.com/ViPDanger/dajs/pkg/logger/v3"
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
