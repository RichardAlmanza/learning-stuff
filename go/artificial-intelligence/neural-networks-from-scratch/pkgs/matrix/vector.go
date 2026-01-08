package matrix

import "math"

func AreEqual(vector1, vector2 []float64) bool {
	if len(vector1) != len(vector2) {
		return false
	}

	for i := 0; i < len(vector1); i++ {
		if vector1[i] != vector2[i] {
			return false
		}
	}

	return true
}

func SumVectors(vector1, vector2 []float64) []float64 {
	panicVectorSize(vector1, vector2)

	newVector := make([]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		newVector[i] = vector1[i] + vector2[i]
	}

	return newVector
}

func Sum(vector []float64) float64 {
	var result float64 = 0

	for i := 0; i < len(vector); i++ {
		result += vector[i]
	}

	return result
}

func Avg(vector []float64) float64 {
	return Sum(vector) / float64(len(vector))
}

func Min(vector []float64)  (int, float64) {
	result := vector[0]
	index := 0

	for i := 1; i < len(vector); i++ {
		if result > vector[i] {
			result = vector[i]
			index = i
		}
	}

	return index, result
}

func Max(vector []float64) (int, float64) {
	result := vector[0]
	index := 0

	for i := 1; i < len(vector); i++ {
		if result < vector[i] {
			result = vector[i]
			index = i
		}
	}

	return index, result
}

func MapFunc(vector []float64, f func(float64) float64) []float64 {
	newVector := make([]float64, len(vector))

	for i := 0; i < len(vector); i++ {
		newVector[i] = f(vector[i])
	}

	return newVector
}

func MapFuncIndex(vector []float64, f func(int, float64) float64) []float64 {
	newVector := make([]float64, len(vector))

	for i := 0; i < len(vector); i++ {
		newVector[i] = f(i, vector[i])
	}

	return newVector
}

func FilterFunction(vector []float64, f func(float64) bool) []float64 {
	newVector := make([]float64, 0, len(vector)/2)

	for _, value := range vector {
		if f(value) {
			newVector = append(newVector, value)
		}
	}

	return newVector
}

func SoftMax(vector []float64) []float64 {
	_, max := Max(vector)
	newVector := MapFunc(vector, func(f float64) float64 { return math.Pow(math.E, f-max) })

	sum := Sum(newVector)

	newVector = MapFunc(newVector, func(f float64) float64 { return f / sum })

	return newVector
}

func CrossEntropyLoss(vector, targets []float64) float64 {
	panicVectorSize(vector, targets)

	var sum float64 = 0

	for i := 0; i < len(vector); i++ {
		if targets[i] == 0 {
			continue
		}

		if vector[i] == 0 {
			return math.MaxFloat64
		}

		sum += targets[i] * math.Log(vector[i])
	}

	return -sum
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
