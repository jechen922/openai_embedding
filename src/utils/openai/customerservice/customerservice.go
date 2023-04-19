package customerservice

import "github.com/sashabaranov/go-openai"

type IAI interface {
}

func New(cli *openai.Client) IAI {
	return &customerService{client: cli}
}

type customerService struct {
	client *openai.Client
}