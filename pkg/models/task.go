package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CompleteBy time.Time `json:"complete_by"`
}

type Tasks []Task
