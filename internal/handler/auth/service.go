package auth

import (
	"context"
)

type Service interface {
	CredentialsRegisterLinkSend(ctx context.Context, req CredentialsRegisterLinkSendRequest) error
	CredentialsRegisterLinkConfirm(ctx context.Context, token string, req CredentialsRegisterLinkConfirmRequest) error
	CredentialsLogin(ctx context.Context, req CredentialsLoginRequest) (res LoginResponce, refreshToken string, err error)
	TokenRefresh(ctx context.Context, refreshToken string) (res TokenRefreshResponce, err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) CredentialsRegisterLinkSend(ctx context.Context, req CredentialsRegisterLinkSendRequest) error {
	return nil
}

func (s *service) CredentialsRegisterLinkConfirm(ctx context.Context, token string, req CredentialsRegisterLinkConfirmRequest) error {
	return nil
}

func (s *service) CredentialsLogin(ctx context.Context, req CredentialsLoginRequest) (res LoginResponce, refreshToken string, err error) {
	return LoginResponce{}, "", nil
}

func (s *service) TokenRefresh(ctx context.Context, refreshToken string) (res TokenRefreshResponce, err error) {
	return TokenRefreshResponce{}, nil
}
