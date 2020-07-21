package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSqlDatabaseProvider struct {
	connectionString string
}

func NewProvider(connectionString string) PostgreSqlDatabaseProvider {
	return PostgreSqlDatabaseProvider{
		connectionString: connectionString,
	}
}

func (p PostgreSqlDatabaseProvider) CreateConnection() (grm *gorm.DB, err error) {
	grm, err = gorm.Open("postgres", p.connectionString)
	if err != nil {
		err = fmt.Errorf("error when try to connect to postgresql db: %s", err)
		return
	}

	return
}
