package usecase

import (
	"context"
	"my-arch/domain"
)

func (u *Usecase) AddUserNote(ctx context.Context, userId int, noteId int) (*domain.UserNote, error) {
	userNote := &domain.UserNote{
		UserID: userId,
		NoteID: noteId,
	}
	err := u.repos.UserNote.Create(ctx, userNote)
	if err != nil {
		return nil, err
	}
	return userNote, nil
}
