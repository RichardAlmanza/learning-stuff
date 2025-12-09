package main

import (
	"fmt"

	layerbasedneurons "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
)

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)

	spiral.Plot(x,y)

	input := &matrix.MatrixFloat64{
		Rows:    3,
		Columns: 4,
		Data: []float64{
			1.0, 2.0, 3.0, 2.5,
			2.0, 5.0, -1.0, 2.0,
			-1.5, 2.7, 3.3, -0.8,
		},
	}

	layer1Weights := &matrix.MatrixFloat64{
		Rows:    3,
		Columns: 4,
		Data: []float64{
			0.2, 0.8, -0.5, 1.0,
			0.5, -0.91, 0.26, -0.5,
			-0.26, -0.27, 0.17, 0.87,
		},
	}

	layer1Biases := []float64{2.0, 3.0, 0.5}

	layer2Weights := &matrix.MatrixFloat64{
		Rows:    3,
		Columns: 3,
		Data: []float64{
			0.1, -0.14, 0.5,
			-0.5, 0.12, -0.33,
			-0.44, 0.73, -0.13,
		},
	}

	layer2Biases := []float64{-1, 2, -0.5}

	// original static example

	layer1Outputs := input.Product(layer1Weights.Transpose()).AddVectorPerRow(layer1Biases)
	layer2Outputs := layer1Outputs.Product(layer2Weights.Transpose()).AddVectorPerRow(layer2Biases)

	fmt.Println(layer2Outputs)

	// Static example
	layer1 := layerbasedneurons.NewLayerBasedNeurons(3, 4)
	layer1.SetWeights(*layer1Weights)
	layer1.SetBiases(layer1Biases)

	layer2 := layerbasedneurons.NewLayerBasedNeurons(3, 3)
	layer2.SetWeights(*layer2Weights)
	layer2.SetBiases(layer2Biases)

	layer1Outputs2 := layer1.ForwardBatch(input)
	layer2Outputs2 := layer2.ForwardBatch(layer1Outputs2)

	fmt.Println(layer2Outputs2)

	// Dynamic example

	sizeLayer1 := 3
	randLayer1 := layerbasedneurons.NewLayerBasedNeurons(sizeLayer1, input.Columns)
	randLayer1.RandomizeWeights()
	randLayer1.RandomizeBiases()

	sizeLayer2 := 3
	randLayer2 := layerbasedneurons.NewLayerBasedNeurons(sizeLayer2, randLayer1.Size)
	randLayer2.RandomizeWeights()
	randLayer2.RandomizeBiases()

	randLayer1Outputs := randLayer1.ForwardBatch(input)
	randLayer2Outputs := randLayer2.ForwardBatch(randLayer1Outputs)

	fmt.Println(randLayer2Outputs)
	fmt.Println(randLayer2)
}
