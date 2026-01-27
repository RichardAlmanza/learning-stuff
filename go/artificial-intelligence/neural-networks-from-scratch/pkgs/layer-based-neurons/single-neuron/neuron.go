package singleneuron

import (
	"errors"
	"math/rand"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type Neuron[T vector.Real] struct {
	InputSize          int
	Bias               T
	Weights            []T
	ActivationFunction func(T) T
}

func NewNeuron[T vector.Real](numInputs int, function func(T) T) *Neuron[T] {
	return &Neuron[T]{
		InputSize:          numInputs,
		Bias:               0,
		Weights:            make([]T, numInputs),
		ActivationFunction: function,
	}
}

func (n *Neuron[T]) SetWeights(newWeights []T) error {
	if len(newWeights) != n.InputSize {
		return errors.New("size mismatch")
	}

	n.Weights = newWeights

	return nil
}

func (n *Neuron[T]) RandomizeWeights() {
	for i := 0; i < n.InputSize; i++ {
		n.Weights[i] = T((rand.Float64() - 0.5) * 0.01)
	}
}

func (n *Neuron[T]) Synapsis(input []T) (T, error) {
	if len(input) != n.InputSize {
		return 0, errors.New("size mismatch")
	}

	var output T = 0

	for i := 0; i < len(input); i++ {
		output += n.Weights[i] * input[i]
	}

	output += n.Bias

	return n.ActivationFunction(output), nil
}

func (n *Neuron[T]) Copy() *Neuron[T] {

	return &Neuron[T]{
		InputSize:          n.InputSize,
		Bias:               n.Bias,
		Weights:            vector.NewCopy(n.Weights),
		ActivationFunction: n.ActivationFunction,
	}
}
