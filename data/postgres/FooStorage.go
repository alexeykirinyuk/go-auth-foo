package postgres

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type FooStorage struct {
	db *gorm.DB
}

func (f FooStorage) Add(foo models.Foo) error {
	if err := f.db.Create(&foo).Error; err != nil {
		return fmt.Errorf("error when trying crete new foo: %s", err)
	}

	return nil
}

func (f FooStorage) GetById(id uuid.UUID) (foo models.Foo, err error) {
	err = f.db.First(&foo, "id = ?", id).Error
	if err != nil {
		err = fmt.Errorf("error when trying get foo by Id: %s", err)
		return
	}

	return
}

func (f FooStorage) Update(foo models.Foo) error {
	if err := f.db.Save(&foo).Error; err != nil {
		return fmt.Errorf("error when trying update foo with ID %s: %s", foo.Id, err)
	}

	return nil
}

func (f FooStorage) Delete(id uuid.UUID) error {
	if err := f.db.Where("id = ?", id).Delete(&models.Foo{}).Error; err != nil {
		return fmt.Errorf("error when trying delete foo with ID %s: %s", id, err)
	}

	return nil
}

func (f FooStorage) GetAll() (items []models.Foo, err error) {
	err = f.db.Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get foos: %s", err)
		return
	}

	return
}