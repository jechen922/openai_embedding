package calculator

func CosineSimilarity(v1, v2 []float32) float32 {
	// 計算兩個向量的內積
	dotProduct := float32(0.0)
	for i := 0; i < len(v1); i++ {
		dotProduct += v1[i] * v2[i]
	}

	// 計算向量的模長
	v1Length := float32(0.0)
	for _, x := range v1 {
		v1Length += x * x
	}
	v1Length = sqrt(v1Length)

	v2Length := float32(0.0)
	for _, x := range v2 {
		v2Length += x * x
	}
	v2Length = sqrt(v2Length)

	// 計算余弦相似度
	return dotProduct / (v1Length * v2Length)
}

func sqrt(x float32) float32 {
	if x == 0 {
		return 0
	}

	z := float32(1.0)
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}
