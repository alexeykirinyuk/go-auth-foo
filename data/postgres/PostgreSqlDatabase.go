package postgres

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSqlDatabase struct {
	db *gorm.DB
}

func (p PostgreSqlDatabase) GetFooStorage() data.IFooStorage {
	return FooStorage{db: p.db}
}

func (p PostgreSqlDatabase) GetBarStorage() data.IBarStorage {
	return BarStorage{db: p.db}
}

func (p PostgreSqlDatabase) GetSigmaStorage() data.ISigmaStorage {
	return SigmaStorage{db: p.db}
}

func (p PostgreSqlDatabase) GetUserStorage() data.IUserStorage {
	return UserStorage{db: p.db}
}

func NewPostgresDatabase() (db data.IDatabase, err error) {
	// TODO: move connection string to configuration file
	// TODO: move set up db to data module
	grm, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=fun password=postgres sslmode=disable")
	if err != nil {
		err = fmt.Errorf("error when try to connect to postgresql db: %s", err)
		return
	}

	err = grm.AutoMigrate(&models.User{}, &models.Foo{}, &models.Bar{}, &models.Sigma{}).Error
	if err != nil {
		err = fmt.Errorf("error when run db migrations: %s", err)
		return
	}

	db = &PostgreSqlDatabase{db: grm}

	return
}
