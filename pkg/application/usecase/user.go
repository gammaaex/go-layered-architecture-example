package usecase

import (
	"go-layered-architecture-example/pkg/application/input"
	"go-layered-architecture-example/pkg/application/output"
	"go-layered-architecture-example/pkg/domain/model"
	"go-layered-architecture-example/pkg/domain/repository"
)

type User struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) *User {
	return &User{userRepository: userRepository}
}

func (u *User) Create(in *input.UserCreate) (*output.UserCreate, error) {
	m := &model.User{
		Name: in.Name,
	}

	if err := u.userRepository.Create(m); err != nil {
		return nil, err
	}

	return &output.UserCreate{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}
