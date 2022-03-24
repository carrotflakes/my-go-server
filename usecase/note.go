package usecase

import (
	"context"
	"my-arch/domain"
)

func (u *Usecase) AddNote(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	_, err = u.repos.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	note, err = u.repos.Note.Create(ctx, note)
	if err != nil {
		return nil, err
	}

	err = u.repos.UserNote.Create(ctx, &domain.UserNote{
		UserID: userID,
		NoteID: note.ID,
	})
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (u *Usecase) GetAllNotes(ctx context.Context) ([]*domain.Note, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	return u.repos.Note.GetNotesByUserID(ctx, userID)
}
