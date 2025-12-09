package main

import (
	"fmt"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	layerbasedneurons "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
)

func main() {
	outputs := matrix.NewMatrixFromSlice(1, 3, []float64{4.8, 1.21, 2.385})
	fmt.Println(outputs)
	fmt.Println(outputs.SoftMax())

	x, _ := spiral.NewSpiralData(100, 3, 0)

	// original dynamic example

	dense1 := denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
	outputs1 := dense1.Forward(x)

	outputs1.Rows = 5 // This will limit the output to only show the first 5 rows
	fmt.Println(outputs1)

	// Dynamic example

	layer1 := layerbasedneurons.NewLayerBasedNeurons(3, 2, denselayer.ReLuFunction)
	layer1.RandomizeWeights()

	outputs2 := layer1.ForwardBatch(x)
	outputs2.Rows = 5
	fmt.Println(outputs2)
}
