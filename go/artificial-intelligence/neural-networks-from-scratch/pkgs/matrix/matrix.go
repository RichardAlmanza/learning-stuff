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

func (m *MatrixFloat64) Scale(scalar float64) {
	for i := 0; i < len(m.Data); i++ {
		m.Data[i] *= scalar
	}
}

func (m *MatrixFloat64) Dot(vector []float64) []float64 {
	if len(vector) != m.Columns {
		panic("Length mismatch")
	}

	result := make([]float64, m.Rows)

	for row := 0; row < m.Rows; row++ {
		result[row] = DotProduct(vector, m.GetRow(row))
	}

	return result
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
	if m.Columns != m2.Columns || m.Rows != m2.Rows {
		panic("Shape mismatch")
	}

	newMatrix := NewMatrix(m.Rows, m.Columns)

	for row := 0; row < m.Rows; row++ {
		// Using the power of slices referencing to the original memory array
		result := SumVectors(m.GetRow(row), m2.GetRow(row))
		copy(newMatrix.GetRow(row), result)
	}

	return newMatrix
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
