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

func (m *MatrixFloat64) Scale(scalar float64) *MatrixFloat64 {
	return WrapSlice(m.Rows, m.Columns, MapFunc(m.Data, func(f float64) float64 { return f * scalar }))
}

func (m *MatrixFloat64) MapFunc(f func(float64) float64) *MatrixFloat64 {
	return WrapSlice(m.Rows, m.Columns, MapFunc(m.Data, f))
}

func (m *MatrixFloat64) Dot(vector []float64) *MatrixFloat64 {
	return m.Product(WrapSlice(1, len(vector), vector))
}

func (m *MatrixFloat64) Product(m2 *MatrixFloat64) *MatrixFloat64 {
	if m2.Rows != m.Columns {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix(m.Rows, m2.Columns)

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m2.Columns; col++ {
			index := row*m2.Columns + col
			newMatrix.Data[index] = DotProduct(m.GetRow(row), m2.GetColumn(col))
		}
	}

	return newMatrix
}

func (m *MatrixFloat64) Transpose() *MatrixFloat64 {
	newMatrix := NewMatrix(m.Columns, m.Rows)

	for col := 0; col < m.Columns; col++ {
		// Using the power of slices referencing to the original memory array
		row := newMatrix.GetRow(col)
		copy(row, m.GetColumn(col))
	}

	return newMatrix
}

func (m *MatrixFloat64) Add(m2 *MatrixFloat64) *MatrixFloat64 {
	panicShapeMismatch(m, m2)

	return WrapSlice(m.Rows, m.Columns, SumVectors(m.Data, m2.Data))
}

func (m *MatrixFloat64) AddVectorPerRow(vector []float64) *MatrixFloat64 {
	if m.Columns != len(vector) {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix(m.Rows, m.Columns)

	for row := 0; row < m.Rows; row++ {
		// Using the power of slices referencing to the original memory array
		result := SumVectors(m.GetRow(row), vector)
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

func (m *MatrixFloat64) AddVectorPerColumn(vector []float64) *MatrixFloat64 {
	if m.Rows != len(vector) {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix(m.Rows, m.Columns)

	for row := 0; row < m.Rows; row++ {
		result := MapFunc(m.GetRow(row), func(f float64) float64 { return f + vector[row] })
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

func (m *MatrixFloat64) SoftMax() *MatrixFloat64 {
	newMatrix := NewMatrix(m.Rows, m.Columns)

	for row := 0; row < m.Rows; row++ {
		result := SoftMax(m.GetRow(row))
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
}

func (m *MatrixFloat64) DerivativeSoftMax() []MatrixFloat64 {
	// Jacobian matrices
	newMatrices := make([]MatrixFloat64, m.Rows)

	for row := 0; row < m.Rows; row++ {
		newMatrices[row] = *DerivativeSoftMax(m.GetRow(row))
	}

	return newMatrices
}

func (m *MatrixFloat64) CrossEntropyLossPerRow(targets *MatrixFloat64) []float64 {
	panicShapeMismatch(m, targets)

	newVector := make([]float64, m.Rows)

	for i := 0; i < m.Rows; i++ {
		newVector[i] = CrossEntropyLoss(m.GetRow(i), targets.GetRow(i))
	}

	return newVector
}

func (m *MatrixFloat64) DerivativeCrossEntropyLossPerRow(targets *MatrixFloat64) []float64 {
	panicShapeMismatch(m, targets)

	newVector := make([]float64, m.Rows)

	for i := 0; i < m.Rows; i++ {
		newVector[i] = DerivativeCrossEntropyLoss(m.GetRow(i), targets.GetRow(i))
	}

	return newVector
}

func (m *MatrixFloat64) DerivativeCrossEntropyLossSoftMaxPerRow(targetsOneHot []int) *MatrixFloat64 {
	if len(targetsOneHot) != m.Rows {
		panic("Size mismatch")
	}

	newMatrix := m.Copy()

	for row := 0; row < m.Rows; row++ {
		col := targetsOneHot[row]
		panicOutOfBound(m.Columns, col)

		newValue := newMatrix.GetAt(row, col) - 1
		newMatrix.Set(row, col, newValue)
	}

	return newMatrix
}

func (m *MatrixFloat64) ToOneHot() []int {
	onehot := make([]int, m.Rows)

	for row := 0; row < m.Rows; row++ {
		index, _ := Max(m.GetRow(row))

		onehot[row] = index
	}

	return onehot
}

func (m *MatrixFloat64) Accuracy(targets *MatrixFloat64) float64 {
	panicShapeMismatch(m, targets)

	var counter int = 0

	for i := 0; i < m.Rows; i++ {
		maxIndex, _ := Max(m.GetRow(i))
		maxTargetIndex, _ := Max(targets.GetRow(i))

		if maxIndex == maxTargetIndex {
			counter++
		}
	}

	return float64(counter) / float64(m.Rows)
}

func (m *MatrixFloat64) IsEqual(m2 *MatrixFloat64) bool {
	if m.Columns != m2.Columns || m.Rows != m2.Rows {
		return false
	}

	return AreEqual(m.Data, m2.Data)
}

func (m *MatrixFloat64) Copy() *MatrixFloat64 {
	newData := make([]float64, len(m.Data))
	copy(newData, m.Data)

	return &MatrixFloat64{
		Rows:    m.Rows,
		Columns: m.Columns,
		Data:    newData,
	}
}

func (m *MatrixFloat64) String() string {
	var sb strings.Builder

	// Assuming a precision of 11 digits, integer and comma, one comma and one space per element
	// and finally 200 character for extra buffer, but it still can grow larger, this is just to reduce
	// the times it needs to reallocate memory
	sb.Grow(15*m.Columns*m.Rows + 200)

	fmt.Fprintf(&sb, "Rows: %d \n", m.Rows)
	fmt.Fprintf(&sb, "Columns: %d \n", m.Columns)

	for row := 0; row < m.Rows; row++ {
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

func panicShapeMismatch(m1, m2 *MatrixFloat64) {
	if m1.Columns != m2.Columns || m1.Rows != m2.Rows || len(m1.Data) != len(m2.Data) {
		panic("Shape or data size mismatch")
	}
}
