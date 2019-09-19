package registry

import (
	"go-layered-architecture-example/pkg/domain/repository"
	"go-layered-architecture-example/pkg/domain/service"
	"go-layered-architecture-example/pkg/domain/service_impl"
)

type Service interface {
	NewUser(repository.User) service.User
}

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

func (*serviceImpl) NewUser(userRepository repository.User) service.User {
	return service_impl.NewUser(userRepository)
}
