package neuralnetwork

import (
	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Optimizer[T vector.Real] struct {
	LearningRate T
}

type StochasticGradientDescent[T vector.Real] Optimizer[T]

func NewSGD[T vector.Real](learningRate T) *StochasticGradientDescent[T] {
	return &StochasticGradientDescent[T]{LearningRate: learningRate}
}
func (oSGD *StochasticGradientDescent[T]) UpdateParameters(l *denselayer.Layer[T]) {
	l.TWeights = l.DWeights.Scale(-float64(oSGD.LearningRate)).Add(l.TWeights)
	l.Biases = vector.Map2Func(l.DBiases, l.Biases, func(_ int, dBias, bias T) T { return -oSGD.LearningRate*dBias + bias })
}

type Batch[T vector.Real] struct {
	Input          *matrix.Matrix[T]
	ExpectedResult *matrix.Matrix[T]
}
