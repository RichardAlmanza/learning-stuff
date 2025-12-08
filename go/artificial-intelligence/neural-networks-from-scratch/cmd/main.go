package main

import (
	"fmt"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/single-neuron"
	"math/rand"
)

func mustWork(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func mustSynapsis(n *singleneuron.NeuronFloat64, input []float64) float64 {
	output, err := n.Synapsis(input)
	mustWork(err)

	return output
}

func main() {
	input := []float64{1.0, 2.0, 3.0, 2.5}

	weights1 := []float64{0.2, 0.8, -0.5, 1.0}
	weights2 := []float64{0.5, -0.91, 0.26, -0.5}
	weights3 := []float64{-0.26, -0.27, 0.17, 0.87}

	bias1 := 2.0
	bias2 := 3.0
	bias3 := 0.5

	// original static example

	outputs1 := []float64{
		// neuron 1
		input[0] * weights1[0] +
		input[1] * weights1[1] +
		input[2] * weights1[2] +
		input[3] * weights1[3] +
		bias1,

		// neuron 2
		input[0] * weights2[0] +
		input[1] * weights2[1] +
		input[2] * weights2[2] +
		input[3] * weights2[3] +
		bias2,

		// neuron 3
		input[0] * weights3[0] +
		input[1] * weights3[1] +
		input[2] * weights3[2] +
		input[3] * weights3[3] +
		bias3,
	}

	fmt.Println(outputs1)

	// Static example
	neuron1 := singleneuron.NewNeuron(4)
	neuron2 := singleneuron.NewNeuron(4)
	neuron3 := singleneuron.NewNeuron(4)

	var err error

	neuron1.Bias = bias1
	err = neuron1.SetWeights(weights1)
	mustWork(err)

	neuron2.Bias = bias2
	err = neuron2.SetWeights(weights2)
	mustWork(err)

	neuron3.Bias = bias3
	err = neuron3.SetWeights(weights3)
	mustWork(err)

	outputs2 := []float64{
		mustSynapsis(neuron1, input),
		mustSynapsis(neuron2, input),
		mustSynapsis(neuron3, input),
	}

	fmt.Println(outputs2)
	fmt.Println(neuron1)
	fmt.Println(neuron2)
	fmt.Println(neuron3)

	// Dynamic example

	sizeLayer := 3
	layer := make([]singleneuron.NeuronFloat64, sizeLayer)
	outputs3 := make([]float64, sizeLayer)

	for i := 0; i < sizeLayer; i++ {
		neuron := singleneuron.NewNeuron(4)
		neuron.RandomizeWeights()
		neuron.Bias = (rand.Float64() - 0.5) * 10

		layer[i] = *neuron
		outputs3[i] = mustSynapsis(neuron, input)
	}

	fmt.Println(outputs3)
	fmt.Println(layer)
}
