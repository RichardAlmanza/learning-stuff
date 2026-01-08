package matrix_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

func TestMatrix_GetAt(t *testing.T) {
	testCases := []struct {
		name           string
		matrix         *matrix.MatrixFloat64
		position       [2]int
		expectedResult float64
	}{
		{
			name:           "Shape (1,1)",
			matrix:         matrix.NewMatrixFromSlice(1, 1, []float64{5}),
			position:       [2]int{0, 0},
			expectedResult: 5.0,
		},
		{
			name: "Shape (1,20)",
			matrix: func() *matrix.MatrixFloat64 {
				nm := matrix.NewMatrix(1, 20)
				nm.Data[13] = 6.2
				return nm
			}(),
			position:       [2]int{0, 13},
			expectedResult: 6.2,
		},
		{
			name: "Shape (20,1)",
			matrix: func() *matrix.MatrixFloat64 {
				nm := matrix.NewMatrix(20, 1)
				nm.Data[13] = 6.5
				return nm
			}(),
			position:       [2]int{13, 0},
			expectedResult: 6.5,
		},
		{
			name: "Shape (20,20)",
			matrix: func() *matrix.MatrixFloat64 {
				nm := matrix.NewMatrix(20, 20)
				nm.Data[33] = 2.2
				return nm
			}(),
			position:       [2]int{33 / 20, 33 % 20},
			expectedResult: 2.2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			number := tc.matrix.GetAt(tc.position[0], tc.position[1])

			if number != tc.expectedResult {
				t.Errorf("\nfloat expected: %.3f | got: %.3f \n%v", tc.expectedResult, number, tc.matrix)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	testCases := []struct {
		name           string
		shape          [2]int
		position       [2]int
		expectedResult float64
	}{
		{
			name:           "Shape (1,1)",
			shape:          [2]int{1, 1},
			position:       [2]int{0, 0},
			expectedResult: 5.0,
		},
		{
			name:           "Shape (1,20)",
			shape:          [2]int{1, 20},
			position:       [2]int{0, 13},
			expectedResult: 6.2,
		},
		{
			name:           "Shape (20,1)",
			shape:          [2]int{20, 1},
			position:       [2]int{13, 0},
			expectedResult: 6.5,
		},
		{
			name:           "Shape (20,20)",
			shape:          [2]int{20, 20},
			position:       [2]int{4, 15},
			expectedResult: 2.2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm := matrix.NewMatrix(tc.shape[0], tc.shape[1])
			nm.Set(tc.position[0], tc.position[1], tc.expectedResult)
			number := nm.GetAt(tc.position[0], tc.position[1])

			if number != tc.expectedResult {
				t.Errorf("\nfloat expected: %.3f | got: %.3f \n%v", tc.expectedResult, number, nm)
			}
		})
	}
}

func TestMatrix_GetRow(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          [2]int
		row            int
		expectedResult []float64
	}{
		{
			name:           "Shape (1,1)",
			shape:          [2]int{1, 1},
			row:            0,
			expectedResult: []float64{0},
		},
		{
			name:  "Shape (1,20)",
			shape: [2]int{1, 20},
			row:   0,
			expectedResult: func() []float64 {
				v := make([]float64, 20)

				for i := 0; i < 20; i++ {
					v[i] = float64(i)
				}

				return v
			}(),
		},
		{
			name:           "Shape (20,1)",
			shape:          [2]int{20, 1},
			row:            5,
			expectedResult: []float64{5},
		},
		{
			name:  "Shape (20,20)",
			shape: [2]int{20, 20},
			row:   12,
			expectedResult: func() []float64 {
				v := make([]float64, 20)
				start := 20 * 12
				end := start + 20

				var index int = 0
				for i := start; i < end; i++ {
					v[index] = float64(i)
					index++
				}

				return v
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.NewMatrixFromSlice(tc.shape[0], tc.shape[1], baseVector[:size])
			row := nm.GetRow(tc.row)

			if !matrix.AreEqual(row, tc.expectedResult) {
				t.Errorf("\nexpected vector: %v | got: %v \n%v", tc.expectedResult, row, nm)
			}
		})
	}
}

func TestMatrix_GetColumn(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          [2]int
		col            int
		expectedResult []float64
	}{
		{
			name:           "Shape (1,1)",
			shape:          [2]int{1, 1},
			col:            0,
			expectedResult: []float64{0},
		},
		{
			name:           "Shape (1,20)",
			shape:          [2]int{1, 20},
			col:            6,
			expectedResult: []float64{6},
		},
		{
			name:  "Shape (20,1)",
			shape: [2]int{20, 1},
			col:   0,
			expectedResult: func() []float64 {
				v := make([]float64, 20)

				for i := 0; i < 20; i++ {
					v[i] = float64(i)
				}

				return v
			}(),
		},
		{
			name:  "Shape (20,20)",
			shape: [2]int{20, 20},
			col:   12,
			expectedResult: func() []float64 {
				v := make([]float64, 20)

				value := float64(12)
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
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.NewMatrixFromSlice(tc.shape[0], tc.shape[1], baseVector[:size])
			col := nm.GetColumn(tc.col)

			if !matrix.AreEqual(col, tc.expectedResult) {
				t.Errorf("\nexpected vector: %v | got: %v \n%v", tc.expectedResult, col, nm)
			}
		})
	}
}

