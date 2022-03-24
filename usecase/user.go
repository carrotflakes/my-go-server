package usecase

import (
	"errors"
	"fmt"
	"my-arch/domain"
	"strconv"
)

func (u *Usecase) UserAdd(ctx Context, user *domain.User) (*domain.User, error) {
	return u.repos.User.Create(ctx, user)
}

func (u *Usecase) UserGetAll(ctx Context) ([]*domain.User, error) {
	return u.repos.User.GetAll(ctx)
}

func (u *Usecase) SignIn(ctx Context, email, password string) (*domain.User, error) {
	user, err := u.repos.User.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	fmt.Printf("user: %v\n", user)
	fmt.Printf("password: %v\n", password)

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	ctx.SetCookie("userID", strconv.Itoa(user.ID), 60*60*24, "/", "", false, false)

	return user, nil
}
