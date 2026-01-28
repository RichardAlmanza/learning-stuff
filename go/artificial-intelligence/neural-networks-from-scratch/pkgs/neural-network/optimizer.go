package neuralnetwork

import (
	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Optimizer[T vector.Real] struct {
	LearningRate T
}

type Batch[T vector.Real] struct {
	Input          *matrix.Matrix[T]
	ExpectedResult *matrix.Matrix[T]
}

func (o *Optimizer[T]) UpdateParameters(l *denselayer.Layer[T]) {
	l.TWeights = l.DWeights.Scale(-float64(o.LearningRate)).Add(l.TWeights)
	l.Biases = vector.Map2Func(l.DBiases, l.Biases, func(_ int, dBias, bias T) T { return -o.LearningRate*dBias + bias })
}
