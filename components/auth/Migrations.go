package auth

import (
	"github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}
