package domain

import "context"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

type UserRepository interface {
	Get(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByNoteID(ctx context.Context, noteID int) ([]*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}