func TestMatrix_Randomize(t *testing.T) {
	sizeX := 2000
	sizeY := 2000
	nm := matrix.NewMatrix(sizeY, sizeX)
	nm.Randomize()

	probabilityDistribution := matrix.SoftMax(nm.Data)
	_, min := matrix.Min(probabilityDistribution)
	_, max := matrix.Max(probabilityDistribution)

	maxDelta := 5e-7
	actualDelta := max - min

	testName := fmt.Sprintf("Randomize keeping Normal distribution Shape (%d,%d) MaxDelta %v", sizeY, sizeX, maxDelta)

	t.Run(testName, func(t *testing.T) {

		if actualDelta > maxDelta {
			t.Errorf("\nexpected maximum difference: %v | got: %v", maxDelta, actualDelta)
		}
	})
}

func TestMatrix_Scale(t *testing.T) {
	scalar := -2.0
	sizeX := 10
	sizeY := 10
	baseVector := make([]float64, sizeX*sizeY)

	for i := 0; i < len(baseVector); i++ {
		baseVector[i] = 10
	}

	nm := matrix.WrapSlice(sizeY, sizeX, baseVector).Scale(scalar)

	expectedMatrix := matrix.NewMatrix(sizeY, sizeX)
	for i := 0; i < len(expectedMatrix.Data); i++ {
		expectedMatrix.Data[i] = -20
	}

	testName := fmt.Sprintf("Scale matrix %v Shape (%d,%d)", scalar, sizeY, sizeX)

	t.Run(testName, func(t *testing.T) {
		if !expectedMatrix.IsEqual(nm) {
			t.Errorf("\nexpected : \n%v got: \n%v", expectedMatrix, nm)
		}
	})
}

func TestMatrix_MapFunction(t *testing.T) {
	f := func(v float64) float64 { return 10 }
	sizeX := 10
	sizeY := 10

	nm := matrix.NewMatrix(sizeY, sizeX).MapFunc(f)

	expectedMatrix := matrix.NewMatrix(sizeY, sizeX)
	for i := 0; i < len(expectedMatrix.Data); i++ {
		expectedMatrix.Data[i] = 10
	}

	testName := fmt.Sprintf("Map matrix Shape (%d,%d)", sizeY, sizeX)

	t.Run(testName, func(t *testing.T) {
		if !expectedMatrix.IsEqual(nm) {
			t.Errorf("\nexpected : \n%v got: \n%v", expectedMatrix, nm)
		}
	})
}

