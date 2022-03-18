package registry

import (
	"my-arch/domain"
)

type Repository interface {
	NewUser() domain.UserRepository
	NewNote() domain.NoteRepository
}
