package matrix

import (
	"fmt"
	"math/rand"
	"strings"
)

type MatrixFloat64 struct {
	Rows    int
	Columns int
	Data    []float64
}

func NewMatrix(row, col int) *MatrixFloat64 {
	return &MatrixFloat64{
		Rows:    row,
		Columns: col,
		Data:    make([]float64, row*col),
	}
}

func NewMatrixOneHot(row, col int, indexes []int) *MatrixFloat64 {
	if len(indexes) != row {
		panic("Size mismatch")
	}

	newMatrix := NewMatrix(row, col)

	for i := 0; i < row; i++ {
		panicOutOfBound(col, indexes[i])
		newMatrix.Set(i, indexes[i], 1)
	}

	return newMatrix
}

func NewMatrixRand(row, col int) *MatrixFloat64 {
	newMatrix := NewMatrix(row, col)
	newMatrix.Randomize()

	return newMatrix
}

func NewMatrixFromSlice(row, col int, slice []float64) *MatrixFloat64 {
	if row*col != len(slice) {
		panic("Size mismatch!")
	}

	newData := make([]float64, len(slice))

	copy(newData, slice)

	return &MatrixFloat64{
		Rows:    row,
		Columns: col,
		Data:    newData,
	}
}

func NewMatrixDiagonalFromSlice(slice []float64) *MatrixFloat64 {
	length := len(slice)
	newData := make([]float64, length*length)

	for i := 0; i < length; i++ {
		index := i*length + i
		newData[index] = slice[i]
	}

	return &MatrixFloat64{
		Rows:    length,
		Columns: length,
		Data:    newData,
	}
}

func WrapSlice(row, col int, slice []float64) *MatrixFloat64 {
	if row*col != len(slice) {
		panic("Size mismatch!")
	}

	return &MatrixFloat64{
		Rows:    row,
		Columns: col,
		Data:    slice,
	}
}

func (m *MatrixFloat64) Set(row, col int, value float64) {
	panicOutOfBound(m.Rows, row)
	panicOutOfBound(m.Columns, col)

	index := row * m.Columns
	index += col

	m.Data[index] = value
}

func (m *MatrixFloat64) GetAt(row, col int) float64 {
	panicOutOfBound(m.Rows, row)
	panicOutOfBound(m.Columns, col)

	index := row * m.Columns
	index += col

	return m.Data[index]
}

func (m *MatrixFloat64) GetRow(index int) []float64 {
	panicOutOfBound(m.Rows, index)

	start := index * m.Columns
	end := start + m.Columns

	return m.Data[start:end]
}

// GetColumn return a NEW array with the values of the queried column, ordered top to bottom.
// If you need to query the columns tons of times, is more efficient to transpose (Transpose) the matrix
// one time and then query the Rows (GetRow) of the transposed matrix.
func (m *MatrixFloat64) GetColumn(index int) []float64 {
	panicOutOfBound(m.Columns, index)

	sliceIndex := 0
	newSlice := make([]float64, m.Rows)

	for i := index; i < len(m.Data); i += m.Columns {
		newSlice[sliceIndex] = m.Data[i]
		sliceIndex++
	}

	return newSlice
}

func (m *MatrixFloat64) Randomize() {
	for i := 0; i < len(m.Data); i++ {
		m.Data[i] = rand.Float64() - 0.5
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

	fmt.Fprintf(&sb, "Columns: %d \n", m.Columns)
	fmt.Fprintf(&sb, "Rows: %d \n", m.Rows)

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

func panicShapeMismatch(m1, m2 *MatrixFloat64) {
	if m1.Columns != m2.Columns || m1.Rows != m2.Rows || len(m1.Data) != len(m2.Data) {
		panic("Shape or data size mismatch")
	}
}
