package auth

import (
	"context"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
)

type serverStore struct {
	storage userStorage
}

func NewServerStore(dbProvider data.IDatabaseProvider) *serverStore {
	db, err := dbProvider.CreateConnection()
	if err != nil {
		panic(err)
	}

	return &serverStore{storage: userStorage{db: db}}
}

func (s serverStore) LoadByConfirmSelector(_ context.Context, selector string) (authboss.ConfirmableUser, error) {
	user, err := s.storage.getByConfirmSelector(selector)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s serverStore) New(_ context.Context) authboss.User {
	return &User{
		Id:   uuid.New(),
		Role: RoleMember,
	}
}

func (s serverStore) Create(_ context.Context, user authboss.User) error {
	u := user.(*User)
	return s.storage.add(*u)
}

func (s serverStore) Load(_ context.Context, key string) (authboss.User, error) {
	user, err := s.storage.getByEmail(key)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s serverStore) Save(_ context.Context, user authboss.User) error {
	castedUser := user.(*User)
	return s.storage.update(*castedUser)
}

func (s *serverStore) Close() error {
	if err := s.Close(); err != nil {
		panic("обработаю эту ошибку коннекта к базе позже")
	}

	return nil
}
