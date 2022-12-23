package graph

import (
	"context"
	"fmt"
	"my-arch/graph/model"
	"my-arch/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase *usecase.Usecase
}

func NewResolver(usecase *usecase.Usecase) *Resolver {
	return &Resolver{
		usecase,
	}
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	uCtx := Context{ctx}
	notes, err := r.usecase.GetAllNotes(uCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all notes; %w", err)
	}

	gNotes := []*model.Note{}

	for _, note := range notes {
		gNotes = append(gNotes, &model.Note{
			ID:   fmt.Sprintf("%d", note.ID),
			Text: note.Text,
			Done: false,
			User: &model.User{
				ID:   "dum",
				Name: "dum",
			},
		})
	}

	return gNotes, nil
}

type Context struct {
	context.Context
}

func (c Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	// TODO
}

func (c Context) GetCookie(name string) (string, error) {
	// TODO
	return "", fmt.Errorf("not implemented")
}
