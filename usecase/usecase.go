package usecase

import (
	"my-arch/domain"
	"my-arch/registry"
)

type Usacase struct {
	userRepo domain.UserRepository
	noteRepo domain.NoteRepository
}

func New(repoRegi registry.Repository) *Usacase {
	return &Usacase{
		userRepo: repoRegi.NewUser(),
		noteRepo: repoRegi.NewNote(),
	}
}
