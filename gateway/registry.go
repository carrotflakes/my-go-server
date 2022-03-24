package gateway

import (
	"my-arch/domain"
	"my-arch/mydb"
)

func NewRepositories(db *mydb.MyDB) *domain.Repositories {
	return &domain.Repositories{
		User:     NewUser(db),
		Note:     NewNote(db),
		UserNote: NewUserNote(db),
	}
}
