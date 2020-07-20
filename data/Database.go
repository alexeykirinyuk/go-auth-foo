package data

import (
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
)

type IUserStorage interface {
	Add(user models.User) error
	GetByEmail(email string) (models.User, error)
	GetById(id uuid.UUID) (models.User, error)
	GetByConfirmSelector(selector string) (models.User, error)
	Update(user models.User) error
	GetAll() ([]models.User, error)
}

type IFooStorage interface {
	Add(foo models.Foo) error
	GetById(id uuid.UUID) (models.Foo, error)
	GetAll() (items []models.Foo, err error)
	Update(foo models.Foo) error
	Delete(id uuid.UUID) error
}

type IBarStorage interface {
	Add(foo models.Bar) error
	GetById(id uuid.UUID) (models.Bar, error)
	Update(foo models.Bar) error
	Delete(id uuid.UUID) error
	GetAll() ([]models.Bar, error)
}

type ISigmaStorage interface {
	Add(foo models.Sigma) error
	GetById(id uuid.UUID) (models.Sigma, error)
	Update(foo models.Sigma) error
	Delete(id uuid.UUID) error
}

type IDatabase interface {
	GetUserStorage() IUserStorage
	GetFooStorage() IFooStorage
	GetBarStorage() IBarStorage
	GetSigmaStorage() ISigmaStorage
}
