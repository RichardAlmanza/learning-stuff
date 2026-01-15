package vector

import (
	"math"
)

type Real interface {
	~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Number interface {
	Real | Integer
}

// Vector is a generic type that represents a mathematical vector.
type Vector[T Number] interface {
	~[]T
}

func panicVectorSize[V Vector[T], T Number](vector1, vector2 V) {
	if len(vector1) != len(vector2) {
		panic("Size mismatch!")
	}
}

// AreEqual compares two vectors and returns true if they are equal.
func AreEqual[V Vector[T], T Number](vector1, vector2 V) bool {
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

// Reduce applies a function to each element of the vector and returns the accumulated result.
func Reduce[V Vector[T], T Number](vector V, initialValue T, f func(int, T) T) T {
	result := initialValue

	for i := 0; i < len(vector); i++ {
		result += f(i, vector[i])
	}

	return result
}

// Map applies a function to each element of the vector and returns a new vector with the results.
func MapFunc[V Vector[T], T Number](vector V, f func(int, T) T) V {
	newVector := make(V, len(vector))

	for i := 0; i < len(vector); i++ {
		newVector[i] = f(i, vector[i])
	}

	return newVector
}

// Map2Func applies a function that receives two vectors and returns a new vector with
// the results of applying the function to each pair of elements.
func Map2Func[V Vector[T], T Number](vector1, vector2 V, f func(int, T, T) T) V {
	panicVectorSize(vector1, vector2)

	newVector := make(V, len(vector1))

	for i := 0; i < len(vector1); i++ {
		newVector[i] = f(i, vector1[i], vector2[i])
	}

	return newVector
}

// Filter applies a function to each element of the vector
// and returns a new vector with the elements for which
// the function returns true.
func FilterFunc[V Vector[T], T Number](vector V, f func(int, T) bool) V {
	newVector := make(V, 0, len(vector)/2)

	for i := 0; i < len(vector); i++ {
		if f(i, vector[i]) {
			newVector = append(newVector, vector[i])
		}
	}

	return newVector
}

// AddVectors returns a new vector that is the element-wise sum of two vectors.
func AddVectors[V Vector[T], T Number](vector1, vector2 V) V {
	return Map2Func(vector1, vector2, func(_ int, n1, n2 T) T { return n1 + n2 })
}

// Sum returns the sum of all elements in a vector
func Sum[V Vector[T], T Number](vector V) T {
	return Reduce(vector, 0, func(_ int, n T) T { return n })
}

// Avg returns the average of all elements in a vector
func Avg[V Vector[T], T Number](vector V) float64 {
	return float64(Sum(vector)) / float64(len(vector))
}

// Min returns the minimum value in a vector and its index.
func Min[V Vector[T], T Number](vector V) (int, T) {
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

// Max returns the maximum value in a vector and its index.
func Max[V Vector[T], T Number](vector V) (int, T) {
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

// DotProduct returns the dot product of two vectors.
func DotProduct[V Vector[T], T Number](vector1, vector2 V) T {
	return Sum(Map2Func(vector1, vector2, func(i int, n1, n2 T) T { return n1 * n2 }))
}

// SoftMax returns the softmax of a vector.
func SoftMax[V Vector[T], T Real](vector V) V {
	_, max := Max(vector)
	newVector := MapFunc(vector, func(_ int, r T) T { return T(math.Pow(math.E, float64(r-max))) })

	sum := Sum(newVector)

	newVector = MapFunc(newVector, func(_ int, r T) T { return r / sum })

	return newVector
}

// CrossEntropyLoss returns the cross entropy loss of a vector based on the targets.
func CrossEntropyLoss[V Vector[T], T Number](vector, targets V) float64 {
	panicVectorSize(vector, targets)

	var sum float64 = 0

	for i := 0; i < len(vector); i++ {
		if targets[i] == 0 {
			continue
		}

		if vector[i] == 0 {
			return math.MaxFloat64
		}

		sum += float64(targets[i]) * math.Log(float64(vector[i]))
	}

	return -sum
}

// func DerivativeSoftMax(vector []float64) *MatrixFloat64 {
// 	diagVector := NewMatrixDiagonalFromSlice(vector)
// 	m1 := WrapSlice(len(vector), 1, vector)
// 	m1 = m1.Product(m1.Transpose()).Scale(-1)

// 	return diagVector.Add(m1)
// }

// func DerivativeCrossEntropyLoss(vector, targets []float64) float64 {
// 	return -Sum(Map2VectorFunc(vector, targets, func(f1, f2 float64) float64 { return f2 / f1 }))
// }
