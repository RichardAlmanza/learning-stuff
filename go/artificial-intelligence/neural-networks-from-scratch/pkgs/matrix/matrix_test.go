package matrix_test

import (
	"fmt"
	// "math"
	"testing"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func TestMatrix_GetAt_Float64(t *testing.T) {
	testCases := []struct {
		name           string
		matrix         *matrix.Matrix[float64]
		position       []int
		expectedResult float64
	}{
		{
			name:           "Shape (1,1)",
			matrix:         matrix.NewMatrixFromSlice([]int{1, 1}, []float64{5}),
			position:       []int{0, 0},
			expectedResult: 5.0,
		},
		{
			name: "Shape (1,20)",
			matrix: func() *matrix.Matrix[float64] {
				nm := matrix.NewMatrix[float64]([]int{1, 20})
				nm.Data[13] = 6.2
				return nm
			}(),
			position:       []int{0, 13},
			expectedResult: 6.2,
		},
		{
			name: "Shape (20,1)",
			matrix: func() *matrix.Matrix[float64] {
				nm := matrix.NewMatrix[float64]([]int{20, 1})
				nm.Data[13] = 6.5
				return nm
			}(),
			position:       []int{13, 0},
			expectedResult: 6.5,
		},
		{
			name: "Shape (20,20)",
			matrix: func() *matrix.Matrix[float64] {
				nm := matrix.NewMatrix[float64]([]int{20, 20})
				nm.Data[33] = 2.2
				return nm
			}(),
			position:       []int{33 / 20, 33 % 20},
			expectedResult: 2.2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			number := tc.matrix.GetAt(tc.position)

			if number != tc.expectedResult {
				t.Errorf("\nfloat expected: %.3f | got: %.3f \n%v", tc.expectedResult, number, tc.matrix)
			}
		})
	}
}

func TestMatrix_Set_Int8(t *testing.T) {
	testCases := []struct {
		name           string
		shape          []int
		position       []int
		expectedResult int8
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			position:       []int{0, 0},
			expectedResult: 5,
		},
		{
			name:           "Shape (1,20)",
			shape:          []int{1, 20},
			position:       []int{0, 13},
			expectedResult: 6,
		},
		{
			name:           "Shape (20,1)",
			shape:          []int{20, 1},
			position:       []int{13, 0},
			expectedResult: 4,
		},
		{
			name:           "Shape (20,20)",
			shape:          []int{20, 20},
			position:       []int{4, 15},
			expectedResult: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm := matrix.NewMatrix[int8](tc.shape)
			nm.Set(tc.position, tc.expectedResult)
			number := nm.GetAt(tc.position)

			if number != tc.expectedResult {
				t.Errorf("\nfloat expected: %v | got: %v \n%v", tc.expectedResult, number, nm)
			}
		})
	}
}

func TestMatrix_GetRow_Float32(t *testing.T) {
	baseVector := make([]float32, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float32(i)
	}

	testCases := []struct {
		name           string
		shape          vector.Shape
		row            int
		expectedResult []float32
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			row:            0,
			expectedResult: []float32{0},
		},
		{
			name:  "Shape (1,20)",
			shape: []int{1, 20},
			row:   0,
			expectedResult: func() []float32 {
				v := make([]float32, 20)

				for i := 0; i < 20; i++ {
					v[i] = float32(i)
				}

				return v
			}(),
		},
		{
			name:           "Shape (20,1)",
			shape:          []int{20, 1},
			row:            5,
			expectedResult: []float32{5},
		},
		{
			name:  "Shape (20,20)",
			shape: []int{20, 20},
			row:   12,
			expectedResult: func() []float32 {
				v := make([]float32, 20)
				start := 20 * 12
				end := start + 20

				var index int = 0
				for i := start; i < end; i++ {
					v[index] = float32(i)
					index++
				}

				return v
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape.TotalSize()
			nm := matrix.NewMatrixFromSlice(tc.shape, baseVector[:size])
			row := nm.GetRow(tc.row)

			if !vector.AreEqual(row, tc.expectedResult) {
				t.Errorf("\nexpected vector: %v | got: %v \n%v", tc.expectedResult, row, nm)
			}
		})
	}
}

