package service_impl

import (
	"go-layered-architecture-example/pkg/domain/model"
	"go-layered-architecture-example/pkg/domain/repository"
)

type User struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) *User {
	return &User{userRepository: userRepository}
}

func (s *User) Create(m *model.User) error {
	if err := s.userRepository.Create(m); err != nil {
		return nil
	}

	return nil
}
