package bar

import (
	"github.com/google/uuid"
	"time"
)

type bar struct {
	Id uuid.UUID

	Title string
	Description string
	Address string
	OpeningDate time.Time
}