func TestMatrix_GetColumn_Int16(t *testing.T) {
	baseVector := make([]int16, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = int16(i)
	}

	testCases := []struct {
		name           string
		shape          vector.Shape
		col            int
		expectedResult []int16
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			col:            0,
			expectedResult: []int16{0},
		},
		{
			name:           "Shape (1,20)",
			shape:          []int{1, 20},
			col:            6,
			expectedResult: []int16{6},
		},
		{
			name:  "Shape (20,1)",
			shape: []int{20, 1},
			col:   0,
			expectedResult: func() []int16 {
				v := make([]int16, 20)

				for i := 0; i < 20; i++ {
					v[i] = int16(i)
				}

				return v
			}(),
		},
		{
			name:  "Shape (20,20)",
			shape: []int{20, 20},
			col:   12,
			expectedResult: func() []int16 {
				v := make([]int16, 20)

				value := int16(12)
				for i := 0; i < 20; i++ {
					v[i] = value
					value += 20
				}

				return v
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape.TotalSize()
			nm := matrix.NewMatrixFromSlice(tc.shape, baseVector[:size])
			col := nm.GetColumn(tc.col)

			if !vector.AreEqual(col, tc.expectedResult) {
				t.Errorf("\nexpected vector: %v | got: %v \n%v", tc.expectedResult, col, nm)
			}
		})
	}
}

func TestMatrix_Randomize_Float32(t *testing.T) {
	sizeX := 2000
	sizeY := 2000
	nm := matrix.NewMatrix[float32]([]int{sizeY, sizeX})
	nm.Randomize()

	probabilityDistribution := vector.SoftMax(nm.Data)
	_, min := vector.Min(probabilityDistribution)
	_, max := vector.Max(probabilityDistribution)

	maxDelta := float32(5e-7)
	actualDelta := max - min

	testName := fmt.Sprintf("Randomize keeping Normal distribution Shape (%d,%d) MaxDelta %v", sizeY, sizeX, maxDelta)

	t.Run(testName, func(t *testing.T) {

		if actualDelta > maxDelta {
			t.Errorf("\nexpected maximum difference: %v | got: %v", maxDelta, actualDelta)
		}
	})
}

func TestMatrix_Scale_Int32(t *testing.T) {
	scalar := -2.0
	var shape vector.Shape = []int{10, 10}

	baseVector := make([]int32, shape.TotalSize())

	for i := 0; i < len(baseVector); i++ {
		baseVector[i] = 10
	}

	nm := matrix.WrapSlice(shape, baseVector).Scale(scalar)

	expectedMatrix := matrix.NewMatrix[int32](shape)
	for i := 0; i < len(expectedMatrix.Data); i++ {
		expectedMatrix.Data[i] = -20
	}

	testName := fmt.Sprintf("Scale matrix %v Shape %v", scalar, shape)

	t.Run(testName, func(t *testing.T) {
		if !expectedMatrix.IsEqual(nm) {
			t.Errorf("\nexpected : \n%v got: \n%v", expectedMatrix, nm)
		}
	})
}

func TestMatrix_MapFunction_Int(t *testing.T) {
	f := func(_, v int) int { return 10 }
	var shape vector.Shape = []int{10, 10}

	nm := matrix.NewMatrix[int](shape).MapFunc(f)

	expectedMatrix := matrix.NewMatrix[int](shape)
	for i := 0; i < len(expectedMatrix.Data); i++ {
		expectedMatrix.Data[i] = 10
	}

	testName := fmt.Sprintf("Map matrix Shape %v", shape)

	t.Run(testName, func(t *testing.T) {
		if !expectedMatrix.IsEqual(nm) {
			t.Errorf("\nexpected : \n%v got: \n%v", expectedMatrix, nm)
		}
	})
}

