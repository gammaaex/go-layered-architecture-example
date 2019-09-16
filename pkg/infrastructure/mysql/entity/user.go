package entity

import (
	"time"

	"go-layered-architecture-example/pkg/domain/model"
)

type User struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(m *model.User) *User {
	return &User{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (e *User) Write(m *model.User) {
	m.ID = e.ID
	m.Name = e.Name
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}
