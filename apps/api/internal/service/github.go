package service

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v68/github"

	"github.com/abgeo/follytics/internal/config"
)

type GithubService interface {
	CreateJWT() (string, error)
	WithToken(token string) *Github
	WithInstallationToken(ctx context.Context) (*Github, error)

	GetAPIRateLimits(ctx context.Context) (*github.RateLimits, *github.Response, error)
	GetUser(ctx context.Context, username string) (*github.User, *github.Response, error)
	GetUserByID(ctx context.Context, id int64) (*github.User, *github.Response, error)
	GetUserFollowers(ctx context.Context, username string, page int, limit int) ([]*github.User, *github.Response, error)
	CollectUserFollowers(ctx context.Context, username string, limit int) ([]*github.User, error)
}

type Github struct {
	config *config.Config
	logger *slog.Logger

	client *github.Client
}

var _ GithubService = (*Github)(nil)

func NewGithub(config *config.Config, logger *slog.Logger) *Github {
	client := github.NewClient(nil)

	return &Github{
		config: config,
		logger: logger.With(
			slog.String("component", "service"),
			slog.String("service", "github"),
		),
		client: client,
	}
}

func (s *Github) CreateJWT() (string, error) {
	key, err := s.loadPrivateKey()
	if err != nil {
		return "", err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(s.config.GitHub.JWTExpiration) * time.Minute).Unix(),
		"iss": s.config.GitHub.AppClientID,
		"alg": "RS256",
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func (s *Github) WithToken(token string) *Github {
	return &Github{
		config: s.config,
		logger: s.logger,
		client: s.client.WithAuthToken(token),
	}
}

func (s *Github) WithInstallationToken(ctx context.Context) (*Github, error) {
	token, err := s.CreateJWT()
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT: %w", err)
	}

	client := s.client.WithAuthToken(token)

	data, _, err := client.Apps.CreateInstallationToken(ctx, s.config.GitHub.AppInstallationID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to craete installation token: %w", err)
	}

	return s.WithToken(*data.Token), nil
}

func (s *Github) GetAPIRateLimits(ctx context.Context) (*github.RateLimits, *github.Response, error) {
	meta, res, err := s.client.RateLimit.Get(ctx)
	if err != nil {
		return nil, res, fmt.Errorf("failed to get API Rate Limits from GitHub: %w", err)
	}

	return meta, res, nil
}

func (s *Github) GetUser(ctx context.Context, username string) (*github.User, *github.Response, error) {
	user, res, err := s.client.Users.Get(ctx, username)
	if err != nil {
		return nil, res, fmt.Errorf("failed to get user from GitHub: %w", err)
	}

	return user, res, nil
}

func (s *Github) GetUserByID(ctx context.Context, id int64) (*github.User, *github.Response, error) {
	user, res, err := s.client.Users.GetByID(ctx, id)
	if err != nil {
		return nil, res, fmt.Errorf("failed to get user by ID from GitHub: %w", err)
	}

	return user, res, nil
}

func (s *Github) GetUserFollowers(
	ctx context.Context,
	username string,
	page int,
	limit int,
) ([]*github.User, *github.Response, error) {
	opt := &github.ListOptions{
		Page:    page,
		PerPage: limit,
	}

	users, res, err := s.client.Users.ListFollowers(ctx, username, opt)
	if err != nil {
		return nil, res, fmt.Errorf("failed to get followers from GitHub: %w", err)
	}

	return users, res, nil
}

func (s *Github) CollectUserFollowers(
	ctx context.Context,
	username string,
	limit int,
) ([]*github.User, error) {
	var followers []*github.User

	page := 1
	logger := s.logger.With(slog.Any("username", username))

	for {
		logger.DebugContext(ctx, "fetching followers", slog.Int("page", page))

		users, res, err := s.GetUserFollowers(ctx, username, page, limit)
		if err != nil {
			return nil, fmt.Errorf("failed to get user followers: %w", err)
		}

		followers = append(followers, users...)

		if res.NextPage == 0 {
			logger.DebugContext(ctx, "last page reached")

			break
		}

		page = res.NextPage
	}

	return followers, nil
}

func (s *Github) loadPrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := os.ReadFile(s.config.GitHub.AppPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA key: %w", err)
	}

	return rsaKey, nil
}
