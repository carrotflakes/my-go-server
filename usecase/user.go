package usecase

import (
	"my-arch/domain"
)

func (u *Usecase) UserAdd(ctx Context, user *domain.User) (*domain.User, error) {
	return u.repos.User.Create(ctx, user)
}

func (u *Usecase) UserGetAll(ctx Context) ([]*domain.User, error) {
	return u.repos.User.GetAll(ctx)
}
