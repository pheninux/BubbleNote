package model

import "time"

type Note struct {
	ID        int
	Content   string
	CreatedAt time.Time
	Done      bool
	Archived  bool
}