func TestMatrix_Transpose_Float64(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          vector.Shape
		expectedResult *matrix.Matrix[float64]
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			expectedResult: matrix.WrapSlice([]int{1, 1}, []float64{0}),
		},
		{
			name:           "Shape (1,10) -> (10,1)",
			shape:          []int{1, 10},
			expectedResult: matrix.WrapSlice([]int{10, 1}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name:           "Shape (10,1) -> (1,10)",
			shape:          []int{10, 1},
			expectedResult: matrix.WrapSlice([]int{1, 10}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name:  "Shape (5,5)",
			shape: []int{5, 5},
			expectedResult: matrix.WrapSlice(
				[]int{5, 5},
				[]float64{
					0, 5, 10, 15, 20,
					1, 6, 11, 16, 21,
					2, 7, 12, 17, 22,
					3, 8, 13, 18, 23,
					4, 9, 14, 19, 24,
				},
			),
		},
		{
			name:  "Shape (16,25) -> (25,16)",
			shape: []int{16, 25},
			expectedResult: matrix.WrapSlice(
				[]int{25, 16},
				func() []float64 {
					nv := make([]float64, 400)

					var index int = 0
					for row := 0; row < 25; row++ {
						for col := row; col < 400; col += 25 {
							nv[index] = float64(col)
							index++
						}
					}

					return nv
				}(),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape.TotalSize()
			nm := matrix.WrapSlice(tc.shape, baseVector[:size]).Transpose()

			if !tc.expectedResult.IsEqual(nm) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nm)
			}
		})
	}
}

func TestMatrix_Add_Int32(t *testing.T) {
	baseVector := make([]int32, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = int32(i)
	}

	testCases := []struct {
		name           string
		shape          vector.Shape
		times          int
		expectedResult *matrix.Matrix[int32]
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			times:          20,
			expectedResult: matrix.NewMatrixFromSlice([]int{1, 1}, []int32{0}),
		},
		{
			name:  "Shape (1,10)",
			shape: []int{1, 10},
			times: 4,
			expectedResult: matrix.NewMatrixFromSlice(
				[]int{1, 10},
				vector.MapFunc(make([]int32, 10), func(i int, _ int32) int32 { return int32(i) * 4 }),
			),
		},
		{
			name:  "Shape (10,1)",
			shape: []int{10, 1},
			times: 15,
			expectedResult: matrix.NewMatrixFromSlice(
				[]int{10, 1},
				vector.MapFunc(make([]int32, 10), func(i int, _ int32) int32 { return int32(i) * 15 }),
			),
		},
		{
			name:  "Shape (5,5)",
			shape: []int{5, 5},
			times: 2,
			expectedResult: matrix.NewMatrixFromSlice(
				[]int{5, 5},
				vector.MapFunc(make([]int32, 25), func(i int, _ int32) int32 { return int32(i) * 2 }),
			),
		},
		{
			name:  "Shape (16,25)",
			shape: []int{16, 25},
			times: 10,
			expectedResult: matrix.NewMatrixFromSlice(
				[]int{16, 25},
				vector.MapFunc(make([]int32, 400), func(i int, _ int32) int32 { return int32(i) * 10 }),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape.TotalSize()
			nm := matrix.WrapSlice(tc.shape, baseVector[:size])

			nmSum := matrix.NewMatrix[int32](tc.shape)
			for i := 0; i < tc.times; i++ {
				nmSum = nmSum.Add(nm)
			}

			if !tc.expectedResult.IsEqual(nmSum) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nmSum)
			}
		})
	}
}

func TestMatrix_AddVectorPerRow_Int16(t *testing.T) {
	baseVector := make([]int16, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = int16(i)
	}

	testCases := []struct {
		name           string
		shape          vector.Shape
		vector         []int16
		expectedResult *matrix.Matrix[int16]
	}{
		{
			name:           "Shape (1,1)",
			shape:          []int{1, 1},
			vector:         []int16{5},
			expectedResult: matrix.WrapSlice([]int{1, 1}, []int16{5}),
		},
		{
			name:           "Shape (1,10)",
			shape:          []int{1, 10},
			vector:         []int16{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expectedResult: matrix.WrapSlice([]int{1, 10}, []int16{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}),
		},
		{
			name:           "Shape (10,1)",
			shape:          []int{10, 1},
			vector:         []int16{10},
			expectedResult: matrix.WrapSlice([]int{10, 1}, []int16{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}),
		},
		{
			name:   "Shape (5,5)",
			shape:  []int{5, 5},
			vector: []int16{-10, 10, -10, 10, -10},
			expectedResult: matrix.WrapSlice(
				[]int{5, 5},
				func() []int16 {
					nv := make([]int16, 25)
					for row := 0; row < 5; row++ {
						value := -10.0
						for col := 0; col < 5; col++ {
							index := 5*row + col
							nv[index] = int16(index) + int16(value)
							value *= -1
						}
					}
					return nv
				}(),
			),
		},
		{
			name:  "Shape (16,25)",
			shape: []int{16, 25},
			vector: func() []int16 {
				nv := make([]int16, 25)
				var value int16 = -10
				for i := 0; i < len(nv); i++ {
					nv[i] = value
					value *= -1
				}
				return nv
			}(),
			expectedResult: matrix.WrapSlice(
				[]int{16, 25},
				func() []int16 {
					nv := make([]int16, 400)
					var value int16
					for row := 0; row < 16; row++ {
						value = -10
						for col := 0; col < 25; col++ {
							index := 25*row + col
							nv[index] = int16(index) + value
							value *= -1
						}
					}
					return nv
				}(),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape.TotalSize()
			nm := matrix.WrapSlice(tc.shape, baseVector[:size]).AddVectorPerRow(tc.vector)

			if !tc.expectedResult.IsEqual(nm) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nm)
			}
		})
	}
}

