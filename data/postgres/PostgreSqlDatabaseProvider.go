package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSqlDatabaseProvider struct {
}

func NewProvider() PostgreSqlDatabaseProvider  {
	return PostgreSqlDatabaseProvider{}
}

func (p PostgreSqlDatabaseProvider) CreateConnection() (grm *gorm.DB, err error) {
	// TODO: move connection pattern string to configuration file
	grm, err = gorm.Open("postgres", "host=localhost port=5432 templates=postgres dbname=fun password=postgres sslmode=disable")
	if err != nil {
		err = fmt.Errorf("error when try to connect to postgresql db: %s", err)
		return
	}

	return
}
