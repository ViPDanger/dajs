package request

import (
	"DAJ/internal/domain/entity"
	logger "DAJ/pkg/logger/v3"
	"net/http"
)

type RequestRepository struct {
	http.Client
	Log          logger.Ilogger
	Host         string
	accessToken  entity.AccessToken
	refreshToken entity.RefreshToken
}

func NewHttpRepository(logger logger.Ilogger, host string) *RequestRepository {
	return &RequestRepository{
		Log:  logger,
		Host: host,
	}
}
