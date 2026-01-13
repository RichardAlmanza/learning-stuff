package denselayer

import (
	"math"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

var (
	sigmoid func(float64) float64 = func(f float64) float64 { return 1 / (1 + math.Pow(math.E, f)) }

	ReLuFunction *ActivationFunction = &ActivationFunction{
		Function: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.MapFunc(func(f float64) float64 { return math.Max(0, f) })
		},
		Derivative: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.MapFunc(func(f float64) float64 {
				if f > 0 {
					return 1
				}
				return 0
			})
		},
	}

	LinearFunction *ActivationFunction = &ActivationFunction{
		Function: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 { return m },
		Derivative: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.MapFunc(func(f float64) float64 { return 1 })
		},
	}

	SigmoidFunction *ActivationFunction = &ActivationFunction{
		Function: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.MapFunc(func(f float64) float64 { return sigmoid(f) })
		},
		Derivative: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.MapFunc(func(f float64) float64 {
				sig := sigmoid(f)
				return sig * (1 - sig)
			})
		},
	}

	SoftMaxFunction *ActivationFunction = &ActivationFunction{
		Function: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return m.SoftMax()
		},
		Derivative: func(m *matrix.MatrixFloat64, _ *Layer) *matrix.MatrixFloat64 {
			return &m.DerivativeSoftMax()[0]
		},
	}
)

type ActivationFunction struct {
	Function   func(*matrix.MatrixFloat64, *Layer) *matrix.MatrixFloat64
	Derivative func(*matrix.MatrixFloat64, *Layer) *matrix.MatrixFloat64
}

func (af *ActivationFunction) Forward(input *matrix.MatrixFloat64, layer *Layer) {
	layer.LastInput = input.Copy()
	layer.LastOutput = af.Function(input.Product(layer.TWeights).AddVectorPerRow(layer.Biases), layer)

	// return layer.LastOutput
}

func (af *ActivationFunction) Backward(dValues *matrix.MatrixFloat64, layer *Layer) {
	dFunction := af.Derivative(dValues, layer)

	layer.DInputs = dFunction.Product(layer.TWeights.Transpose())
	layer.DWeights = layer.LastInput.Transpose().Product(dFunction)
	layer.DBiases = make([]float64, len(layer.Biases))

	for i := 0; i < len(layer.DBiases); i++ {
		layer.DBiases[i] = matrix.Sum(dFunction.GetColumn(i))
	}

	// return l.DInputs
}
