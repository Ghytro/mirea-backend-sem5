package auth

import (
	"context"
	"time"
)

type Repository interface {
	StartSession(ctx context.Context, username string, password string) (string, time.Time, error)
}
