package bar

import "github.com/jinzhu/gorm"

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&bar{}).Error
}
