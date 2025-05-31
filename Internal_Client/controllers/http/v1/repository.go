package v1

import (
	"DAJ/pkg/logger"
	"net/http"
)

type HttpRepository struct {
	http.Client
	log          logger.Ilogger
	Host         string
	accessToken  string
	refreshToken string
}

func NewHttpRepository(logger logger.Ilogger, host string) *HttpRepository {
	return &HttpRepository{
		log:  logger,
		Host: host,
	}
}
