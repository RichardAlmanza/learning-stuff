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
)

type ActivationFunction struct {
	Function   func(*matrix.MatrixFloat64, *Layer) *matrix.MatrixFloat64
	Derivative func(*matrix.MatrixFloat64, *Layer) *matrix.MatrixFloat64
}

type Layer struct {
	TWeights           *matrix.MatrixFloat64
	Biases             []float64
	LastInput          *matrix.MatrixFloat64
	LastOutput         *matrix.MatrixFloat64
	DInputs            *matrix.MatrixFloat64
	DWeights           *matrix.MatrixFloat64
	DBiases            []float64
	DeltaParameters    float64
	ActivationFunction *ActivationFunction
}

func NewLayer(nNeurons, inputSize int, function *ActivationFunction) *Layer {
	newTWeights := matrix.NewMatrixRand(inputSize, nNeurons).Scale(0.01)

	return &Layer{
		TWeights:           newTWeights,
		Biases:             make([]float64, nNeurons),
		ActivationFunction: function,
	}
}

func NewLayerFrom(tWeights *matrix.MatrixFloat64, biases []float64, function *ActivationFunction) *Layer {
	if len(biases) != tWeights.Columns {
		panic("Size mismatch")
	}

	return &Layer{
		TWeights:           tWeights,
		Biases:             biases,
		ActivationFunction: function,
	}
}

func (l *Layer) Forward(input *matrix.MatrixFloat64) *matrix.MatrixFloat64 {
	l.LastInput = input.Copy()
	l.LastOutput = l.ActivationFunction.Function(input.Product(l.TWeights).AddVectorPerRow(l.Biases), l)

	return l.LastOutput
}

func (l *Layer) Backward(dValues *matrix.MatrixFloat64) *matrix.MatrixFloat64 {
	dFunction := l.ActivationFunction.Derivative(dValues, l)

	l.DInputs = dFunction.Product(l.TWeights.Transpose())
	l.DWeights = l.LastInput.Transpose().Product(dFunction)
	l.DBiases = make([]float64, len(l.Biases))

	for i := 0; i < len(l.DBiases); i++ {
		l.DBiases[i] = matrix.Sum(dFunction.GetColumn(i))
	}

	return l.DInputs
}

func (l *Layer) UpdateParameters() {
	l.TWeights = l.DWeights.Scale(-l.DeltaParameters).Add(l.TWeights)
	l.Biases = matrix.Map2VectorFunc(l.DBiases, l.Biases, func(f1, f2 float64) float64 { return -l.DeltaParameters*f1 + f2 })
}

func (l *Layer) Copy() *Layer {
	newBiases := make([]float64, len(l.Biases))
	copy(newBiases, l.Biases)

	return &Layer{
		TWeights:           l.TWeights.Copy(),
		Biases:             newBiases,
		ActivationFunction: l.ActivationFunction,
	}
}
