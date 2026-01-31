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

type RootMeanSquarePropagation[T vector.Real] struct {
	LearningRate        T
	InitialLearningRate T
	Decay               T
	Epsilon             T
	Rho                 T
}

func NewRMSProp[T vector.Real](learningRate, decay, epsilon, rho T) *RootMeanSquarePropagation[T] {
	return &RootMeanSquarePropagation[T]{
		LearningRate:        learningRate,
		InitialLearningRate: learningRate,
		Decay:               decay,
		Epsilon:             epsilon,
		Rho:                 rho,
	}
}

func (oRMSP *RootMeanSquarePropagation[T]) UpdateLearningRate(epoch int) {
	if oRMSP.Decay != 0 {
		oRMSP.LearningRate = oRMSP.InitialLearningRate / (1 + oRMSP.Decay*T(epoch))
	}
}

func (oRMSP *RootMeanSquarePropagation[T]) UpdateParameters(l *denselayer.Layer[T]) {
	// cache = rho * cache + (1 - rho) * dWeights ^ 2
	// l.weights += -CurrentLearningRate * dWeights / (epsilon + sqrt(weightsCache))
	fADivB := func(_ int, a, b T) T { return a / b }
	fCacheRhoPlusA2 := func(_ int, c, a T) T { return oRMSP.Rho*c + (1-oRMSP.Rho)*T(math.Pow(float64(a), 2)) }
	fEpsilonPlusSqrtA := func(_ int, a T) T { return oRMSP.Epsilon + T(math.Sqrt(float64(a))) }

	// Create layer cache if it doesn't exist.
	if l.Cache.W == nil || l.Cache.B == nil {
		l.Cache.W = matrix.NewMatrix[T](l.Base.W.Shape())
		l.Cache.B = make([]T, len(l.Base.B))
	}

	// cache = rho * cache + (1 - rho) * dWeights ^ 2
	l.Cache.W = matrix.WrapSlice(l.Cache.W.Shape(),
		vector.Map2Func(l.Cache.W.Data, l.DBase.W.Data, fCacheRhoPlusA2))
	l.Cache.B = vector.Map2Func(l.Cache.B, l.DBase.B, fCacheRhoPlusA2)

	// A = -CurrentLearningRate * dWeights
	wUpdates := l.DBase.W.Scale(-float64(oRMSP.LearningRate))
	bUpdates := vector.MapFunc(l.DBase.B, func(_ int, dBias T) T { return -oRMSP.LearningRate * dBias })

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

type AdaptiveMomentum[T vector.Real] struct {
	LearningRate        T
	InitialLearningRate T
	Decay               T
	Epsilon             T
	Beta1               T
	Beta2               T
}

func NewAdam[T vector.Real](learningRate, decay, epsilon, beta1, beta2 T) *AdaptiveMomentum[T] {
	return &AdaptiveMomentum[T]{
		LearningRate:        learningRate,
		InitialLearningRate: learningRate,
		Decay:               decay,
		Epsilon:             epsilon,
		Beta1:               beta1,
		Beta2:               beta2,
	}
}

func (oAM *AdaptiveMomentum[T]) UpdateLearningRate(epoch int) {
	if oAM.Decay != 0 {
		oAM.LearningRate = oAM.InitialLearningRate / (1 + oAM.Decay*T(epoch))
	}
}

func (oAM *AdaptiveMomentum[T]) UpdateParameters(l *denselayer.Layer[T], epoch int) {
	// cache = beta2 * cache + (1 - beta2) * dWeights^2
	// l.weights = -CurrentLearningRate * momentumCorrected / (epsilon + sqrt(cacheCorrected))
	fADivB := func(_ int, a, b T) T { return a / b }
	fCacheBeta2PlusA2 := func(_ int, c, a T) T { return oAM.Beta2*c + (1-oAM.Beta2)*T(math.Pow(float64(a), 2)) }
	fCorrectCache := func(_ int, m T) T { return m / T(1-math.Pow(float64(oAM.Beta2), float64(epoch+1))) }
	fMomentumDerivative := func(_ int, m, d T) T { return oAM.Beta1*m + (1-oAM.Beta1)*d }
	fCorrectMomentum := func(_ int, m T) T { return m / T(1-math.Pow(float64(oAM.Beta1), float64(epoch+1))) }
	fEpsilonPlusSqrtA := func(_ int, a T) T { return oAM.Epsilon + T(math.Sqrt(float64(a))) }

	// Create layer cache if it doesn't exist.
	if l.Cache.W == nil || l.Cache.B == nil {
		l.Cache.W = matrix.NewMatrix[T](l.Base.W.Shape())
		l.Cache.B = make([]T, len(l.Base.B))
	}

	// Create layer momentum if it doesn't exist.
	if l.Momentum.W == nil || l.Momentum.B == nil {
		l.Momentum.W = matrix.NewMatrix[T](l.Base.W.Shape())
		l.Momentum.B = make([]T, len(l.Base.B))
	}

	// momentum = beta1 * weightMomentum + (1 - beta1) * dWeights
	l.Momentum.W = matrix.WrapSlice(l.Momentum.W.Shape(),
		vector.Map2Func(l.Momentum.W.Data, l.DBase.W.Data, fMomentumDerivative))
	l.Momentum.B = vector.Map2Func(l.Momentum.B, l.DBase.B, fMomentumDerivative)

	// momentumCorrected = momentum / (1 - beta1 ^ (epoch + 1))
	momentumCorrectedW := l.Momentum.W.MapFunc(fCorrectMomentum)
	momentumCorrectedB := vector.MapFunc(l.Momentum.B, fCorrectMomentum)

	// cache = beta2 * cache + (1 - beta2) * dWeights^2
	l.Cache.W = matrix.WrapSlice(l.Cache.W.Shape(),
		vector.Map2Func(l.Cache.W.Data, l.DBase.W.Data, fCacheBeta2PlusA2))
	l.Cache.B = vector.Map2Func(l.Cache.B, l.DBase.B, fCacheBeta2PlusA2)

	// cacheCorrected = cache / (1 - beta2 ^ (epoch + 1))
	cacheCorrectedW := l.Cache.W.MapFunc(fCorrectCache)
	cacheCorrectedB := vector.MapFunc(l.Cache.B, fCorrectCache)

	// A = -CurrentLearningRate * momentumCorrected
	wUpdates := momentumCorrectedW.Scale(-float64(oAM.LearningRate))
	bUpdates := vector.MapFunc(momentumCorrectedB, func(_ int, dBias T) T { return -oAM.LearningRate * dBias })

	// B = Epsilon + sqrt(cacheCorrected)
	sqrtEpsilonW := cacheCorrectedW.MapFunc(fEpsilonPlusSqrtA)
	sqrtEpsilonB := vector.MapFunc(cacheCorrectedB, fEpsilonPlusSqrtA)

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
