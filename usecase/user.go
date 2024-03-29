package usecase

import (
	"context"
	"errors"
	"fmt"
	"my-arch/domain"
)

func (u *Usecase) UserAdd(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := u.repos.User.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user; %w", err)
	}
	return user, nil
}

func (u *Usecase) UserGetAll(ctx context.Context) ([]*domain.User, error) {
	user, err := u.repos.User.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user; %w", err)
	}
	return user, nil
}

func (u *Usecase) GetNoteUsers(ctx context.Context, noteID int) ([]*domain.User, error) {
	users, err := u.repos.User.GetByNoteID(ctx, noteID)
	if err != nil {
		return nil, fmt.Errorf("failed to get users; %w", err)
	}

	return users, nil
}

func (u *Usecase) SignIn(ctx context.Context, email, password string) (*domain.User, error) {
	sessionState, err := getSessionState(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sessionState; %w", err)
	}

	user, err := u.repos.User.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user; %w", err)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	fmt.Printf("user: %v\n", user)
	fmt.Printf("password: %v\n", password)

	if user.Password != password {
		return nil, errors.New("invalid password") // TODO: userやパスワードが正しくないという事情を外にはもらさない
	}

	sessionState.UserID = user.ID

	return user, nil
}
