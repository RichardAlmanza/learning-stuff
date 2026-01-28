package denselayer

import (
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Layer[T vector.Real] struct {
	TWeights           *matrix.Matrix[T]
	Biases             []T
	LastInput          *matrix.Matrix[T]
	LastOutput         *matrix.Matrix[T]
	DInputs            *matrix.Matrix[T]
	DWeights           *matrix.Matrix[T]
	DBiases            []T
	ActivationFunction *ActivationFunction[T]
}

func NewLayer[T vector.Real](nNeurons, inputSize int, function *ActivationFunction[T]) *Layer[T] {
	shape := []int{inputSize, nNeurons}
	newTWeights := matrix.NewMatrixRand[T](shape).Scale(0.01)

	return &Layer[T]{
		TWeights:           newTWeights,
		Biases:             make([]T, nNeurons),
		ActivationFunction: function,
	}
}

func NewLayerFrom[T vector.Real](tWeights *matrix.Matrix[T], biases []T, function *ActivationFunction[T]) *Layer[T] {
	if len(biases) != tWeights.Shape()[1] {
		panic("Size mismatch")
	}

	return &Layer[T]{
		TWeights:           tWeights,
		Biases:             biases,
		ActivationFunction: function,
	}
}

func (l *Layer[T]) Forward(input *matrix.Matrix[T]) *matrix.Matrix[T] {
	l.ActivationFunction.Forward(input, l)

	return l.LastOutput
}

func (l *Layer[T]) Backward(dValues *matrix.Matrix[T]) *matrix.Matrix[T] {
	l.ActivationFunction.Backward(dValues, l)

	return l.DInputs
}

func (l *Layer[T]) Copy() *Layer[T] {
	return &Layer[T]{
		TWeights:           l.TWeights.Copy(),
		Biases:             vector.NewCopy(l.Biases),
		ActivationFunction: l.ActivationFunction,
	}
}
