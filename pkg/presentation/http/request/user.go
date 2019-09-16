package request

import (
	"go-layered-architecture-example/pkg/application/input"
)

type User struct {
	Name string `json:"name"`
}

func (r *User) Write(in *input.UserCreate) {
	in.Name = r.Name
}
