package handler

import (
	"github.com/ekkx/tcmrsv-web/server/internal/modules/user/usecase"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
)

type HandlerImpl struct {
	useCase usecase.UseCase
}

func New(useCase usecase.UseCase) userv1connect.UserServiceHandler {
	return &HandlerImpl{
		useCase: useCase,
	}
}
