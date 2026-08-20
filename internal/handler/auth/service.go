package auth

import "context"

type Service interface {
	login(ctx context.Context) error
	logout(ctx context.Context) error
	register(ctx context.Context) error
	refresh(ctx context.Context) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) register(ctx context.Context) error {
	return nil
}

func (s *service) login(ctx context.Context) error {
	return nil
}

func (s *service) logout(ctx context.Context) error {
	return nil
}

func (s *service) refresh(ctx context.Context) error {
	return nil
}
