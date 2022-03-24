package gateway

import (
	"context"
	"my-arch/domain"
	"my-arch/mydb"
)

type User struct {
	db *mydb.MyDB
}

var _ domain.UserRepository = (*User)(nil)

func NewUser(db *mydb.MyDB) *User {
	return &User{
		db: db,
	}
}

func (g *User) Get(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

func (g *User) GetAll(ctx context.Context) ([]*domain.User, error) {
	table, err := g.db.GetTable("user")
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(*table))
	for i, v := range *table {
		users[i] = v.(*domain.User)
	}

	return users, nil
}

func (g *User) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	table, err := g.db.GetTable("user")
	if err != nil {
		return nil, err
	}

	// IDのauto increment
	id := 0
	for _, v := range *table {
		user := v.(*domain.User)
		if id < user.ID+1 {
			id = user.ID + 1
		}
	}

	user = &domain.User{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}

	*table = append(*table, user)

	return user, nil
}
