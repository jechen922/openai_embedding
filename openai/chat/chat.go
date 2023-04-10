package chat

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var defaultChatMessage = []openai.ChatCompletionMessage{
	{Role: openai.ChatMessageRoleSystem, Content: "I want you act as the customer service of 弈樂科技."},
	{Role: openai.ChatMessageRoleSystem, Content: "Your name is 小P."},
	{Role: openai.ChatMessageRoleSystem, Content: "I will ask you many question about your company. You will answer the question accurately in Traditional Chinese"},
}

func client() *openai.Client {
	return openai.NewClient("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")
}

func GetTitle(question string) string {
	content := fmt.Sprintf(`%s
Please analyze which of the following answers the above description belongs to, you only need to choose the answer, no explanation is required
- 個人資料
- 公司
- 出差管理辦法
- unknown`, question)
	resp, err := client().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: content},
			},
		},
	)
	if err != nil {
		return err.Error()
	}
	fmt.Println("content", content)
	fmt.Println(resp.Choices[0].Message.Content)
	//re := regexp.MustCompile(`\((.*?)\)`)
	//match := re.FindStringSubmatch(resp.Choices[0].Message.Content)
	//if len(match) > 1 {
	//	return match[1]
	//}
	return resp.Choices[0].Message.Content
}

func AnswerPersonalQuestion(question string) string {
	resp, err := client().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: append(
				defaultChatMessage,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				}),
		},
	)
	if err != nil {
		return err.Error()
	}
	return resp.Choices[0].Message.Content
}

func Chat(question string, provideAnswers ...string) string {
	similarAnswers := "context:"
	for _, m := range provideAnswers {
		similarAnswers += fmt.Sprintf("%s ", m)
	}
	completionMessages := append(defaultChatMessage,
		openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: `"You are a customer service,follow the highest principles,
if the answer is not contained within the context,do not say anything, just say "抱歉，您的問題已超出我可回覆範圍".`},
		openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: similarAnswers},
		openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: question},
	)
	resp, err := client().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: completionMessages,
		},
	)
	if err != nil {
		return err.Error()
	}
	return resp.Choices[0].Message.Content
}
