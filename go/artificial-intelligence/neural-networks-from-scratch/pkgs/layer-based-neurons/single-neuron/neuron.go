package singleneuron

import (
	"errors"
	"math/rand"
)

type NeuronFloat64 struct {
	InputSize int
	Bias      float64
	Weights   []float64
}

func NewNeuron(numInputs int) *NeuronFloat64 {
	return &NeuronFloat64{
		InputSize: numInputs,
		Bias:      0,
		Weights:   make([]float64, numInputs),
	}
}

func (n *NeuronFloat64) SetWeights(newWeights []float64) error {
	if len(newWeights) != n.InputSize {
		return errors.New("size mismatch")
	}

	n.Weights = newWeights

	return nil
}

func (n *NeuronFloat64) RandomizeWeights() {
	for i := 0; i < n.InputSize; i++ {
		n.Weights[i] = (rand.Float64() - 0.5) * 10
	}
}

func (n *NeuronFloat64) Synapsis(input []float64) (float64, error) {
	if len(input) != n.InputSize {
		return 0, errors.New("size mismatch")
	}

	var output float64 = 0

	for i := 0; i < len(input); i++ {
		output += n.Weights[i] * input[i]
	}

	output += n.Bias

	return output, nil
}
