package userRole

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

const (
	RoleMember = "Member"
	RoleAdmin  = "Admin"
)

type userRoleService struct {
	storage userRoleStorage
}

func newService(dbProvider data.IDatabaseProvider) userRoleService {
	return userRoleService{storage: newStorage(dbProvider)}
}

func (b userRoleService) get(id uuid.UUID) (item userRole, errs []libs.ValidationError, ok bool) {
	item, err := b.storage.getById(id)
	if err == gorm.ErrRecordNotFound {
		return userRole{}, libs.SingleValidationError("Id", fmt.Sprintf("userRole with ID '%s' doesn't exists")), false
	}
	if err != nil {
		panic(err)
	}

	return item, libs.Valid(), true
}

func (b userRoleService) updateRole(id uuid.UUID, role string) (errs []libs.ValidationError, ok bool) {
	if errs, ok = validateUpdateRole(b.storage, id, role); !ok {
		return
	}

	if err := b.storage.updateRole(id, role); err != nil {
		panic(err)
	}

	return libs.Valid(), true
}

func validateUpdateRole(storage userRoleStorage, id uuid.UUID, role string) (errs []libs.ValidationError, ok bool) {
	validator := libs.NewValidator()
	validator = validator.NotEmpty("Role", role)

	msg := fmt.Sprintf("The role must be either '%s' or '%s'", RoleAdmin, RoleMember)
	validator = validator.Must("Role", msg, func() bool {
		return role == RoleMember || role == RoleAdmin
	})
	validator = validator.Must("Id", "User doesn't exists", func() bool {
		_, err := storage.getById(id)
		if err == gorm.ErrRecordNotFound {
			return false
		}
		if err != nil{
			panic(err)
		}

		return true
	})

	return validator.Errors, validator.Ok()
}
