package registry

import (
	"my-arch/domain"
	"my-arch/gateway"
	"my-arch/memdb"
)

type RepositoryImpl struct {
	db *memdb.MemDb
}

var _ Repository = (*RepositoryImpl)(nil)

func New(db *memdb.MemDb) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) NewUser() domain.UserRepository {
	return gateway.NewUser(r.db)
}

func (r *RepositoryImpl) NewNote() domain.NoteRepository {
	return gateway.NewNote(r.db)
}
