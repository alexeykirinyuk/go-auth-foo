package userRole

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type userRoleStorage struct {
	db *gorm.DB
}

type userRole struct {
	Id uuid.UUID

	FirstName string
	LastName  string
	Role      string

	Email string
}

const tableName = "users"

func newStorage(dbProvider data.IDatabaseProvider) userRoleStorage {
	db, err := dbProvider.CreateConnection()
	if err != nil {
		panic(err)
	}

	return userRoleStorage{db: db}
}

func (u userRoleStorage) getById(id uuid.UUID) (item userRole, err error) {
	if err = u.db.Table(tableName).First(&item, "id = ?", id).Error; err != nil {
		return
	}

	return
}

func (u userRoleStorage) getAll() (items []userRole, err error) {
	err = u.db.Table(tableName).Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get bars: %s", err)
		return
	}

	return
}

func (u userRoleStorage) updateRole(id uuid.UUID, role string) error {
	if err := u.db.Table(tableName).Where("id = ?", id).Update("role", role).Error; err != nil {
		return fmt.Errorf("error when trying update templates with ID %s: %s", id, err)
	}

	return nil
}
