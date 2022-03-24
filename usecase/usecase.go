package usecase

import (
	"my-arch/registry"
)

type Usecase struct {
	repos *registry.Repositories
}

func New(repoRegi *registry.Repositories) *Usecase {
	return &Usecase{
		repos: repoRegi,
	}
}
