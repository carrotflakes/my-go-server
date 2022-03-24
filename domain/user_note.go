package domain

import "context"

type UserNote struct {
	UserID int
	NoteID int
}

type UserNoteRepository interface {
	Create(ctx context.Context, userNote *UserNote) error
	Delete(ctx context.Context, userID int, noteID int) error
}
