package postgres

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type SigmaStorage struct {
	db *gorm.DB
}


func (s SigmaStorage) Add(Sigma models.Sigma) error {
	if err := s.db.Create(Sigma).Error; err != nil {
		return fmt.Errorf("error when trying crete new sigma: %s", err)
	}

	return nil
}

func (s SigmaStorage) GetById(id uuid.UUID) (Sigma models.Sigma, err error) {
	err = s.db.First(&Sigma, id).Error
	if err != nil {
		err = fmt.Errorf("error when trying get sigma by Id: %s", err)
		return
	}

	return
}

func (s SigmaStorage) Update(Sigma models.Sigma) error {
	if err := s.db.Update(Sigma).Error; err != nil {
		return fmt.Errorf("error when trying update sigma with ID %s: %s", Sigma.Id, err)
	}

	return nil
}

func (s SigmaStorage) Delete(id uuid.UUID) error {
	if err := s.db.Where("id = ?", id).Delete(&models.Sigma{}); err != nil {
		return fmt.Errorf("error when trying delete sigma with ID %s: %s", id, err)
	}

	return nil
}
