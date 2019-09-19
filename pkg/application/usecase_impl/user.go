package usecase_impl

import (
	"go-layered-architecture-example/pkg/application/input"
	"go-layered-architecture-example/pkg/application/output"
	"go-layered-architecture-example/pkg/domain/model"
	"go-layered-architecture-example/pkg/domain/service"
)

type User struct {
	userService service.User
}

func NewUser(userService service.User) *User {
	return &User{userService: userService}
}

func (uc *User) Create(in *input.UserCreate) (*output.UserCreate, error) {
	m := &model.User{
		Name: in.Name,
	}

	if err := uc.userService.Create(m); err != nil {
		return nil, err
	}

	return &output.UserCreate{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}
