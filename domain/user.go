package domain

import "context"

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Get(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}
