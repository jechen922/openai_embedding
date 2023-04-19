package openai

import (
	"openaigo/config"
	"openaigo/src/utils/openai/customerservice"

	"github.com/sashabaranov/go-openai"
)

type IAI interface {
	CustomerService() customerservice.IAI
}

func New(cfg config.IConfig) IAI {
	cli := openai.NewClient(cfg.GetSystemENV().ChatGPTToken)
	return &openAI{
		customerServiceAI: customerservice.New(cli),
	}
}

type openAI struct {
	customerServiceAI customerservice.IAI
}

func (ai *openAI) CustomerService() customerservice.IAI {
	return ai.customerServiceAI
}
