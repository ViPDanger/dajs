package v1_01

import "DAJ/pkg/logger"

type HttpRepository struct {
	log          logger.Ilogger
	baseURL      string
	accessToken  string
	refreshToken string
}

func NewHttpRepository(logger logger.Ilogger, baseURL string) *HttpRepository {
	return &HttpRepository{
		baseURL: baseURL,
	}
}
