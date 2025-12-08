package main

import (
	"fmt"
	"math/rand"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/single-neuron"
)

func mustWork(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	input := []float64{1.0, 2.0, 3.0, 2.5}

	// Static example
	neuron := singleneuron.NewNeuron(4)
	weights := []float64{0.2, 0.8, -0.5, 1.0}

	neuron.Bias = 2.0
	err := neuron.SetWeights(weights)
	mustWork(err)

	output, err := neuron.Synapsis(input)
	mustWork(err)

	fmt.Printf("%f \n", output)
	fmt.Println(neuron)

	// Dynamic example
	neuron2 := singleneuron.NewNeuron(4)
	neuron2.Bias = (rand.Float64() - 0.5) * 10
	neuron2.RandomizeWeights()

	output2, err := neuron2.Synapsis(input)
	mustWork(err)

	fmt.Printf("%f \n", output2)
	fmt.Println(neuron2)
}