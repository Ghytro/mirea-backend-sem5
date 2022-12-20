package model

import (
	"backendmirea/pr3/internal/database"
	"backendmirea/pr3/internal/entity"
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type AuthModel struct {
	db          database.DBI
	redisClient *redis.Client
}

func NewAuthModel(db database.DBI, redisClient *redis.Client) *AuthModel {
	return &AuthModel{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *AuthModel) StartSession(ctx context.Context, username, password string) (string, time.Time, error) {
	var authedUser entity.AuthedUser
	expires := time.Now().Add(time.Hour)
	if err := r.db.ModelContext(ctx, &authedUser).
		Where("username = ? AND password = crypt(?, password)", username, password).
		Select(); err != nil {
		return "", time.Time{}, err
	}

	token := strings.ReplaceAll(uuid.New().String(), "-", "")
	if err := r.redisClient.HSet(ctx, "token:"+token, map[string]interface{}{
		"id":       authedUser.Id,
		"username": authedUser.UserName,
		"password": authedUser.Password,
		"email":    authedUser.Email,
		"name":     authedUser.Name,
		"is_admin": authedUser.IsAdmin,
	}).Err(); err != nil {
		return "", time.Time{}, err
	}
	return token, expires, nil
}
