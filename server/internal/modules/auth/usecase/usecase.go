package usecase

import (
	"context"

	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/repository"
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/service"
	"github.com/ekkx/tcmrsv-web/server/pkg/jwt"
)

type UseCase interface {
	Authorize(ctx context.Context, params *AuthorizeInput) (*AuthorizeOutput, error)
}

type UseCaseImpl struct {
	jwtManager  *jwt.JWTManager
	userRepo    repository.Repository
	userService service.Service
}

func New(
	jwtManager *jwt.JWTManager,
	userRepo repository.Repository,
	userService service.Service,
) UseCase {
	return &UseCaseImpl{
		jwtManager:  jwtManager,
		userRepo:    userRepo,
		userService: userService,
	}
}
