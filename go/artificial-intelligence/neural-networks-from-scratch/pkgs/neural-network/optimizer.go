package neuralnetwork

import (
	"math"

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
	// l.weightsMomentum = momentum * l.weightsMomentum - LearningRate * dWeights
	// l.Weights += l.weightsMomentum

	// A = -LearningRate * dWeights
	wUpdates := l.DBase.W.Scale(-float64(oSGD.LearningRate))
	bUpdates := vector.MapFunc(l.DBase.B, func(_ int, dBias T) T { return -oSGD.LearningRate * dBias })

	if oSGD.Momentum != 0 {
		// Create layer momentum if it doesn't exist.
		if l.Momentum.W == nil || l.Momentum.B == nil {
			l.Momentum.W = matrix.NewMatrix[T](l.Base.W.Shape())
			l.Momentum.B = make([]T, len(l.Base.B))
		}

		// Update momentums.
		// B = momentum * l.WeightsMomentum
		// A += B
		wUpdates = l.Momentum.W.Scale(float64(oSGD.Momentum)).Add(wUpdates)
		// l.WeightsMomentum = A
		l.Momentum.W = wUpdates

		bUpdates = vector.Map2Func(bUpdates, l.Momentum.B, func(_ int, bUpdate, bMomentum T) T { return bUpdate + oSGD.Momentum*bMomentum })
		l.Momentum.B = bUpdates
	}

	// l.Weights += l.weightsMomentum
	l.Base.W = l.Base.W.Add(wUpdates)
	l.Base.B = vector.AddVectors(l.Base.B, bUpdates)
}

type AdaptiveGradientOptimizer[T vector.Real] struct {
	LearningRate        T
	InitialLearningRate T
	Decay               T
	Epsilon             T
}

func NewAdaGrad[T vector.Real](learningRate, decay, epsilon T) *AdaptiveGradientOptimizer[T] {
	return &AdaptiveGradientOptimizer[T]{
		LearningRate:        learningRate,
		InitialLearningRate: learningRate,
		Decay:               decay,
		Epsilon:             epsilon,
	}
}

func (oAG *AdaptiveGradientOptimizer[T]) UpdateLearningRate(epoch int) {
	if oAG.Decay != 0 {
		oAG.LearningRate = oAG.InitialLearningRate / (1 + oAG.Decay*T(epoch))
	}
}

func (oAG *AdaptiveGradientOptimizer[T]) UpdateParameters(l *denselayer.Layer[T]) {
	// weightsCache += dWeights^2 ; The power operation is element-wise, not matrix-wise
	// l.weights += -CurrentLearningRate * dWeights / (epsilon + sqrt(weightsCache))
	fADivB := func(_ int, a, b T) T { return a / b }
	fAPlusB2 := func(_ int, c, a T) T { return c + T(math.Pow(float64(a), 2)) }
	fEpsilonPlusSqrtA := func(_ int, a T) T { return oAG.Epsilon + T(math.Sqrt(float64(a))) }

	// Create layer cache if it doesn't exist.
	if l.Cache.W == nil || l.Cache.B == nil {
		l.Cache.W = matrix.NewMatrix[T](l.Base.W.Shape())
		l.Cache.B = make([]T, len(l.Base.B))
	}

	// weightsCache += dWeights^2
	l.Cache.W = matrix.WrapSlice(l.Cache.W.Shape(),
		vector.Map2Func(l.Cache.W.Data, l.DBase.W.Data, fAPlusB2))
	l.Cache.B = vector.Map2Func(l.Cache.B, l.DBase.B, fAPlusB2)

	// A = -CurrentLearningRate * dWeights
	wUpdates := l.DBase.W.Scale(-float64(oAG.LearningRate))
	bUpdates := vector.MapFunc(l.DBase.B, func(_ int, dBias T) T { return -oAG.LearningRate * dBias })

	// B = Epsilon + sqrt(weightsCache)
	sqrtEpsilonW := l.Cache.W.MapFunc(fEpsilonPlusSqrtA)
	sqrtEpsilonB := vector.MapFunc(l.Cache.B, fEpsilonPlusSqrtA)

	// A /= B
	wUpdates = matrix.WrapSlice(wUpdates.Shape(),
		vector.Map2Func(wUpdates.Data, sqrtEpsilonW.Data, fADivB))
	bUpdates = vector.Map2Func(bUpdates, sqrtEpsilonB, fADivB)

	// Update weights and biases
	l.Base.W = l.Base.W.Add(wUpdates)
	l.Base.B = vector.AddVectors(l.Base.B, bUpdates)
}

type Batch[T vector.Real] struct {
	Input          *matrix.Matrix[T]
	ExpectedResult *matrix.Matrix[T]
}
