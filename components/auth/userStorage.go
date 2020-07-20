package auth

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type userStorage struct {
	db *gorm.DB
}

func (u userStorage) getByEmail(email string) (user User, err error) {
	if err = u.db.First(&user, "email = ?", email).Error; err != nil {
		err = fmt.Errorf("error when trying get templates by email: %s", err)
		return
	}

	return
}

func (u userStorage) getByConfirmSelector(selector string) (user User, err error) {
	if err = u.db.First(&user, "confirmed_selector = ?", selector).Error; err != nil {
		err = fmt.Errorf("error when trying get templates by confirmed_selector: %s", err)
		return
	}

	return
}

func (u userStorage) add(user User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return fmt.Errorf("error when trying crete new templates: %s", err)
	}

	return nil
}

func (u userStorage) update(user User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return fmt.Errorf("error when trying update templates with ID %s: %s",user.Id, err)
	}

	return nil
}
