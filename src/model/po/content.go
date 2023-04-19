package po

type OpenAICategory struct {
	ID        int64
	Category  string
	Tokens    int
	Embedding []float32
}

type OpenAIContent struct {
	ID        int64
	Category  string
	Heading   string
	Content   string
	Tokens    int
	Embedding []float32
}
