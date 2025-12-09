package matrix

func SumVectors(vector1, vector2 []float64) []float64 {
	panicVectorSize(vector1, vector2)

	newVector := make([]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		newVector[i] = vector1[i] + vector2[i]
	}

	return newVector
}

func DotProduct(vector1, vector2 []float64) float64 {
	panicVectorSize(vector1, vector2)

	var result float64 = 0

	for i := 0; i < len(vector1); i++ {
		result += vector1[i] * vector2[i]
	}

	return result
}

func panicVectorSize(vector1, vector2 []float64) {
	if len(vector1) != len(vector2) {
		panic("Size mismatch!")
	}
}
