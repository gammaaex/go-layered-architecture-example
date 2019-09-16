package response

import (
	"go-layered-architecture-example/pkg/application/output"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(out *output.UserCreate) *User {
	return &User{
		ID:        out.ID,
		Name:      out.Name,
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}
}
