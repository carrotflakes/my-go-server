package usecase

import (
	"context"
	"my-arch/domain"
)

func (u *Usecase) UserAdd(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.repos.User.Create(ctx, user)
}

func (u *Usecase) UserGetAll(ctx context.Context) ([]*domain.User, error) {
	return u.repos.User.GetAll(ctx)
}
