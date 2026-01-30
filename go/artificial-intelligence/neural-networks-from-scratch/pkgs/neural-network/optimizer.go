package neuralnetwork

import (
	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Optimizer[T vector.Real] struct {
	LearningRate        T
	InitialLearningRate T
	Decay               T
}

type StochasticGradientDescent[T vector.Real] Optimizer[T]

func NewSGD[T vector.Real](learningRate, decay T) *StochasticGradientDescent[T] {
	return &StochasticGradientDescent[T]{
		LearningRate:        learningRate,
		InitialLearningRate: learningRate,
		Decay:               decay,
	}
}

func (oSGD *StochasticGradientDescent[T]) UpdateLearningRate(epoch int) {
	if oSGD.Decay != 0 {
		oSGD.LearningRate = oSGD.InitialLearningRate / (1 + oSGD.Decay*T(epoch))
	}
}

func (oSGD *StochasticGradientDescent[T]) UpdateParameters(l *denselayer.Layer[T]) {
	l.Base.W = l.DBase.W.Scale(-float64(oSGD.LearningRate)).Add(l.Base.W)
	l.Base.B = vector.Map2Func(l.DBase.B, l.Base.B, func(_ int, dBias, bias T) T { return -oSGD.LearningRate*dBias + bias })
}

type Batch[T vector.Real] struct {
	Input          *matrix.Matrix[T]
	ExpectedResult *matrix.Matrix[T]
}