func TestMatrix_AddVectorPerColumn(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          [2]int
		vector         []float64
		expectedResult *matrix.MatrixFloat64
	}{
		{
			name:   "Shape (1,1)",
			shape:  [2]int{1, 1},
			vector: []float64{5},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 1,
				Data:    []float64{5},
			},
		},
		{
			name:   "Shape (1,10)",
			shape:  [2]int{1, 10},
			vector: []float64{10},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 10,
				Data:    []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
			},
		},
		{
			name:   "Shape (10,1)",
			shape:  [2]int{10, 1},
			vector: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    10,
				Columns: 1,
				Data:    []float64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			},
		},
		{
			name:   "Shape (5,5)",
			shape:  [2]int{5, 5},
			vector: []float64{-10, 10, -10, 10, 10},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    5,
				Columns: 5,
				Data: func() []float64 {
					nv := make([]float64, 25)
					value := -10.0
					for row := 0; row < 5; row++ {
						for col := 0; col < 5; col++ {
							index := 5*row + col
							nv[index] = float64(index) + value
						}
						value *= -1
						if row == 3 {
							value *= -1
						}
					}
					return nv
				}(),
			},
		},
		{
			name:  "Shape (16,25)",
			shape: [2]int{16, 25},
			vector: func() []float64 {
				nv := make([]float64, 16)
				value := -10.0
				for i := 0; i < len(nv); i++ {
					nv[i] = value
					value *= -1
				}
				return nv
			}(),
			expectedResult: &matrix.MatrixFloat64{
				Rows:    16,
				Columns: 25,
				Data: func() []float64 {
					nv := make([]float64, 400)
					value := -10.0
					for row := 0; row < 16; row++ {
						for col := 0; col < 25; col++ {
							index := 25*row + col
							nv[index] = float64(index) + value
						}
						value *= -1
					}
					return nv
				}(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.WrapSlice(tc.shape[0], tc.shape[1], baseVector[:size]).AddVectorPerColumn(tc.vector)

			if !tc.expectedResult.IsEqual(nm) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nm)
			}
		})
	}
}

func TestMatrix_Product(t *testing.T) {
	testCases := []struct {
		name           string
		matrixA        *matrix.MatrixFloat64
		matrixB        *matrix.MatrixFloat64
		expectedResult *matrix.MatrixFloat64
	}{
		{
			name: "Shapes (1,3) * (3, 1)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 3,
				Data: []float64{
					1, 2, 3,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 1,
				Data: []float64{
					4,
					5,
					6,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 1,
				Data:    []float64{32},
			},
		},
		{
			name: "Shapes (3,1) * (1, 3)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 1,
				Data: []float64{
					1,
					2,
					3,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 3,
				Data: []float64{
					4, 5, 6,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 3,
				Data: []float64{
					4, 5, 6,
					8, 10, 12,
					12, 15, 18,
				},
			},
		},
		{
			name: "Shapes (2,3) * (3,2)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 3,
				Data: []float64{
					1, 2, 3,
					4, 5, 6,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 2,
				Data: []float64{
					7, 8,
					9, 10,
					11, 12,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					58, 64,
					139, 154,
				},
			},
		},
		{
			name: "Shapes (2,2) * (2,2)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					1, 2,
					3, 4,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					2, 0,
					1, 2,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					4, 4,
					10, 8,
				},
			},
		},
		{
			name: "Shapes (2,2) * (2,2)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					2, 0,
					1, 2,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					1, 2,
					3, 4,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    2,
				Columns: 2,
				Data: []float64{
					2, 4,
					7, 10,
				},
			},
		},
		{
			name: "Shapes (1,3) * (3,4)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 3,
				Data: []float64{
					3, 4, 2,
				},
			},
			matrixB: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 4,
				Data: []float64{
					13, 9, 7, 15,
					8, 7, 4, 6,
					6, 4, 0, 3,
				},
			},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 4,
				Data: []float64{
					83, 63, 37, 75,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm := tc.matrixA.Product(tc.matrixB)

			if !tc.expectedResult.IsEqual(nm) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nm)
			}
		})
	}
}

