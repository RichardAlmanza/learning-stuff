package vector_test

import (
	"testing"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func TestDot_RealNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		vectorA  []float64
		vectorB  []float64
		expected float64
	}{
		{
			name:     "3x3 vectors",
			vectorA:  []float64{-0.26, 65, 12},
			vectorB:  []float64{1, -2, 0.3},
			expected: -126.66,
		},
		{
			name:     "0x0 vectors",
			vectorA:  []float64{},
			vectorB:  []float64{},
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vector.DotProduct(tc.vectorA, tc.vectorB)

			if result != tc.expected {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestSum_Integers(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []int
		expected int
	}{
		{
			name:     "9 vector",
			vector:   []int{-26, 6, 12, 15, 10, -10, -25, 59, -20},
			expected: 21,
		},
		{
			name:     "Empty vector",
			vector:   []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vector.Sum(tc.vector)

			if result != tc.expected {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestSum_SmallReal(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []float32
		expected float32
	}{
		{
			name:     "3 vector",
			vector:   []float32{-26, 12e-4, 121},
			expected: 95.0012,
		},
		{
			name:     "Empty vector",
			vector:   []float32{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vector.Sum(tc.vector)

			if result != tc.expected {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestMapFunc_RealNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []float64
		function func(int, float64) float64
		expected []float64
	}{
		{
			name:     "Adding 1 to each element",
			vector:   []float64{-0.2, -5, 2},
			function: func(i int, f float64) float64 { return f + 1 },
			expected: []float64{0.8, -4, 3},
		},
		{
			name:     "Square the element of the vector",
			vector:   []float64{-0.5, -5, 2, 3},
			function: func(i int, f float64) float64 { return f * f },
			expected: []float64{0.25, 25, 4, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vector.MapFunc(tc.vector, tc.function)

			if !vector.AreEqual(result, tc.expected) {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestFilterFunc_32Integers(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []int32
		function func(int, int32) bool
		expected []int32
	}{
		{
			name:   "Filter even numbers",
			vector: []int32{1, 2, 3, 4, 5, 6},
			function: func(i int, n int32) bool {
				return n%2 == 0
			},
			expected: []int32{2, 4, 6},
		},
		{
			name:   "Filter numbers greater than 3",
			vector: []int32{1, 2, 3, 4, 5, 6},
			function: func(i int, n int32) bool {
				return n > 3
			},
			expected: []int32{4, 5, 6},
		},
		{
			name:   "Filter every third position element",
			vector: []int32{-1, 0, 1, 2, 3, 4, 5, 6},
			function: func(i int, n int32) bool {
				return (i+1)%3 == 0
			},
			expected: []int32{1, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := vector.FilterFunc(tc.vector, tc.function)

			if !vector.AreEqual(result, tc.expected) {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestMax_Integers(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []int
		expected []int
	}{
		{
			name:     "Max integers in vector of negatives numbers",
			vector:   []int{-2, -3, -4, -1, -50, -60},
			expected: []int{3, -1},
		},
		{
			name:     "Max integers in vector of positives numbers",
			vector:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{5, 6},
		},
		{
			name:     "Max integers in vector of mixed numbers",
			vector:   []int{-10, 20, -3, 40, -5, 6},
			expected: []int{3, 40},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, v := vector.Max(tc.vector)
			result := []int{i, v}

			if !vector.AreEqual(result, tc.expected) {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}

func TestMin_SmallReal(t *testing.T) {
	testCases := []struct {
		name     string
		vector   []float32
		expected []float32
	}{
		{
			name:     "Min Float32 in vector of negatives numbers",
			vector:   []float32{-2.25, -3.9, -4.001, -0.002, -50, -60},
			expected: []float32{5, -60},
		},
		{
			name:     "Min Float32 in vector of positives numbers",
			vector:   []float32{1, 2, 3, 4, 5, 6},
			expected: []float32{0, 1},
		},
		{
			name:     "Min Float32 in vector of mixed numbers",
			vector:   []float32{-10, 20, -3, 40, -5, 6},
			expected: []float32{0, -10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			i, v := vector.Min(tc.vector)
			result := []float32{float32(i), v}

			if !vector.AreEqual(result, tc.expected) {
				t.Errorf("\nexpected result: %v \ngot: \n%v", tc.expected, result)
			}
		})
	}
}
