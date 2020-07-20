package models

import (
	"github.com/google/uuid"
	"time"
)

type Bar struct {
	Id uuid.UUID

	Title string
	Description string
	Address string
	OpeningDate time.Time
}
