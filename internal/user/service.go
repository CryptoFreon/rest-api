package user

import (
	"context"
	"rest-api/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context) (u User, err error) {
	//TODO for next one
	return
}
