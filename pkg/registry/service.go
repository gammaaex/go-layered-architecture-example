package registry

import (
	"go-layered-architecture-example/pkg/application/usecase"
	"go-layered-architecture-example/pkg/domain/repository"
	"go-layered-architecture-example/pkg/domain/service"
)

type Service interface {
	NewUser(userRepository repository.User) service.User
}

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

func (*serviceImpl) NewUser(userRepository repository.User) service.User {
	return usecase.NewUser(userRepository)
}
