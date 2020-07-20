package foo

import (
"fmt"
"github.com/alexeykirinyuk/tech-task-go/data"
"github.com/alexeykirinyuk/tech-task-go/libs"
"github.com/google/uuid"
"github.com/jinzhu/gorm"
)

type fooService struct {
	storage fooStorage
}

func newService(dbProvider data.IDatabaseProvider) fooService {
	return fooService{storage: newStorage(dbProvider)}
}

func (b fooService) create(foo foo) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateFoo(foo); !ok {
		return
	}

	if err := b.storage.add(foo); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b fooService) get(id uuid.UUID) (item foo, errs []libs.ValidationError, ok bool) {
	item, err := b.storage.getById(id)
	if err == gorm.ErrRecordNotFound {
		return foo{}, libs.SingleValidationError("Id", fmt.Sprintf("foo with ID '%s' doesn't exists")), false
	}
	if err != nil {
		panic(err)
	}

	return item, libs.Valid(), true
}

func (b fooService) update(item foo) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateFoo(item); !ok {
		return
	}

	if err := b.storage.update(item); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func (b fooService) delete(id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateDeleteFoo(b.storage, id); !ok {
		return
	}

	if err := b.storage.delete(id); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}


func validateFoo(foo foo) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator().NotEmpty("Title", foo.Title)
	validator = validator.NotEmpty("Description", foo.Description)

	return validator.Errors, validator.Ok()
}

func validateDeleteFoo(storage fooStorage, id uuid.UUID) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator()
	validator.Must("Id", fmt.Sprintf("foo with ID '%s' doesn't exists.", id), func() bool {
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
