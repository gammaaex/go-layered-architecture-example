package service

import "go-layered-architecture-example/pkg/domain/model"

type User interface {
	Create(*model.User) error
}
