package foo

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type fooStorage struct {
	db *gorm.DB
}

func newStorage(dbProvider data.IDatabaseProvider) fooStorage {
	db, err := dbProvider.CreateConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&foo{})

	return fooStorage{db: db}
}

func (f fooStorage) add(foo foo) error {
	if err := f.db.Create(&foo).Error; err != nil {
		return fmt.Errorf("error when trying crete new foo: %s", err)
	}

	return nil
}

func (f fooStorage) getById(id uuid.UUID) (foo foo, err error) {
	err = f.db.First(&foo, "id = ?", id).Error
	if err != nil {
		err = fmt.Errorf("error when trying get foo by Id: %s", err)
		return
	}

	return
}

func (f fooStorage) update(foo foo) error {
	if err := f.db.Save(&foo).Error; err != nil {
		return fmt.Errorf("error when trying update foo with ID %s: %s", foo.Id, err)
	}

	return nil
}

func (f fooStorage) delete(id uuid.UUID) error {
	if err := f.db.Where("id = ?", id).Delete(&foo{}).Error; err != nil {
		return fmt.Errorf("error when trying delete foo with ID %s: %s", id, err)
	}

	return nil
}

func (f fooStorage) getAll() (items []foo, err error) {
	err = f.db.Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get foos: %s", err)
		return
	}

	return
}