package embedding

import (
	"context"
	"openaigo/src/model/po"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type Section struct {
	Title   string
	Heading string
	Content string
	Tokens  int
	Vectors []float32
}

func client() *openai.Client {
	return openai.NewClient("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")
}

func Category(category po.OpenAICategory) (po.OpenAICategory, error) {
	embResp, err := client().CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: strings.Split(category.Category, ""),
			Model: openai.AdaEmbeddingV2,
			User:  "somebody",
		})
	if err != nil {
		return po.OpenAICategory{}, err
	}

	ca := po.OpenAICategory{
		Category:  category.Category,
		Tokens:    embResp.Usage.TotalTokens,
		Embedding: embResp.Data[0].Embedding,
	}
	return ca, nil
}

func Content(section Section) (Section, error) {
	embResp, err := client().CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: strings.Split(section.Content, ""),
			Model: openai.AdaEmbeddingV2,
			User:  "somebody",
		})
	if err != nil {
		return Section{}, err
	}

	result := Section{
		Title:   section.Title,
		Heading: section.Heading,
		Content: section.Content,
		Tokens:  embResp.Usage.TotalTokens,
		Vectors: embResp.Data[0].Embedding,
	}
	return result, nil
}
