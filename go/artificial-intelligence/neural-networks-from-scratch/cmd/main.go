package main

import (
	"fmt"
	"math"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)
	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)

	// original dynamic example
	var lowestLoss float64 = math.MaxFloat64
	bestDense1 := denselayer.NewLayer(3, 2, denselayer.NewReLuFloat64())
	bestDense2 := denselayer.NewLayer(3, 3, denselayer.NewLinearFloat64())

	var dense2 *denselayer.Layer[float64]
	var dense1 *denselayer.Layer[float64]

	auxWeights1 := matrix.NewMatrix[float64](bestDense1.TWeights.Shape())
	auxBiases1 := matrix.NewMatrix[float64]([]int{1, len(bestDense1.Biases)})
	auxWeights2 := matrix.NewMatrix[float64](bestDense2.TWeights.Shape())
	auxBiases2 := matrix.NewMatrix[float64]([]int{1, len(bestDense2.Biases)})

	for i := 0; i < 1000; i++ {
		dense1 = bestDense1.Copy()
		dense2 = bestDense2.Copy()

		auxWeights1.Randomize()
		auxBiases1.Randomize()
		auxWeights2.Randomize()
		auxBiases2.Randomize()

		dense1.TWeights = dense1.TWeights.Add(auxWeights1.Scale(0.05))
		dense1.Biases = vector.AddVectors(dense1.Biases, auxBiases1.Scale(0.05).Data)
		dense2.TWeights = dense2.TWeights.Add(auxWeights2.Scale(0.05))
		dense2.Biases = vector.AddVectors(dense2.Biases, auxBiases2.Scale(0.05).Data)

		outputs1 := dense1.Forward(x)
		outputs2 := matrix.SoftMax(dense2.Forward(outputs1))

		loss := vector.Avg(matrix.CrossEntropyLossPerRow(outputs2, expectedLabels))
		accuracy := outputs2.Accuracy(expectedLabels)

		if loss < lowestLoss {
			fmt.Printf("New set of weights found!\nIteration: %d\nLoss: %v\nAccuracy: %v\n", i, loss, accuracy)

			bestDense1 = dense1.Copy()
			bestDense2 = dense2.Copy()
			lowestLoss = loss
		}
	}

	fmt.Println(bestDense1, "\n", bestDense2)
}
