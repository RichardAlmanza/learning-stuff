package denselayer

import (
	"math"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

var (
	reLuFunctionFloat64 *ActivationFunction[float64] = nil
	reLuFunctionFloat32 *ActivationFunction[float32] = nil

	linearFunctionFloat64 *ActivationFunction[float64] = nil
	linearFunctionFloat32 *ActivationFunction[float32] = nil

	sigmoidFunctionFloat64 *ActivationFunction[float64] = nil
	sigmoidFunctionFloat32 *ActivationFunction[float32] = nil

	softMaxFunctionFloat64 *ActivationFunction[float64] = nil
	softMaxFunctionFloat32 *ActivationFunction[float32] = nil

	softMaxWCrossEntropyFunctionFloat64 *ActivationFunction[float64] = nil
	softMaxWCrossEntropyFunctionFloat32 *ActivationFunction[float32] = nil
)

type ActivationFunction[T vector.Real] struct {
	Function   func(*matrix.Matrix[T], *Layer[T]) *matrix.Matrix[T]
	Derivative func(*matrix.Matrix[T], *Layer[T]) *matrix.Matrix[T]
}

func (af *ActivationFunction[T]) Forward(input *matrix.Matrix[T], layer *Layer[T]) {
	layer.LastInput = input.Copy()
	layer.LastOutput = af.Function(input.Product(layer.TWeights).AddVectorPerRow(layer.Biases), layer)

	// return layer.LastOutput
}

func (af *ActivationFunction[T]) Backward(dValues *matrix.Matrix[T], layer *Layer[T]) {
	dFunction := af.Derivative(dValues, layer).Scale(1 / float64(dValues.Shape()[0]))

	layer.DInputs = dFunction.Product(layer.TWeights.Transpose())
	layer.DWeights = layer.LastInput.Transpose().Product(dFunction)
	layer.DBiases = make([]T, len(layer.Biases))

	for i := 0; i < len(layer.DBiases); i++ {
		layer.DBiases[i] = vector.Sum(dFunction.GetColumn(i))
	}

	// return l.DInputs
}

func sigmoid[T vector.Real](r T) T {
	return T(1 / (1 + math.Pow(math.E, float64(r))))
}

func newReLu[T vector.Real]() *ActivationFunction[T] {
	return &ActivationFunction[T]{
		Function: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return m.MapFunc(func(_ int, f T) T { return T(math.Max(0, float64(f))) })
		},
		Derivative: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return m.MapFunc(func(_ int, f T) T {
				if f > 0 {
					return 1
				}
				return 0
			})
		},
	}
}

func newLinear[T vector.Real]() *ActivationFunction[T] {
	return &ActivationFunction[T]{
		Function: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] { return m },
		Derivative: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return m.MapFunc(func(_ int, f T) T { return 1 })
		},
	}
}

func newSigmoid[T vector.Real]() *ActivationFunction[T] {
	return &ActivationFunction[T]{
		Function: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return m.MapFunc(func(_ int, f T) T { return sigmoid(f) })
		},
		Derivative: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return m.MapFunc(func(_ int, f T) T {
				sig := sigmoid(f)
				return sig * (1 - sig)
			})
		},
	}
}

func newSoftMax[T vector.Real]() *ActivationFunction[T] {
	return &ActivationFunction[T]{
		Function: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return matrix.SoftMax(m)
		},
		Derivative: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			// Jacobian matrices
			newMatrices := make([]matrix.Matrix[T], m.Shape()[0])

			for row := 0; row < m.Shape()[0]; row++ {
				currRow := m.GetRow(row)

				diagMatrix := matrix.NewMatrixDiagonalFromSlice(currRow)
				m1 := matrix.WrapSlice([]int{m.Shape()[1], 1}, currRow)
				m1 = m1.Product(m1.Transpose()).Scale(-1)

				newMatrices[row] = *diagMatrix.Add(m1)
			}

			return &newMatrices[0]
		},
	}
}

func newSoftMaxWCrossEntropy[T vector.Real]() *ActivationFunction[T] {
	return &ActivationFunction[T]{
		Function: func(m *matrix.Matrix[T], _ *Layer[T]) *matrix.Matrix[T] {
			return matrix.SoftMax(m)
		},
		Derivative: func(m *matrix.Matrix[T], l *Layer[T]) *matrix.Matrix[T] {
			return l.LastInput.DerivativeCrossEntropyLossSoftMaxPerRow(m.ToOneHot())
		},
	}
}

func singletonAF[T vector.Real](af *ActivationFunction[T], newAF func() *ActivationFunction[T]) *ActivationFunction[T] {
	if af == nil {
		af = newAF()
	}

	return af
}

// Exported functions for creating ActivationFunctions.

func NewReLuFloat64() *ActivationFunction[float64] {
	return singletonAF(reLuFunctionFloat64, newReLu)
}

func NewReLuFloat32() *ActivationFunction[float32] {
	return singletonAF(reLuFunctionFloat32, newReLu)
}

func NewLinearFloat64() *ActivationFunction[float64] {
	return singletonAF(linearFunctionFloat64, newLinear)
}

func NewLinearFloat32() *ActivationFunction[float32] {
	return singletonAF(linearFunctionFloat32, newLinear)
}
func NewSigmoidFloat64() *ActivationFunction[float64] {
	return singletonAF(sigmoidFunctionFloat64, newSigmoid)
}

func NewSigmoidFloat32() *ActivationFunction[float32] {
	return singletonAF(sigmoidFunctionFloat32, newSigmoid)
}

func NewSoftMaxFloat64() *ActivationFunction[float64] {
	return singletonAF(softMaxFunctionFloat64, newSoftMax)
}

func NewSoftMaxFloat32() *ActivationFunction[float32] {
	return singletonAF(softMaxFunctionFloat32, newSoftMax)
}

func NewSoftMaxWCrossEntropyFloat64() *ActivationFunction[float64] {
	return singletonAF(softMaxWCrossEntropyFunctionFloat64, newSoftMaxWCrossEntropy)
}

func NewSoftMaxWCrossEntropyFloat32() *ActivationFunction[float32] {
	return singletonAF(softMaxWCrossEntropyFunctionFloat32, newSoftMaxWCrossEntropy)
}
