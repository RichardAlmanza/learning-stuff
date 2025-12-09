package matrix

import "math/rand"

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

func (m *MatrixFloat64) GetAt(indexCol, indexRow int) float64 {
	panicOutOfBound(m.Rows, indexRow)
	panicOutOfBound(m.Columns, indexCol)

	index := indexRow * m.Columns
	index += indexCol

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

func (m *MatrixFloat64) Add(m2 *MatrixFloat64) *MatrixFloat64 {
	if m.Columns != m2.Columns || m.Rows != m2.Rows {
		panic("Shape mismatch")
	}

	newMatrix := NewMatrix(m.Rows, m.Columns)

	for i := 0; i < len(m.Data); i++ {
		newMatrix.Data[i] = m.Data[i] + m2.Data[i]
	}

	return newMatrix
}

func panicOutOfBound(max, value int) {
	if value < 0 || value >= max {
		panic("Be serious, you're out of bounds!")
	}
}
