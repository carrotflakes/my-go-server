package gateway

import (
	"context"
	"fmt"
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
	table, err := g.db.GetTable("user")
	if err != nil {
		return nil, fmt.Errorf("get user table failed; %w", err)
	}

	for _, v := range *table {
		if v.(*domain.User).ID == id {
			return v.(*domain.User), nil
		}
	}

	return nil, nil // not found
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

func (g *User) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	table, err := g.db.GetTable("user")
	if err != nil {
		return nil, err
	}

	for _, v := range *table {
		if v.(*domain.User).Email == email {
			return v.(*domain.User), nil
		}
	}

	return nil, nil // not found
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

	u := *user
	u.ID = id

	*table = append(*table, &u)
	fmt.Printf("user added: %v\n", u)

	return &u, nil
}

func (g *User) GetByNoteID(ctx context.Context, noteID int) ([]*domain.User, error) {
	userNoteTable, err := g.db.GetTable("userNote")
	if err != nil {
		return nil, fmt.Errorf("get userNote table failed; %w", err)
	}

	userTable, err := g.db.GetTable("user")
	if err != nil {
		return nil, fmt.Errorf("get user table failed; %w", err)
	}

	users := []*domain.User{}

	for _, v := range *userNoteTable {
		userNote := v.(*domain.UserNote)
		if userNote.NoteID == noteID {
			for _, v := range *userTable {
				user := v.(*domain.User)
				if user.ID == userNote.UserID {
					users = append(users, user)
				}
			}
		}
	}

	return users, nil
}
