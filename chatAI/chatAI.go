package chatAI

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Ask(question string) string {
	client := openai.NewClient("")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: "你是弈樂科技的客服系統"},
				{Role: openai.ChatMessageRoleSystem, Content: "你的名字叫做小天"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬公司名稱:弈樂科技股份有限公司"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬公司代表人:許順宗"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬公司地址:407台中市西屯區安和東路3號6樓"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬公司電話:04-23592667"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬資本總額(元):500,000,000"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬實收資本額(元):170,000,000"},
				{Role: openai.ChatMessageRoleSystem, Content: "所屬公司簡介:弈樂科技成立於2018年，由一群熱愛遊戲、專注研發及運營行銷的專業營運團隊所組成，秉持著「創新、用心」的經營理念，致力於最優質的遊戲開發。我們以提供現代人線上休閒娛樂管道為本，持續投入先進研發技術，開發多款具備娛樂感的遊戲，將休閒娛樂的氛圍帶給玩家，提供玩家多元與極致的遊戲體驗。我們將以把產品帶向全球化為目標，讓弈樂成為首屈一指的國際知名線上遊戲公司 !"},
				{Role: openai.ChatMessageRoleSystem, Content: `所屬公司事蹟:
2018/06: 弈樂成立、包你發娛樂城雙平台上線
2018/08: 包你發娛樂城 AppStore 平台第一名 (Casino Games/娛樂場遊戲類遊戲比較)
2020/03: 包你發娛樂城 GooglePlay 平台第一名 (Casino Games/娛樂場遊戲類遊戲比較)
2021/11: 包你發娛樂城獲得由「數位遊戲產業自律推動委員會」舉辦之友善遊戲軟體評選
2022/05: 弈樂科技運用 MongoDB Atlas 搶攻全球遊戲商機
2023/02: 包你發娛樂城 連續榮獲「友善遊戲環境指標」認證
`},
				{Role: openai.ChatMessageRoleSystem, Content: "公司未來願景:以提供現代人線上休閒娛樂管道為本，持續投入先進研發技術，開發更多產品，將我們的產品帶向全球化，成為首屈一指的國際知名線上遊戲公司。"},
				{Role: openai.ChatMessageRoleSystem, Content: "企業文化:弈樂秉持以人為本，塑造正向積極，運動健康的企業文化，希望夥伴們工作、健康、家庭三方平衡，Work Hard Play Hard !我們追求創新，激發跳耀思維，走在技術尖端。我們樂於學習，保持汲取新知，鼓勵實踐理論 。我們重視團隊，尊重不同聲音，堅信眾志成城。我們無畏挑戰，面對市場變遷，開創自我價值。"},
				{Role: openai.ChatMessageRoleSystem, Content: `公司福利:
幸福職場:免費供應午餐及晚餐、咖啡飲料及零食區
運動健康:員工專屬健身房、平日下午運動時間、多樣化運動性質社團
福利補助:新人到職即預給特休、年度旅遊補助（享有資格依職務別不同）、結婚禮金、生育補助、傷病慰問、生日禮金、喪葬慰問
學習成長:讀書日、員工教育訓練補助、不定期海外觀展
健康守護:除勞健保外，另有團體保險(含壽險、意外險、醫療險、癌症險)、健康檢查、職護諮詢服務
凝聚感情:部門聚餐補助、不定期舉辦大型活動(中秋烤肉、聖誕節交換禮物、除夕圍爐、員工旅遊)`},
				{Role: openai.ChatMessageRoleSystem, Content: "公司產品:包你發娛樂城、聚寶Online、Pocket Casino"},
				{Role: openai.ChatMessageRoleSystem, Content: `公司相關資訊:
官網:https://www.yile.com.tw/
facebook:https://www.facebook.com/YileTechnology/
Instagram:https://www.instagram.com/yiletechnology/
youtube:https://www.youtube.com/channel/UCoaw-TsMrYYKlDJAhl6CVUQ
linkedin:https://pse.is/3jk42e`},
				{Role: openai.ChatMessageRoleUser, Content: question},
			},
		},
	)
	if err != nil {
		return err.Error()
	}
	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

func AskTest(question string) string {
	client := openai.NewClient("")
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:       openai.GPT3TextDavinci003,
			Prompt:      question,
			Suffix:      "",
			MaxTokens:   300,
			Temperature: 0,
			User:        openai.ChatMessageRoleUser,
		})
	if err != nil {
		return err.Error()
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text
}

var embeddingInputs = []string{
	"所屬公司名稱:弈樂科技股份有限公司",
	"所屬公司代表人:許順宗",
	"所屬公司地址:407台中市西屯區安和東路3號6樓",
	"所屬公司電話:04-23592667",
	"所屬資本總額(元):500,000,000",
	"所屬實收資本額(元):170,000,000",
}

func CreateEmbeddings() string {
	client := openai.NewClient("")
	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: embeddingInputs,
			Model: openai.AdaEmbeddingV2,
			User:  openai.ChatMessageRoleSystem,
		},
	)
	if err != nil {
		return err.Error()
	}
	fmt.Println(fmt.Sprintf("%+v", resp))
	return "success"
}
