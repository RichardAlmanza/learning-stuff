package denselayer

import (
	"math"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

var (
	reLuFunctionFloat64 *ReLu[float64] = nil
	reLuFunctionFloat32 *ReLu[float32] = nil

	linearFunctionFloat64 *Linear[float64] = nil
	linearFunctionFloat32 *Linear[float32] = nil

	sigmoidFunctionFloat64 *Sigmoid[float64] = nil
	sigmoidFunctionFloat32 *Sigmoid[float32] = nil

	softMaxFunctionFloat64 *SoftMax[float64] = nil
	softMaxFunctionFloat32 *SoftMax[float32] = nil

	softMaxWithCrossEntropyFunctionFloat64 *SoftMaxWithCrossEntropy[float64] = nil
	softMaxWithCrossEntropyFunctionFloat32 *SoftMaxWithCrossEntropy[float32] = nil
)

type ActivationFunction[T vector.Real] struct {
	Input  *matrix.Matrix[T]
	Output *matrix.Matrix[T]
	DInput *matrix.Matrix[T]
}

type AFConstraint[T vector.Real] interface {
	ActivationFunction[T] | ReLu[T] | Linear[T] |
		Sigmoid[T] | SoftMax[T] | SoftMaxWithCrossEntropy[T]
}

type ReLu[T vector.Real] ActivationFunction[T]
type Linear[T vector.Real] ActivationFunction[T]
type Sigmoid[T vector.Real] ActivationFunction[T]
type SoftMax[T vector.Real] ActivationFunction[T]
type SoftMaxWithCrossEntropy[T vector.Real] ActivationFunction[T]

func NewReLu[T vector.Real]() *ReLu[T]       { return &ReLu[T]{} }
func NewLinear[T vector.Real]() *Linear[T]   { return &Linear[T]{} }
func NewSigmoid[T vector.Real]() *Sigmoid[T] { return &Sigmoid[T]{} }
func NewSoftMax[T vector.Real]() *SoftMax[T] { return &SoftMax[T]{} }
func NewSoftMaxWithCrossEntropy[T vector.Real]() *SoftMaxWithCrossEntropy[T] {
	return &SoftMaxWithCrossEntropy[T]{}
}

func (rl *ReLu[T]) Forward(input *matrix.Matrix[T]) {
	rl.Input = input.Copy()
	rl.Output = input.MapFunc(func(_ int, f T) T { return T(math.Max(0, float64(f))) })
}

func (rl *ReLu[T]) Backward(dValues *matrix.Matrix[T]) {
	rl.DInput = dValues.MapFunc(func(_ int, f T) T {
		if f > 0 {
			return 1
		}
		return 0
	})
}

func (l *Linear[T]) Forward(input *matrix.Matrix[T]) {
	l.Input = input.Copy()
	l.Output = input.Copy()
}

func (l *Linear[T]) Backward(dValues *matrix.Matrix[T]) {
	l.DInput = dValues.MapFunc(func(_ int, f T) T { return 1 })
}

func sigmoid[T vector.Real](r T) T {
	return T(1 / (1 + math.Pow(math.E, float64(r))))
}

func (s *Sigmoid[T]) Forward(input *matrix.Matrix[T]) {
	s.Input = input.Copy()
	s.Output = input.MapFunc(func(_ int, f T) T { return sigmoid(f) })
}

func (s *Sigmoid[T]) Backward(dValues *matrix.Matrix[T]) {
	s.DInput = dValues.MapFunc(func(_ int, f T) T {
		sig := sigmoid(f)
		return sig * (1 - sig)
	})
}

func (s *SoftMax[T]) Forward(input *matrix.Matrix[T]) {
	s.Input = input.Copy()
	s.Output = matrix.SoftMax(input)
}

func (s *SoftMax[T]) Backward(dValues *matrix.Matrix[T]) {
	// Jacobian matrices
	newMatrices := make([]matrix.Matrix[T], dValues.Shape()[0])

	for row := 0; row < dValues.Shape()[0]; row++ {
		currRow := dValues.GetRow(row)

		diagMatrix := matrix.NewMatrixDiagonalFromSlice(currRow)
		m1 := matrix.WrapSlice([]int{dValues.Shape()[1], 1}, currRow)
		m1 = m1.Product(m1.Transpose()).Scale(-1)

		newMatrices[row] = *diagMatrix.Add(m1)
	}

	s.DInput = &newMatrices[0]
}

func (s *SoftMaxWithCrossEntropy[T]) Forward(input *matrix.Matrix[T]) {
	s.Input = input.Copy()
	s.Output = matrix.SoftMax(input)
}

func (s *SoftMaxWithCrossEntropy[T]) Backward(dValues *matrix.Matrix[T], targets []int) {
	s.DInput = dValues.DerivativeCrossEntropyLossSoftMaxPerRow(targets)
}
