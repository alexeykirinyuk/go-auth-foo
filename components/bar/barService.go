package bar

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type barService struct {
	storage barStorage
}

func newService(dbProvider data.IDatabaseProvider) barService {
	return barService{storage: newStorage(dbProvider)}
}

func (b barService) create(bar bar) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateBar(bar); !ok {
		return
	}

	if err := b.storage.add(bar); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b barService) get(id uuid.UUID) (item bar, errs []libs.ValidationError, ok bool) {
	item, err := b.storage.getById(id)
	if err == gorm.ErrRecordNotFound {
		return bar{}, libs.SingleValidationError("Id", fmt.Sprintf("bar with ID '%s' doesn't exists")), false
	}
	if err != nil {
		panic(err)
	}

	return item, libs.Valid(), true
}

func (b barService) update(bar bar) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateBar(bar); !ok {
		return
	}

	if err := b.storage.update(bar); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b barService) delete(id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateDeleteBar(b.storage, id); !ok {
		return
	}

	if err := b.storage.delete(id); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}


func validateBar(bar bar) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator().NotEmpty("Title", bar.Title)
	validator = validator.NotEmpty("Description", bar.Address)
	validator = validator.NotEmpty("Address", bar.Address)

	return validator.Errors, validator.Ok()
}

func validateDeleteBar(storage barStorage, id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator()
	validator.Must("Id", fmt.Sprintf("bar with ID '%s' doesn't exists.", id), func() bool {
		_, err := storage.getById(id)
		if err == gorm.ErrRecordNotFound {
			return false
		}
		if err != nil {
			panic(err)
		}

		return true
	})

	return validator.Errors, validator.Ok()
}
