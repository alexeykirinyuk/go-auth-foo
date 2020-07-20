package sigma

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type sigmaService struct {
	storage sigmaStorage
}

func newService(dbProvider data.IDatabaseProvider) sigmaService {
	return sigmaService{storage: newStorage(dbProvider)}
}

func (b sigmaService) create(sigma sigma) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateSigma(sigma); !ok {
		return
	}

	if err := b.storage.add(sigma); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b sigmaService) get(id uuid.UUID) (item sigma, errs []libs.ValidationError, ok bool) {
	item, err := b.storage.getById(id)
	if err == gorm.ErrRecordNotFound {
		return sigma{}, libs.SingleValidationError("Id", fmt.Sprintf("sigma with ID '%s' doesn't exists")), false
	}
	if err != nil {
		panic(err)
	}

	return item, libs.Valid(), true
}

func (b sigmaService) update(item sigma) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateSigma(item); !ok {
		return
	}

	if err := b.storage.update(item); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b sigmaService) delete(id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateDeleteSigma(b.storage, id); !ok {
		return
	}

	if err := b.storage.delete(id); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}


func validateSigma(sigma sigma) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator().NotEmpty("Info", sigma.Info)

	return validator.Errors, validator.Ok()
}

func validateDeleteSigma(storage sigmaStorage, id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator()
	validator.Must("Id", fmt.Sprintf("sigma with ID '%s' doesn't exists.", id), func() bool {
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
