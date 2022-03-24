package usecase

import "my-arch/domain"

type Usecase struct {
	repos *domain.Repositories
}

func New(repoRegi *domain.Repositories) *Usecase {
	return &Usecase{
		repos: repoRegi,
	}
}
