package usecase

import (
	"context"
	"fmt"
	"my-arch/domain"
)

func (u *Usecase) AddNote(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	sessionState, err := getSessionState(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sessionState; %w", err)
	}

	_, err = u.repos.User.Get(ctx, sessionState.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user; %w", err)
	}

	note, err = u.repos.Note.Create(ctx, note)
	if err != nil {
		return nil, fmt.Errorf("failed to create note; %w", err)
	}

	err = u.repos.UserNote.Create(ctx, domain.NewUserNote(sessionState.UserID, note.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to create userNote; %w", err)
	}

	return note, nil
}

func (u *Usecase) GetAllNotes(ctx context.Context) ([]*domain.Note, error) {
	sessionState, err := getSessionState(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sessionState; %w", err)
	}

	notes, err := u.repos.Note.GetNotesByUserID(ctx, sessionState.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get notes; %w", err)
	}

	return notes, nil
}
