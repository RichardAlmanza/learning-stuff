package denselayer

import (
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

type Layer struct {
	TWeights           *matrix.MatrixFloat64
	Biases             []float64
	LastInput          *matrix.MatrixFloat64
	LastOutput         *matrix.MatrixFloat64
	DInputs            *matrix.MatrixFloat64
	DWeights           *matrix.MatrixFloat64
	DBiases            []float64
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
	l.ActivationFunction.Forward(input, l)

	return l.LastOutput
}

func (l *Layer) Backward(dValues *matrix.MatrixFloat64) *matrix.MatrixFloat64 {
	l.ActivationFunction.Backward(dValues, l)

	return l.DInputs
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
