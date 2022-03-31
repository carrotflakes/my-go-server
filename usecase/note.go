package usecase

import (
	"fmt"
	"my-arch/domain"
)

func (u *Usecase) AddNote(ctx Context, note *domain.Note) (*domain.Note, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get userID; %w", err)
	}

	_, err = u.repos.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user; %w", err)
	}

	note, err = u.repos.Note.Create(ctx, note)
	if err != nil {
		return nil, fmt.Errorf("failed to create note; %w", err)
	}

	err = u.repos.UserNote.Create(ctx, domain.NewUserNote(userID, note.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to create userNote; %w", err)
	}

	return note, nil
}

func (u *Usecase) GetAllNotes(ctx Context) ([]*domain.Note, error) {
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get userID; %w", err)
	}

	notes, err := u.repos.Note.GetNotesByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes; %w", err)
	}

	return notes, nil
}
