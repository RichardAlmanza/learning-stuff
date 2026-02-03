package denselayer

import (
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

// PeripheralStates represents the peripheral states of a layer.

type PeripheralStates[T vector.Real] struct {
	Input    *matrix.Matrix[T]
	Output   *matrix.Matrix[T]
	Gradient *matrix.Matrix[T]
}

func (ps *PeripheralStates[T]) Copy() *PeripheralStates[T] {
	newPS := &PeripheralStates[T]{}

	if ps.Input != nil {
		newPS.Input = ps.Input.Copy()
	}
	if ps.Output != nil {
		newPS.Output = ps.Output.Copy()
	}
	if ps.Gradient != nil {
		newPS.Gradient = ps.Gradient.Copy()
	}

	return newPS
}

// LayerBase represents the base properties of a layer.
type LayerBase[T vector.Real] struct {
	// Weights and biases
	W *matrix.Matrix[T]
	B []T
}

// Copy creates a deep copy of the LayerBase.
func (lb *LayerBase[T]) Copy() *LayerBase[T] {
	newLB := &LayerBase[T]{}

	if lb.W != nil {
		newLB.W = lb.W.Copy()
	}
	if lb.B != nil {
		newLB.B = vector.NewCopy(lb.B)
	}

	return newLB
}

// Layer represents a dense layer in the neural network.
type Layer[T vector.Real] struct {
	Base     *LayerBase[T]
	DBase    *LayerBase[T]
	Momentum *LayerBase[T]
	Cache    *LayerBase[T]
	IOStates *PeripheralStates[T]
}

// NewLayer creates a new layer with the specified number of neurons and input size.
func NewLayer[T vector.Real](nNeurons, inputSize int) *Layer[T] {
	shape := []int{inputSize, nNeurons}

	return &Layer[T]{
		Base: &LayerBase[T]{
			W: matrix.NewMatrixRand[T](shape).Scale(0.01),
			B: make([]T, nNeurons),
		},
		DBase:    &LayerBase[T]{},
		Momentum: &LayerBase[T]{},
		Cache:    &LayerBase[T]{},
		IOStates: &PeripheralStates[T]{},
	}
}

// NewLayerFrom creates a new layer with the specified weights and biases.
func NewLayerFrom[T vector.Real](tWeights *matrix.Matrix[T], biases []T) *Layer[T] {
	if len(biases) != tWeights.Shape()[1] {
		panic("Size mismatch")
	}

	return &Layer[T]{
		Base: &LayerBase[T]{
			W: tWeights.Copy(),
			B: vector.NewCopy(biases),
		},
		DBase:    &LayerBase[T]{},
		Momentum: &LayerBase[T]{},
		Cache:    &LayerBase[T]{},
		IOStates: &PeripheralStates[T]{},
	}
}

func (l *Layer[T]) Forward(input *matrix.Matrix[T]) *matrix.Matrix[T] {
	l.IOStates.Input = input.Copy()
	l.IOStates.Output = input.Product(l.Base.W).AddVectorPerRow(l.Base.B)

	return l.IOStates.Output
}

func (l *Layer[T]) Backward(dValues *matrix.Matrix[T]) *matrix.Matrix[T] {
	// Dinputs are the Gradients
	l.IOStates.Gradient = dValues.Product(l.Base.W.Transpose())
	l.DBase.W = l.IOStates.Input.Transpose().Product(dValues)
	l.DBase.B = make([]T, len(l.Base.B))

	for i := 0; i < len(l.Base.B); i++ {
		l.DBase.B[i] = vector.Sum(dValues.GetColumn(i))
	}

	return l.IOStates.Gradient
}

// Copy creates a deep copy of the layer.
func (l *Layer[T]) Copy() *Layer[T] {
	return &Layer[T]{
		Base:     l.Base.Copy(),
		DBase:    l.DBase.Copy(),
		Momentum: l.Momentum.Copy(),
		Cache:    l.Cache.Copy(),
		IOStates: l.IOStates.Copy(),
	}
}
