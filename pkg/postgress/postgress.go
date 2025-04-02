package postgres

import "time"

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
}
