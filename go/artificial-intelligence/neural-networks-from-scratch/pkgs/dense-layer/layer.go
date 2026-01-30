package denselayer

import (
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type LayerBase[T vector.Real] struct {
	W *matrix.Matrix[T]
	B []T
}

func (lb *LayerBase[T]) Copy() *LayerBase[T] {
	return &LayerBase[T]{
		W: lb.W.Copy(),
		B: vector.NewCopy(lb.B),
	}
}

type Layer[T vector.Real] struct {
	Base    *LayerBase[T]
	DBase   *LayerBase[T]
	Input   *matrix.Matrix[T]
	Output  *matrix.Matrix[T]
	DInputs *matrix.Matrix[T]
}

func NewLayer[T vector.Real](nNeurons, inputSize int) *Layer[T] {
	shape := []int{inputSize, nNeurons}

	return &Layer[T]{
		Base: &LayerBase[T]{
			W: matrix.NewMatrixRand[T](shape).Scale(0.01),
			B: make([]T, nNeurons),
		},
		DBase: &LayerBase[T]{},
	}
}

func NewLayerFrom[T vector.Real](tWeights *matrix.Matrix[T], biases []T) *Layer[T] {
	if len(biases) != tWeights.Shape()[1] {
		panic("Size mismatch")
	}

	return &Layer[T]{
		Base: &LayerBase[T]{
			W: tWeights.Copy(),
			B: vector.NewCopy(biases),
		},
		DBase: &LayerBase[T]{},
	}
}

func (l *Layer[T]) Forward(input *matrix.Matrix[T]) *matrix.Matrix[T] {
	l.Input = input.Copy()
	l.Output = input.Product(l.Base.W).AddVectorPerRow(l.Base.B)

	return l.Output
}

func (l *Layer[T]) Backward(dValues *matrix.Matrix[T]) *matrix.Matrix[T] {
	// Dinputs are the Gradients
	l.DInputs = dValues.Product(l.Base.W.Transpose())
	l.DBase.W = l.Input.Transpose().Product(dValues)
	l.DBase.B = make([]T, len(l.Base.B))

	for i := 0; i < len(l.Base.B); i++ {
		l.DBase.B[i] = vector.Sum(dValues.GetColumn(i))
	}

	return l.DInputs
}

func (l *Layer[T]) Copy() *Layer[T] {
	return &Layer[T]{
		Base: l.Base.Copy(),
	}
}
