package postgres

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BarStorage struct {
	db *gorm.DB
}
func (f BarStorage) Add(b models.Bar) error {
	if err := f.db.Create(&b).Error; err != nil {
		return fmt.Errorf("error when trying crete new bar: %s", err)
	}

	return nil
}

func (f BarStorage) GetAll() (items []models.Bar, err error) {
	err = f.db.Find(&items).Error
	if err != nil {
		err = fmt.Errorf("error when trying get bars: %s", err)
		return
	}

	return
}


func (f BarStorage) GetById(id uuid.UUID) (Bar models.Bar, err error) {
	err = f.db.First(&Bar, id).Error
	if err != nil {
		err = fmt.Errorf("error when trying get bar by Id: %s", err)
		return
	}

	return
}

func (f BarStorage) Update(b models.Bar) error {
	if err := f.db.Update(&b).Error; err != nil {
		return fmt.Errorf("error when trying update bar with ID %s: %s", b.Id, err)
	}

	return nil
}

func (f BarStorage) Delete(id uuid.UUID) error {
	if err := f.db.Where("id = ?", id).Delete(&models.Bar{}).Error; err != nil {
		return fmt.Errorf("error when trying delete bar with ID %s: %s", id, err)
	}

	return nil
}

