package graph

import (
	"my-arch/domain"
	"my-arch/pubsub"
	"my-arch/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repos   *domain.Repositories
	pubsub  *pubsub.Pubsub
	usecase *usecase.Usecase
}

func NewResolver(repos *domain.Repositories, pubsub *pubsub.Pubsub, usecase *usecase.Usecase) *Resolver {
	return &Resolver{
		repos,
		pubsub,
		usecase,
	}
}
