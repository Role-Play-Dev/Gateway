package auth

import (
	"context"
)

type Service interface {
	register(ctx context.Context, email string) error
	verify(ctx context.Context, token string, username string, password string) error
	login(ctx context.Context, username string, password string) (accessToken string, refreshToken string, err error)
	refresh(ctx context.Context, refreshToken string) (accessToken string, err error)
	logout(ctx context.Context) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) register(ctx context.Context, email string) error {
	return nil
}

func (s *service) verify(ctx context.Context, token string, username string, password string) error {
	return nil
}

func (s *service) login(ctx context.Context, username string, password string) (accessToken string, refreshToken string, err error) {
	return "", "", nil
}

func (s *service) refresh(ctx context.Context, refreshToken string) (accessToken string, err error) {
	return "", nil
}

func (s *service) logout(ctx context.Context) error {
	return nil
}
