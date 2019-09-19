package repository_impl

import (
	"github.com/jinzhu/gorm"
	"go-layered-architecture-example/pkg/domain/model"
	"go-layered-architecture-example/pkg/infrastructure/mysql/entity"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (r *User) Create(m *model.User) error {
	e := entity.NewUser(m)

	if err := r.db.Create(&e).Error; err != nil {
		return err
	}

	e.Write(m)

	return nil
}
