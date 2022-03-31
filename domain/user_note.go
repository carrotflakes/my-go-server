package domain

import "context"

type UserNote struct {
	UserID int
	NoteID int
}

func NewUserNote(userID int, noteID int) *UserNote {
	return &UserNote{
		UserID: userID,
		NoteID: noteID,
	}
}

type UserNoteRepository interface {
	Create(ctx context.Context, userNote *UserNote) error
	Delete(ctx context.Context, userID int, noteID int) error
}