func TestMatrix_Transpose(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          [2]int
		expectedResult *matrix.MatrixFloat64
	}{
		{
			name:  "Shape (1,1)",
			shape: [2]int{1, 1},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 1,
				Data:    []float64{0},
			},
		},
		{
			name:  "Shape (1,10) -> (10,1)",
			shape: [2]int{1, 10},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    10,
				Columns: 1,
				Data:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name:  "Shape (10,1) -> (1,10)",
			shape: [2]int{10, 1},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 10,
				Data:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name:  "Shape (5,5)",
			shape: [2]int{5, 5},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    5,
				Columns: 5,
				Data: []float64{
					0, 5, 10, 15, 20,
					1, 6, 11, 16, 21,
					2, 7, 12, 17, 22,
					3, 8, 13, 18, 23,
					4, 9, 14, 19, 24,
				},
			},
		},
		{
			name:  "Shape (16,25) -> (25,16)",
			shape: [2]int{16, 25},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    25,
				Columns: 16,
				Data: func() []float64 {
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
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.WrapSlice(tc.shape[0], tc.shape[1], baseVector[:size]).Transpose()

			if !tc.expectedResult.IsEqual(nm) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nm)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {
	baseVector := make([]float64, 400)

	for i := 0; i < 400; i++ {
		baseVector[i] = float64(i)
	}

	testCases := []struct {
		name           string
		shape          [2]int
		times          int
		expectedResult *matrix.MatrixFloat64
	}{
		{
			name:  "Shape (1,1)",
			shape: [2]int{1, 1},
			times: 20,
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 1,
				Data:    []float64{0},
			},
		},
		{
			name:  "Shape (1,10)",
			shape: [2]int{1, 10},
			times: 4,
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 10,
				Data: func() []float64 {
					nv := make([]float64, 10)
					for i := 0; i < len(nv); i++ {
						nv[i] = float64(4 * i)
					}
					return nv
				}(),
			},
		},
		{
			name:  "Shape (10,1)",
			shape: [2]int{10, 1},
			times: 15,
			expectedResult: &matrix.MatrixFloat64{
				Rows:    10,
				Columns: 1,
				Data: func() []float64 {
					nv := make([]float64, 10)
					for i := 0; i < len(nv); i++ {
						nv[i] = float64(15 * i)
					}
					return nv
				}(),
			},
		},
		{
			name:  "Shape (5,5)",
			shape: [2]int{5, 5},
			times: 2,
			expectedResult: &matrix.MatrixFloat64{
				Rows:    5,
				Columns: 5,
				Data: func() []float64 {
					nv := make([]float64, 25)
					for i := 0; i < len(nv); i++ {
						nv[i] = float64(2 * i)
					}
					return nv
				}(),
			},
		},
		{
			name:  "Shape (16,25)",
			shape: [2]int{16, 25},
			times: 10,
			expectedResult: &matrix.MatrixFloat64{
				Rows:    16,
				Columns: 25,
				Data: func() []float64 {
					nv := make([]float64, 400)
					for i := 0; i < len(nv); i++ {
						nv[i] = float64(10 * i)
					}
					return nv
				}(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.WrapSlice(tc.shape[0], tc.shape[1], baseVector[:size])

			nmSum := matrix.NewMatrix(tc.shape[0], tc.shape[1])
			for i := 0; i < tc.times; i++ {
				nmSum = nmSum.Add(nm)
			}

			if !tc.expectedResult.IsEqual(nmSum) {
				t.Errorf("\nexpected matrix: %v \ngot: \n%v", tc.expectedResult, nmSum)
			}
		})
	}
}

func TestMatrix_AddVectorPerRow(t *testing.T) {
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
			vector: []float64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    1,
				Columns: 10,
				Data:    []float64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			},
		},
		{
			name:   "Shape (10,1)",
			shape:  [2]int{10, 1},
			vector: []float64{10},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    10,
				Columns: 1,
				Data:    []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
			},
		},
		{
			name:   "Shape (5,5)",
			shape:  [2]int{5, 5},
			vector: []float64{-10, 10, -10, 10, -10},
			expectedResult: &matrix.MatrixFloat64{
				Rows:    5,
				Columns: 5,
				Data: func() []float64 {
					nv := make([]float64, 25)
					for row := 0; row < 5; row++ {
						value := -10.0
						for col := 0; col < 5; col++ {
							index := 5*row + col
							nv[index] = float64(index) + value
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
				nv := make([]float64, 25)
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
					for row := 0; row < 16; row++ {
						value := -10.0
						for col := 0; col < 25; col++ {
							index := 25*row + col
							nv[index] = float64(index) + value
							value *= -1
						}
					}
					return nv
				}(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.shape[0] * tc.shape[1]
			nm := matrix.WrapSlice(tc.shape[0], tc.shape[1], baseVector[:size]).AddVectorPerRow(tc.vector)

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
