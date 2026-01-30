package denselayer

import (
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Layer[T vector.Real] struct {
	TWeights *matrix.Matrix[T]
	Biases   []T
	Input    *matrix.Matrix[T]
	Output   *matrix.Matrix[T]
	DInputs  *matrix.Matrix[T]
	DWeights *matrix.Matrix[T]
	DBiases  []T
}

func NewLayer[T vector.Real](nNeurons, inputSize int) *Layer[T] {
	shape := []int{inputSize, nNeurons}

	return &Layer[T]{
		TWeights: matrix.NewMatrixRand[T](shape).Scale(0.01),
		Biases:   make([]T, nNeurons),
	}
}

func NewLayerFrom[T vector.Real](tWeights *matrix.Matrix[T], biases []T) *Layer[T] {
	if len(biases) != tWeights.Shape()[1] {
		panic("Size mismatch")
	}

	return &Layer[T]{
		TWeights: tWeights.Copy(),
		Biases:   vector.NewCopy(biases),
	}
}

func (l *Layer[T]) Forward(input *matrix.Matrix[T]) *matrix.Matrix[T] {
	l.Input = input.Copy()
	l.Output = input.Product(l.TWeights).AddVectorPerRow(l.Biases)

	return l.Output
}

func (l *Layer[T]) Backward(dValues *matrix.Matrix[T]) *matrix.Matrix[T] {
	// Dinputs are the Gradients
	l.DInputs = dValues.Product(l.TWeights.Transpose())
	l.DWeights = l.Input.Transpose().Product(dValues)
	l.DBiases = make([]T, len(l.Biases))

	for i := 0; i < len(l.Biases); i++ {
		l.DBiases[i] = vector.Sum(dValues.GetColumn(i))
	}

	return l.DInputs
}

func (l *Layer[T]) Copy() *Layer[T] {
	return &Layer[T]{
		TWeights: l.TWeights.Copy(),
		Biases:   vector.NewCopy(l.Biases),
	}
}
