package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

var (
	matrix *Matrix[int8] = NewMatrix[int8]([]int{1, 1})
	_      vector.Dense  = matrix
)

type Matrix[T vector.Number] struct {
	shape       vector.Shape
	elementType vector.NumberType
	Data        []T
}

func (m *Matrix[T]) Type() vector.NumberType {
	if m.elementType == 0 {
		m.elementType = vector.AssertNumber(m.Data[0])
	}

	return m.elementType
}

func (m *Matrix[T]) Shape() vector.Shape {
	return m.shape
}

func NewMatrix[T vector.Number](shape vector.Shape) *Matrix[T] {
	size := shape.TotalSize()
	if size <= 0 || shape.Dims() != 2 {
		panic("Invalid shape!")
	}

	data := make([]T, size)

	return &Matrix[T]{
		shape: shape.Copy(),
		Data:  data,
	}
}

func NewMatrixOneHot[T vector.Number](shape, indexes []int) *Matrix[T] {
	if len(indexes) != shape[0] {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix[T](shape)

	for i := 0; i < shape[0]; i++ {
		panicOutOfBound(shape[1], indexes[i])
		coordinates := []int{i, indexes[i]}

		newMatrix.Set(coordinates, 1)
	}

	return newMatrix
}

func NewMatrixRand[T vector.Number](shape vector.Shape) *Matrix[T] {
	newMatrix := NewMatrix[T](shape)
	newMatrix.Randomize()

	return newMatrix
}

func NewMatrixFromSlice[T vector.Number](shape vector.Shape, slice []T) *Matrix[T] {
	panicDimensionsShape(shape)

	if shape.TotalSize() != len(slice) {
		panic("Size mismatch!")
	}

	return &Matrix[T]{
		shape: shape,
		Data:  vector.NewCopy(slice),
	}
}

func NewMatrixDiagonalFromSlice[T vector.Number](slice []T) *Matrix[T] {
	length := len(slice)
	newMatrix := NewMatrix[T]([]int{length, length})

	for i := 0; i < length; i++ {
		index := i*length + i
		newMatrix.Data[index] = slice[i]
	}

	return newMatrix
}

func WrapSlice[T vector.Number](shape vector.Shape, slice []T) *Matrix[T] {
	panicDimensionsShape(shape)

	if shape.TotalSize() != len(slice) {
		panic("Size mismatch!")
	}

	return &Matrix[T]{
		shape: shape,
		Data:  slice,
	}
}

func (m *Matrix[T]) Set(coordinates vector.Shape, value T) {
	panicDimensionsShape(coordinates)

	row := coordinates[0]
	col := coordinates[1]

	panicOutOfBound(m.shape[0], row)
	panicOutOfBound(m.shape[1], col)

	index := row * m.shape[1]
	index += col

	m.Data[index] = value
}

func (m *Matrix[T]) GetAt(coordinates vector.Shape) T {
	panicDimensionsShape(coordinates)

	row := coordinates[0]
	col := coordinates[1]

	panicOutOfBound(m.shape[0], row)
	panicOutOfBound(m.shape[1], col)

	index := row * m.shape[1]
	index += col

	return m.Data[index]
}

func (m *Matrix[T]) GetRow(index int) []T {
	panicOutOfBound(m.shape[0], index)

	start := index * m.shape[1]
	end := start + m.shape[1]

	return m.Data[start:end]
}

// GetColumn return a NEW array with the values of the queried column, ordered top to bottom.
// If you need to query the columns tons of times, is more efficient to transpose (Transpose) the matrix
// one time and then query the Rows (GetRow) of the transposed matrix.
func (m *Matrix[T]) GetColumn(index int) []T {
	panicOutOfBound(m.shape[1], index)

	sliceIndex := 0
	newSlice := make([]T, m.shape[0])

	for i := index; i < len(m.Data); i += m.shape[1] {
		newSlice[sliceIndex] = m.Data[i]
		sliceIndex++
	}

	return newSlice
}

func (m *Matrix[T]) Randomize() {
	var f func(float64) T

	if m.Type()&vector.IntegerNumber != 0 {
		f = func(f float64) T { return T(math.Round(100 * f)) }
	} else {
		f = func(f float64) T { return T(f) }
	}

	for i := 0; i < len(m.Data); i++ {
		random := rand.Float64() - 0.5
		m.Data[i] = f(random)
	}
}

func (m *Matrix[T]) Scale(scalar float64) *Matrix[T] {
	return WrapSlice(m.shape, vector.MapFunc(m.Data, func(_ int, n T) T { return T(float64(n) * scalar) }))
}

func (m *Matrix[T]) MapFunc(f func(int, T) T) *Matrix[T] {
	return WrapSlice(m.shape, vector.MapFunc(m.Data, f))
}

func (m *Matrix[T]) Dot(vector []T) *Matrix[T] {
	return m.Product(WrapSlice([]int{1, len(vector)}, vector))
}

func (m *Matrix[T]) Product(m2 *Matrix[T]) *Matrix[T] {
	if m2.shape[0] != m.shape[1] {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix[T]([]int{m.shape[0], m2.shape[1]})

	for row := 0; row < m.shape[0]; row++ {
		for col := 0; col < m2.shape[1]; col++ {
			index := row*m2.shape[1] + col
			newMatrix.Data[index] = vector.DotProduct(m.GetRow(row), m2.GetColumn(col))
		}
	}

	return newMatrix
}

func (m *Matrix[T]) Transpose() *Matrix[T] {
	newShape := vector.Shape{m.shape[1], m.shape[0]}
	newMatrix := NewMatrix[T](newShape)

	for col := 0; col < m.shape[1]; col++ {
		// Using the power of slices referencing to the original memory array
		row := newMatrix.GetRow(col)
		copy(row, m.GetColumn(col))
	}

	return newMatrix
}

func (m *Matrix[T]) Add(m2 *Matrix[T]) *Matrix[T] {
	panicShapeMismatch(m, m2)

	return WrapSlice(m.shape, vector.AddVectors(m.Data, m2.Data))
}

func (m *Matrix[T]) AddVectorPerRow(v []T) *Matrix[T] {
	if m.shape[1] != len(v) {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix[T](m.shape)

	for row := 0; row < m.shape[0]; row++ {
		// Using the power of slices referencing to the original memory array
		result := vector.AddVectors(m.GetRow(row), v)
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

func (m *Matrix[T]) AddVectorPerColumn(v []T) *Matrix[T] {
	if m.shape[0] != len(v) {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix[T](m.shape)

	for row := 0; row < m.shape[0]; row++ {
		result := vector.MapFunc(m.GetRow(row), func(_ int, n T) T { return n + v[row] })
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

func SoftMax[T vector.Real](m *Matrix[T]) *Matrix[T] {
	newMatrix := NewMatrix[T](m.shape)

	for row := 0; row < m.shape[0]; row++ {
		result := vector.SoftMax(m.GetRow(row))
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

// func DerivativeSoftMax[T vector.Real](m *Matrix[T]) *Matrix[T] {
// 	// Jacobian matrices
// 	newMatrices := make([]Matrix[T], m.shape[0])

// 	for row := 0; row < m.shape[0]; row++ {
// 		newMatrices[row] = *vector.DerivativeSoftMax(m.GetRow(row))
// 	}

// 	return &newMatrices[0]
// }

func CrossEntropyLossPerRow[T vector.Real](m, targets *Matrix[T]) []float64 {
	panicShapeMismatch(m, targets)

	newVector := make([]float64, m.shape[0])

	for i := 0; i < m.shape[0]; i++ {
		newVector[i] = vector.CrossEntropyLoss(m.GetRow(i), targets.GetRow(i))
	}

	return newVector
}

// func DerivativeCrossEntropyLossPerRow[T vector.Number](m, targets *Matrix[T]) []float64 {
// 	panicShapeMismatch(m, targets)

// 	newVector := make([]float64, m.shape[0])

// 	for i := 0; i < m.shape[0]; i++ {
// 		newVector[i] = vector.DerivativeCrossEntropyLoss(m.GetRow(i), targets.GetRow(i))
// 	}

// 	return newVector
// }

func (m *Matrix[T]) DerivativeCrossEntropyLossSoftMaxPerRow(targetsOneHot []int) *Matrix[T] {
	if len(targetsOneHot) != m.shape[0] {
		panic("Size mismatch")
	}

	newMatrix := m.Copy()

	for row := 0; row < m.shape[0]; row++ {
		col := targetsOneHot[row]
		panicOutOfBound(m.shape[1], col)

		newValue := newMatrix.GetAt([]int{row, col}) - 1
		newMatrix.Set([]int{row, col}, newValue)
	}

	return newMatrix
}

func (m *Matrix[T]) ToOneHot() []int {
	onehot := make([]int, m.shape[0])

	for row := 0; row < m.shape[0]; row++ {
		index, _ := vector.Max(m.GetRow(row))

		onehot[row] = index
	}

	return onehot
}

func (m *Matrix[T]) Accuracy(targets *Matrix[T]) float64 {
	panicShapeMismatch(m, targets)

	var counter int = 0

	for i := 0; i < m.shape[0]; i++ {
		maxIndex, _ := vector.Max(m.GetRow(i))
		maxTargetIndex, _ := vector.Max(targets.GetRow(i))

		if maxIndex == maxTargetIndex {
			counter++
		}
	}

	return float64(counter) / float64(m.shape[0])
}

func (m *Matrix[T]) IsEqual(m2 *Matrix[T]) bool {
	if !vector.AreEqual(m.shape, m2.shape) {
		return false
	}

	return vector.AreEqual(m.Data, m2.Data)
}

func (m *Matrix[T]) Copy() *Matrix[T] {
	return &Matrix[T]{
		shape: m.shape.Copy(),
		Data:  vector.NewCopy(m.Data),
	}
}

func (m *Matrix[T]) String() string {
	var sb strings.Builder

	// Assuming a precision of 11 digits, integer and comma, one comma and one space per element
	// and finally 200 character for extra buffer, but it still can grow larger, this is just to reduce
	// the times it needs to reallocate memory
	sb.Grow(15*m.shape.TotalSize() + 200)

	fmt.Fprintf(&sb, "Rows: %d \n", m.shape[0])
	fmt.Fprintf(&sb, "Columns: %d \n", m.shape[1])

	for row := 0; row < m.shape[0]; row++ {
		fmt.Fprintf(&sb, "%d ", row)
		fmt.Fprint(&sb, m.GetRow(row))
		sb.WriteRune('\n')
	}

	return sb.String()
}

func panicOutOfBound(max, value int) {
	if value < 0 || value >= max {
		panic("Be serious, you're out of bounds!")
	}
}

func panicDimensionsShape(shape vector.Shape) {
	if shape.Dims() != 2 {
		panic("Invalid shape! Expected 2 dimensions.")
	}
}

func panicShapeMismatch[T vector.Number](m1, m2 *Matrix[T]) {
	if !vector.AreEqual(m1.shape, m2.shape) || len(m1.Data) != len(m2.Data) {
		panic("Shape or data size mismatch")
	}
}
