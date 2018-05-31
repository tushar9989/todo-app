package models

import (
	"time"
)

type Task struct {
	ListID string
	Name string
	ID string
	Description string
	CreatedOn time.Time
	UpdatedOn time.Time
	Completed bool
}
