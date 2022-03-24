package registry

import (
	"my-arch/domain"
)

type Repositories struct {
	User     domain.UserRepository
	Note     domain.NoteRepository
	UserNote domain.UserNoteRepository
}
