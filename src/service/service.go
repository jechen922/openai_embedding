package service

import (
	"openaigo/src/database"
	"openaigo/src/service/seedservice"
)

type ICore interface {
	Seed() seedservice.ISeedService
}

type service struct {
	seedService seedservice.ISeedService
}

func New(db database.IDatabase) ICore {
	return &service{
		seedService: seedservice.NewSeed(db.Postgres()),
	}
}

func (s *service) Seed() seedservice.ISeedService {
	return s.seedService
}
