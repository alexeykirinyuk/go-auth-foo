package models

import "github.com/google/uuid"

type Foo struct {
	Id uuid.UUID

	Title string
	Description string
}
