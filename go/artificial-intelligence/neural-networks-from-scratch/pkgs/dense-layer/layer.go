package denselayer

import (
	"math"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

var (
	ReLuFunction    ActivationFunction = func(f float64) float64 { return math.Max(0, f) }
	LinearFunction  ActivationFunction = func(f float64) float64 { return f }
	SigmoidFunction ActivationFunction = func(f float64) float64 { return 1 / (1 + math.Pow(math.E, f)) }
)

type ActivationFunction func(float64) float64

type Layer struct {
	Weights            *matrix.MatrixFloat64
	Biases             []float64
	ActivationFunction ActivationFunction
}

func NewLayer(nNeurons, inputSize int, function ActivationFunction) *Layer {
	newWeights := matrix.NewMatrixRand(inputSize, nNeurons).Scale(0.01)

	return &Layer{
		Weights:            newWeights,
		Biases:             make([]float64, nNeurons),
		ActivationFunction: function,
	}
}

func NewLayerFrom(weights *matrix.MatrixFloat64, biases []float64, function ActivationFunction) *Layer {
	if len(biases) != weights.Columns {
		panic("Size mismatch")
	}

	return &Layer{
		Weights:            weights,
		Biases:             biases,
		ActivationFunction: function,
	}
}

func (l *Layer) Forward(input *matrix.MatrixFloat64) *matrix.MatrixFloat64 {
	return input.Product(l.Weights).AddVectorPerRow(l.Biases).MapFunc(l.ActivationFunction)
}

func (l *Layer) Copy() *Layer {
	newWeights := matrix.NewMatrix(l.Weights.Rows, l.Weights.Columns)
	copy(newWeights.Data, l.Weights.Data)

	newBiases := make([]float64, len(l.Biases))
	copy(newBiases, l.Biases)

	return &Layer{
		Weights:            newWeights,
		Biases:             newBiases,
		ActivationFunction: l.ActivationFunction,
	}
}
