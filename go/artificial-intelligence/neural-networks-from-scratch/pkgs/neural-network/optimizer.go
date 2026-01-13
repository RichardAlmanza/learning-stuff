package neuralnetwork

import (
	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

type Optimizer struct {
	LearningRate float64
}

type Batch struct {
	Input          *matrix.MatrixFloat64
	ExpectedResult *matrix.MatrixFloat64
}

func (o *Optimizer) UpdateParameters(l *denselayer.Layer) {
	l.TWeights = l.DWeights.Scale(-o.LearningRate).Add(l.TWeights)
	l.Biases = matrix.Map2VectorFunc(l.DBiases, l.Biases, func(f1, f2 float64) float64 { return -o.LearningRate*f1 + f2 })
}
