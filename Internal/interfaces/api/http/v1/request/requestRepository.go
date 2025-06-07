package request

import (
	"DAJ/Internal/domain/entity"
	"DAJ/pkg/logger"
	"net/http"
)

type HttpRepository struct {
	http.Client
	Log          logger.Ilogger
	Host         string
	accessToken  entity.AccessToken
	refreshToken entity.RefreshToken
}

func NewHttpRepository(logger logger.Ilogger, host string) *HttpRepository {
	return &HttpRepository{
		Log:  logger,
		Host: host,
	}
}
