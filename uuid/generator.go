package uuid

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() (uuid.UUID) {
	uuid := uuid.NewV4()

	return uuid
}