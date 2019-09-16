package registry

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-layered-architecture-example/pkg/domain/repository"
	"go-layered-architecture-example/pkg/infrastructure/mysql/gateway"
)

type Repository interface {
	NewUser() repository.User
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func createDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql:3306)/example?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	return db
}

func (*repositoryImpl) NewUser() repository.User {
	return gateway.NewUser(createDBConnection())
}
