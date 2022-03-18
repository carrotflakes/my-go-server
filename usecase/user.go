package usecase

import (
	"context"
	"my-arch/domain"
)

func (u *Usacase) UserAdd(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepo.Create(ctx, user)
}

func (u *Usacase) UserGetAll(ctx context.Context) ([]*domain.User, error) {
	return u.userRepo.GetAll(ctx)
}
