package service

import (
	"context"

	"github.com/google/go-github/v71/github"

	"github.com/abgeo/follytics/internal/domain/constant"
)

type AuthService interface {
	CurrentUser(ctx context.Context) *github.User
	Token(ctx context.Context) string
}

type Auth struct{}

var _ AuthService = (*Auth)(nil)

func NewAuth() *Auth {
	return &Auth{}
}

func (s *Auth) CurrentUser(ctx context.Context) *github.User {
	if user, ok := ctx.Value(constant.AuthUserKey).(*github.User); ok {
		return user
	}

	return nil
}

func (s *Auth) Token(ctx context.Context) string {
	if token, ok := ctx.Value(constant.AuthTokenKey).(string); ok {
		return token
	}

	return "<nil>"
}
