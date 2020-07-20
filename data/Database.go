package data

import (
	"github.com/jinzhu/gorm"
)

type IDatabaseProvider interface {
	CreateConnection() (*gorm.DB, error)
}