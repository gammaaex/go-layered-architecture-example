package service

import (
	"go-layered-architecture-example/pkg/application/input"
	"go-layered-architecture-example/pkg/application/output"
)

type User interface {
	Create(*input.UserCreate) (*output.UserCreate, error)
}
