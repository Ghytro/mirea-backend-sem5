package auth

import (
	"context"
	"time"
)

type UseCase interface {
	StartSession(ctx context.Context, username string, password string) (string, time.Time, error)
}
