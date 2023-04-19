package service

import (
	"openaigo/src/database"
	"openaigo/src/repository"
	"openaigo/src/service/customerservice"
	"openaigo/src/service/seedservice"
	"openaigo/src/utils/openai"
)

type ICore interface {
	Seed() seedservice.ISeedService
	CustomerService() customerservice.ICustomerService
}

type service struct {
	seed            seedservice.ISeedService
	customerService customerservice.ICustomerService
}

func New(db database.IDatabase, repo repository.ICore, ai openai.IAI) ICore {
	return &service{
		seed:            seedservice.NewSeed(db.Postgres(), repo),
		customerService: customerservice.NewCustomer(db.Postgres(), repo, ai),
	}
}

func (s *service) Seed() seedservice.ISeedService {
	return s.seed
}

func (s *service) CustomerService() customerservice.ICustomerService {
	return s.customerService
}
