package postgres

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func (u UserStorage) GetByEmail(email string) (user models.User, err error) {
	if err = u.db.First(&user, "email = ?", email).Error; err != nil {
		err = fmt.Errorf("error when trying get user by email: %s", err)
		return
	}

	return
}

func (u UserStorage) GetById(id uuid.UUID) (user models.User, err error) {
	if err = u.db.First(&user, "id = ?", id).Error; err != nil {
		err = fmt.Errorf("error when trying get user by id: %s", err)
		return
	}

	return
}

func (u UserStorage) GetByConfirmSelector(selector string) (user models.User, err error) {
	if err = u.db.First(&user, "confirmed_selector = ?", selector).Error; err != nil {
		err = fmt.Errorf("error when trying get user by confirmed_selector: %s", err)
		return
	}

	return
}

func (u UserStorage) GetAll() (items []models.User, err error) {
	err = u.db.Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get bars: %s", err)
		return
	}

	return
}

func (u UserStorage) Add(user models.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return fmt.Errorf("error when trying crete new user: %s", err)
	}

	return nil
}

func (u UserStorage) Update(user models.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return fmt.Errorf("error when trying update user with ID %s: %s",user.Id, err)
	}

	return nil
}
