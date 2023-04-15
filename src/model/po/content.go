package po

type OpenAIContent struct {
	ID        int64
	Category  string
	Heading   string
	Content   string
	Tokens    int
	Embedding []float32
}
