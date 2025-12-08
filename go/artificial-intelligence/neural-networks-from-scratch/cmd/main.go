package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type NeuronFloat64 struct {
	InputSize int
	Bias float64
	Weights []float64
}

func NewNeuron(numInputs int) *NeuronFloat64 {
	return &NeuronFloat64{
		InputSize: numInputs,
		Bias: 0,
		Weights: make([]float64, numInputs),
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

	return  output, nil
}

func mustWork(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	input := []float64{1.0, 2.0, 3.0, 2.5}

	// Static example
	neuron := NewNeuron(4)
	weights := []float64{0.2, 0.8, -0.5, 1.0}

	neuron.Bias = 2.0
	err := neuron.SetWeights(weights)
	mustWork(err)

	output, err := neuron.Synapsis(input)
	mustWork(err)

	fmt.Printf("%f \n", output)
	fmt.Println(neuron)

	// Dynamic example
	neuron2 := NewNeuron(4)
	neuron2.Bias = (rand.Float64() - 0.5) * 10
	neuron2.RandomizeWeights()

	output2, err := neuron2.Synapsis(input)
	mustWork(err)

	fmt.Printf("%f \n", output2)
	fmt.Println(neuron2)
}