package auth

import (
	"context"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
)

type ServerStore struct {
	store data.IUserStorage
}

func (s ServerStore) LoadByConfirmSelector(ctx context.Context, selector string) (authboss.ConfirmableUser, error) {
	user, err := s.store.GetByConfirmSelector(selector)
	if err != nil{
		return nil, err
	}

	return &user, nil
}

func NewServerStore(db data.IUserStorage) *ServerStore {
	return &ServerStore{store: db}
}

func (s ServerStore) New(ctx context.Context) authboss.User {
	return &models.User{
		Id: uuid.New(),
		Role: models.RoleMember,
	}
}

func (s ServerStore) Create(ctx context.Context, user authboss.User) error {
	u := user.(*models.User)
	return s.store.Add(*u)
}

func (s ServerStore) Load(ctx context.Context, key string) (authboss.User, error) {
	user, err := s.store.GetByEmail(key)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s ServerStore) Save(ctx context.Context, user authboss.User) error {
	castedUser := user.(*models.User)
	return s.store.Update(*castedUser)
}

func (s *ServerStore) Close() error {
	// TODO: create own error
	if err := s.Close(); err != nil {
		panic("обработаю эту ошибку коннекта к базе позже")
	}

	return nil
}
