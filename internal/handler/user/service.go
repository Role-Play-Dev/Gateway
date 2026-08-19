package user

import "context"

type Service interface {
	get(ctx context.Context) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) get(ctx context.Context) error {
	return nil
}
