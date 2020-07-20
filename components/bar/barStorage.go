package bar

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type barStorage struct {
	db *gorm.DB
}

func newStorage(dbProvider data.IDatabaseProvider) barStorage {
	db, err := dbProvider.CreateConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&bar{})

	return barStorage{db: db}
}

func (f barStorage) add(b bar) error {
	if err := f.db.Create(&b).Error; err != nil {
		return fmt.Errorf("error when trying crete new templates: %s", err)
	}

	return nil
}

func (f barStorage) getAll() (items []bar, err error) {
	err = f.db.Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get bars: %s", err)
		return
	}

	return
}

func (f barStorage) getById(id uuid.UUID) (Bar bar, err error) {
	err = f.db.First(&Bar, id).Error
	if err != nil {
		err = fmt.Errorf("error when trying get templates by Id: %s", err)
		return
	}

	return
}

func (f barStorage) update(b bar) error {
	if err := f.db.Update(&b).Error; err != nil {
		return fmt.Errorf("error when trying update templates with ID %s: %s", b.Id, err)
	}

	return nil
}

func (f barStorage) delete(id uuid.UUID) error {
	if err := f.db.Where("id = ?", id).Delete(&bar{}).Error; err != nil {
		return fmt.Errorf("error when trying delete templates with ID %s: %s", id, err)
	}

	return nil
}
