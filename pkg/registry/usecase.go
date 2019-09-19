package registry

import (
	"go-layered-architecture-example/pkg/application/usecase"
	"go-layered-architecture-example/pkg/application/usecase_impl"
	"go-layered-architecture-example/pkg/domain/service"
)

type Usecase interface {
	NewUser(service.User) usecase.User
}

type usecaseImpl struct {
}

func NewUsecase() Usecase {
	return &usecaseImpl{}
}

func (*usecaseImpl) NewUser(userService service.User) usecase.User {
	return usecase_impl.NewUser(userService)
}
