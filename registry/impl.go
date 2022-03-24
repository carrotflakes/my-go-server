package registry

import (
	"my-arch/gateway"
	"my-arch/mydb"
)

func NewRepositories(db *mydb.MyDB) *Repositories {
	return &Repositories{
		User:     gateway.NewUser(db),
		Note:     gateway.NewNote(db),
		UserNote: gateway.NewUserNote(db),
	}
}
