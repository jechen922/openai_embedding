package repository

import (
	"openaigo/src/repository/contentrepo"
	"openaigo/src/repository/seedrepo"
)

type ICore interface {
	Seed() seedrepo.ISeedRepo
	Content() contentrepo.IContentRepo
}

type service struct {
	seedRepo    seedrepo.ISeedRepo
	contentRepo contentrepo.IContentRepo
}

func New() ICore {
	return &service{
		seedRepo:    seedrepo.New(),
		contentRepo: contentrepo.New(),
	}
}

func (s *service) Seed() seedrepo.ISeedRepo {
	return s.seedRepo
}

func (s *service) Content() contentrepo.IContentRepo {
	return s.contentRepo
}
