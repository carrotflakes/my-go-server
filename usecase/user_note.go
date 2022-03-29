package usecase

import (
	"fmt"
	"my-arch/domain"
)

func (u *Usecase) AddUserNote(ctx Context, userId int, noteId int) (*domain.UserNote, error) {
	userNote := &domain.UserNote{
		UserID: userId,
		NoteID: noteId,
	}
	err := u.repos.UserNote.Create(ctx, userNote)
	if err != nil {
		return nil, fmt.Errorf("failed to create userNote; %w", err)
	}
	return userNote, nil
}
