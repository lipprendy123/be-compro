package service

import (
	"compro/config"
	"compro/internal/adapter/repository"
	"compro/internal/core/domain/entity"
	"context"
)

var (
	err  error
	code string
)

type UserServiceInterface interface {
	LoginAdmin(ctx context.Context, req entity.UserEntity) (string, error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
	cfg      *config.Config
}
