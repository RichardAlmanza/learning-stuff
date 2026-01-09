package main

import (
	"fmt"
	"math"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	layerbasedneurons "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
)

func vectorF64ToInt(vectorF []float64) []int {
	vectorInt := make([]int, len(vectorF))

	for i := 0; i < len(vectorF); i++ {
		vectorInt[i] = int(vectorF[i])
	}

	return vectorInt
}

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)

	// original dynamic example
	var dense1 *denselayer.Layer
	var dense2 *denselayer.Layer

	var lowestLoss float64 = math.MaxFloat64
	var bestDense1 *denselayer.Layer
	var bestDense2 *denselayer.Layer

	for i := 0; i < 1000; i++ {
		dense1 = denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
		dense2 = denselayer.NewLayer(3, 3, denselayer.LinearFunction)

		outputs1 := dense1.Forward(x)
		outputs2 := dense2.Forward(outputs1).SoftMax()

		expectedLabels := matrix.NewMatrixOneHot(300, 3, vectorF64ToInt(y.Data))

		loss := matrix.Avg(outputs2.CrossEntropyLossPerRow(expectedLabels))
		accuracy := outputs2.Accuracy(expectedLabels)

		if loss < lowestLoss {
			fmt.Printf("New set of weights found!\nIteration: %d\nLoss: %v\nAccuracy: %v\n", i, loss, accuracy)

			bestDense1 = dense1.Copy()
			bestDense2 = dense2.Copy()
			lowestLoss = loss
		}
	}

	fmt.Println(bestDense1, "\n", bestDense2)

	// Dynamic example

	var layer1 *layerbasedneurons.LayerBasedNeurons
	var layer2 *layerbasedneurons.LayerBasedNeurons

	lowestLoss = math.MaxFloat64
	var bestLayer1 *layerbasedneurons.LayerBasedNeurons
	var bestLayer2 *layerbasedneurons.LayerBasedNeurons

	for i := 0; i < 1000; i++ {

		layer1 = layerbasedneurons.NewLayerBasedNeurons(3, 2, denselayer.ReLuFunction)
		layer2 = layerbasedneurons.NewLayerBasedNeurons(3, 3, denselayer.LinearFunction)

		layer1.RandomizeWeights()
		layer2.RandomizeWeights()

		layer1outputs := layer1.ForwardBatch(x)
		layer2outputs := layer2.ForwardBatch(layer1outputs).SoftMax()

		expectedLabels := matrix.NewMatrixOneHot(300, 3, vectorF64ToInt(y.Data))

		loss := matrix.Avg(layer2outputs.CrossEntropyLossPerRow(expectedLabels))
		accuracy := layer2outputs.Accuracy(expectedLabels)

		if loss < lowestLoss {
			fmt.Printf("New set of weights found!\nIteration: %d\nLoss: %v\nAccuracy: %v\n", i, loss, accuracy)

			bestLayer1 = layer1.Copy()
			bestLayer2 = layer2.Copy()
			lowestLoss = loss
		}
	}

	fmt.Println(bestLayer1, "\n", bestLayer2)

}
