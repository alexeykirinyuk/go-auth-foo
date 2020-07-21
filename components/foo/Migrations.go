package foo

import "github.com/jinzhu/gorm"

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&foo{}).Error
}
