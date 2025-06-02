package request

import (
	"DAJ/pkg/logger"
	"net/http"
)

type HttpRepository struct {
	http.Client
	log           logger.Ilogger
	Host          string
	accessString  string
	refreshString string
}

func NewHttpRepository(logger logger.Ilogger, host string) *HttpRepository {
	return &HttpRepository{
		log:  logger,
		Host: host,
	}
}
