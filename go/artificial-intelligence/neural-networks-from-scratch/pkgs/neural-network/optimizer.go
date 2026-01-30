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
	Momentum            T
}

type StochasticGradientDescent[T vector.Real] Optimizer[T]

func NewSGD[T vector.Real](learningRate, decay, momentum T) *StochasticGradientDescent[T] {
	return &StochasticGradientDescent[T]{
		LearningRate:        learningRate,
		InitialLearningRate: learningRate,
		Decay:               decay,
		Momentum:            momentum,
	}
}

func (oSGD *StochasticGradientDescent[T]) UpdateLearningRate(epoch int) {
	if oSGD.Decay != 0 {
		oSGD.LearningRate = oSGD.InitialLearningRate / (1 + oSGD.Decay*T(epoch))
	}
}

func (oSGD *StochasticGradientDescent[T]) UpdateParameters(l *denselayer.Layer[T]) {
	wUpdates := l.DBase.W.Scale(-float64(oSGD.LearningRate))
	bUpdates := vector.MapFunc(l.DBase.B, func(_ int, dBias T) T { return -oSGD.LearningRate * dBias })

	if oSGD.Momentum != 0 {
		// Create layer momentum if it doesn't exist.
		if l.Momentum.W == nil || l.Momentum.B == nil {
			l.Momentum.W = matrix.NewMatrix[T](l.Base.W.Shape())
			l.Momentum.B = make([]T, len(l.Base.B))
		}

		// Update momentums.
		// Weights
		wUpdates = l.Momentum.W.Scale(float64(oSGD.Momentum)).Add(wUpdates)
		l.Momentum.W = wUpdates

		// Biases
		bUpdates = vector.Map2Func(bUpdates, l.Momentum.B, func(_ int, bUpdate, bMomentum T) T { return bUpdate + oSGD.Momentum*bMomentum })
		l.Momentum.B = bUpdates
	}

	l.Base.W = l.Base.W.Add(wUpdates)
	l.Base.B = vector.AddVectors(l.Base.B, bUpdates)
}

type Batch[T vector.Real] struct {
	Input          *matrix.Matrix[T]
	ExpectedResult *matrix.Matrix[T]
}
