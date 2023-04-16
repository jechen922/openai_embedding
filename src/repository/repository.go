package repository

import (
	"openaigo/src/repository/seedrepo"
)

type ICore interface {
	Seed() seedrepo.ISeedRepo
}

type service struct {
	seedRepo seedrepo.ISeedRepo
}

func New() ICore {
	return &service{
		seedRepo: seedrepo.New(),
	}
}

func (s *service) Seed() seedrepo.ISeedRepo {
	return s.Seed()
}