func TestMatrix_CrossEntropyLossPerRow(t *testing.T) {
	testCases := []struct {
		name           string
		matrixA        *matrix.MatrixFloat64
		matrixTargets  *matrix.MatrixFloat64
		expectedResult []float64
	}{
		{
			name: "One-hot Shape (3,3)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 3,
				Data: []float64{
					0.7, 0.1, 0.2,
					0.1, 0.5, 0.4,
					0.02, 0.9, 0.08,
				},
			},
			matrixTargets:  matrix.NewMatrixOneHot(3, 3, []int{0, 1, 1}),
			expectedResult: []float64{-math.Log(0.7), -math.Log(0.5), -math.Log(0.9)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nv := tc.matrixA.CrossEntropyLossPerRow(tc.matrixTargets)

			if !matrix.AreEqual(nv, tc.expectedResult) {
				t.Errorf("\nexpected vector: %v \ngot: \n%v", tc.expectedResult, nv)
			}
		})
	}
}

func TestMatrix_Accuracy(t *testing.T) {
	testCases := []struct {
		name           string
		matrixA        *matrix.MatrixFloat64
		matrixTargets  *matrix.MatrixFloat64
		expectedResult float64
	}{
		{
			name: "One-hot Shape (3,3)",
			matrixA: &matrix.MatrixFloat64{
				Rows:    3,
				Columns: 3,
				Data: []float64{
					0.7, 0.2, 0.1,
					0.5, 0.1, 0.4,
					0.02, 0.9, 0.08,
				},
			},
			matrixTargets:  matrix.NewMatrixOneHot(3, 3, []int{0, 1, 1}),
			expectedResult: 2.0 / 3.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			acc := tc.matrixA.Accuracy(tc.matrixTargets)

			if acc != tc.expectedResult {
				t.Errorf("\nexpected value: %v \ngot: \n%v", tc.expectedResult, acc)
			}
		})
	}
}

func BenchmarkMatrix_GetRow(b *testing.B) {
	sizeX := 1000
	sizeY := 1000

	nm := matrix.NewMatrixRand(sizeY, sizeX)

	index := 0
	for b.Loop() {
		nm.GetRow(index)
		index = (index + 1) % sizeY
	}
}

func BenchmarkMatrix_GetColumn(b *testing.B) {
	sizeX := 1000
	sizeY := 1000

	nm := matrix.NewMatrixRand(sizeY, sizeX)

	index := 0
	for b.Loop() {
		nm.GetColumn(index)
		index = (index + 1) % sizeX
	}
}

func BenchmarkMatrix_Scale(b *testing.B) {
	sizeX := 20
	sizeY := 20

	nm := matrix.NewMatrixRand(sizeY, sizeX)

	for b.Loop() {
		nm.Scale(-1)
	}
}

func BenchmarkMatrix_Add(b *testing.B) {
	sizeX := 1000
	sizeY := 1000

	nm1 := matrix.NewMatrixRand(sizeY, sizeX)
	nm2 := matrix.NewMatrixRand(sizeY, sizeX)

	for b.Loop() {
		nm1.Add(nm2)
	}
}

func BenchmarkVector_Filter(b *testing.B) {
	baseVector := make([]float64, 40000)

	for i := 0; i < 40000; i++ {
		baseVector[i] = float64(i)
	}

	for b.Loop() {
		matrix.FilterFunction(baseVector, func(v float64) bool {
			return v < 20000
		})
	}
}
