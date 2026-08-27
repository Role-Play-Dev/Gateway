package auth

import (
	"context"
)

type Service interface {
	credentialsRegisterLinkSend(ctx context.Context, req CredentialsRegisterLinkSendRequest) error
	credentialsRegisterLinkConfirm(ctx context.Context, token string, req CredentialsRegisterLinkConfirmRequest) error
	credentialsLogin(ctx context.Context, req CredentialsLoginRequest) (res LoginResponce, refreshToken string, err error)
	tokenRefresh(ctx context.Context, refreshToken string) (res TokenRefreshResponce, err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) credentialsRegisterLinkSend(ctx context.Context, req CredentialsRegisterLinkSendRequest) error {
	return nil
}

func (s *service) credentialsRegisterLinkConfirm(ctx context.Context, token string, req CredentialsRegisterLinkConfirmRequest) error {
	return nil
}

func (s *service) credentialsLogin(ctx context.Context, req CredentialsLoginRequest) (res LoginResponce, refreshToken string, err error) {
	return LoginResponce{}, "", nil
}

func (s *service) tokenRefresh(ctx context.Context, refreshToken string) (res TokenRefreshResponce, err error) {
	return TokenRefreshResponce{}, nil
}
