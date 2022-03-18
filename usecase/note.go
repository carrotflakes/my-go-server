package usecase

import (
	"context"
	"my-arch/domain"
)

func (u *Usacase) NoteAdd(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	return u.noteRepo.Create(ctx, note)
}
