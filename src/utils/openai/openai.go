package openai

import (
	"github.com/sashabaranov/go-openai"
	"openaigo/config"
	"openaigo/src/utils/openai/customerservice"
)

type IAI interface {
	CustomerService() customerservice.IAI
}

func New(cfg config.IConfig) IAI {
	cli := openai.NewClient(cfg.GetSystemENV().ChatGPTToken)
	return &ai{
		customerServiceAI: customerservice.New(cli),
	}
}

type ai struct {
	customerServiceAI customerservice.IAI
}

func (a *ai) CustomerService() customerservice.IAI {
	return a.customerServiceAI
}
