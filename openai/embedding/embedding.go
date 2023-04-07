package embedding

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Section struct {
	Title   string
	Heading string
	Content string
	Vectors []float64
}

func client() *openai.Client {
	return openai.NewClient("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")
}

func Create(sections ...Section) ([]Section, error) {
	texts := make([]string, len(sections))
	for i, s := range sections {
		texts[i] = s.Content
	}
	embResp, err := client().CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: texts,
			Model: openai.AdaEmbeddingV2,
			User:  "somebody",
		})
	if err != nil {
		return nil, err
	}

	results := make([]Section, 0, len(sections))
	for _, datum := range embResp.Data {
		results = append(results, Section{
			Title:   sections[datum.Index].Title,
			Heading: sections[datum.Index].Heading,
			Content: sections[datum.Index].Content,
			Vectors: datum.Embedding,
		})
	}
	return results, nil
}
