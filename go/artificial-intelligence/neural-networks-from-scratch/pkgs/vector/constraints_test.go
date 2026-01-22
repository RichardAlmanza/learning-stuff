package vector_test

import (
	"testing"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func TestAssertNumber(t *testing.T) {
	testCases := []struct {
		input    any
		expected vector.NumberType
	}{
		{
			input:    float32(1.0),
			expected: vector.Float32 | vector.RealNumber,
		},
		{
			input:    float64(1.0),
			expected: vector.Float64 | vector.RealNumber,
		},
		{
			input:    int(1),
			expected: vector.Int | vector.IntegerNumber,
		},
		{
			input:    int8(1),
			expected: vector.Int8 | vector.IntegerNumber,
		},
		{
			input:    int16(1),
			expected: vector.Int16 | vector.IntegerNumber,
		},
		{
			input:    int32(1),
			expected: vector.Int32 | vector.IntegerNumber,
		},
		{
			input:    int64(1),
			expected: vector.Int64 | vector.IntegerNumber,
		},
		{
			input:    "string",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		result := vector.AssertNumber(tc.input)

		if result != tc.expected {
			t.Errorf("AssertNumber(%v) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestShape_TotalSize(t *testing.T) {
	testCases := []struct {
		input    vector.Shape
		expected int
	}{
		{[]int{2, 3}, 6},
		{[]int{1, 2, 3}, 6},
		{[]int{}, 1},
		{[]int{0, 5}, 0},
		{[]int{4, 0}, 0},
	}

	for _, tc := range testCases {
		result := tc.input.TotalSize()

		if result != tc.expected {
			t.Errorf("Shape(%v).TotalSize() = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestShape_Copy(t *testing.T) {
	testCases := []struct {
		input    vector.Shape
		expected vector.Shape
	}{
		{[]int{2, 3}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := tc.input.Copy()

		if !vector.AreEqual(result, tc.expected) {
			t.Errorf("Shape(%v).Copy() = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestShape_Dims(t *testing.T) {
	testCases := []struct {
		input    vector.Shape
		expected int
	}{
		{[]int{2, 3}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{}, 0},
	}

	for _, tc := range testCases {
		result := tc.input.Dims()

		if result != tc.expected {
			t.Errorf("Shape(%v).Dims() = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
