package domain

import (
	"context"
	"time"
)

type Note struct {
	ID        int
	Text      string
	CreatedAt time.Time
}

type NoteRepository interface {
	Get(ctx context.Context, id int) (*Note, error)
	Create(ctx context.Context, note *Note) (*Note, error)

	GetNotesByUserID(ctx context.Context, userID int) ([]*Note, error)
}
