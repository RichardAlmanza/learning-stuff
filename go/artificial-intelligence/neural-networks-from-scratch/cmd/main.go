package main

import (
	"fmt"

	layerbasedneurons "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

func main() {
	input := []float64{1.0, 2.0, 3.0, 2.5}

	weights := &matrix.MatrixFloat64{
		Rows:    3,
		Columns: 4,
		Data: []float64{
			0.2, 0.8, -0.5, 1.0,
			0.5, -0.91, 0.26, -0.5,
			-0.26, -0.27, 0.17, 0.87,
		},
	}

	biases := []float64{2.0, 3.0, 0.5}

	// original static example

	outputs1 := make([]float64, 3)

	for indexNeuron := 0; indexNeuron < 3; indexNeuron++ {
		output := 0.0

		for indexInput := 0; indexInput < 4; indexInput++ {
			output += input[indexInput] * weights.GetAt(indexInput, indexNeuron)
		}

		outputs1[indexNeuron] = output + biases[indexNeuron]
	}

	fmt.Println(outputs1)

	// Static example
	layer1 := layerbasedneurons.NewLayerBasedNeurons(3, 4)
	layer1.SetWeights(*weights)
	layer1.SetBiases(biases)

	outputs2 := layer1.Forward(input)

	fmt.Println(outputs2)
	fmt.Println(layer1)

	// Dynamic example

	sizeLayer := 3
	layer2 := layerbasedneurons.NewLayerBasedNeurons(sizeLayer, len(input))
	layer2.RandomizeWeights()
	layer2.RandomizeBiases()

	outputs3 := layer2.Forward(input)

	fmt.Println(outputs3)
	fmt.Println(layer2)
}
