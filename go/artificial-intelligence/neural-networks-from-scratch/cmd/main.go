package main

import (
	"fmt"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	layerbasedneurons "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
)

func main() {
	x, _ := spiral.NewSpiralData(100, 3, 0)

	// original dynamic example

	dense1 := denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
	dense2 := denselayer.NewLayer(3, 3, denselayer.LinearFunction)

	outputs1 := dense1.Forward(x)
	outputs2 := dense2.Forward(outputs1).SoftMax()

	outputs2.Rows = 5 // This will limit the output to only show the first 5 rows
	fmt.Println(outputs2)

	// Dynamic example

	layer1 := layerbasedneurons.NewLayerBasedNeurons(3, 2, denselayer.ReLuFunction)
	layer1.RandomizeWeights()
	layer2 := layerbasedneurons.NewLayerBasedNeurons(3, 3, denselayer.LinearFunction)
	layer2.RandomizeWeights()

	layer1outputs := layer1.ForwardBatch(x)
	layer2outputs := layer2.ForwardBatch(layer1outputs).SoftMax()

	layer2outputs.Rows = 5
	fmt.Println(layer2outputs)
}
