package main

import (
	"fmt"
	"math"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
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
	expectedLabels := matrix.NewMatrixOneHot(300, 3, vectorF64ToInt(y.Data))

	// original dynamic example
	var lowestLoss float64 = math.MaxFloat64
	bestDense1 := denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
	bestDense2 := denselayer.NewLayer(3, 3, denselayer.LinearFunction)

	var dense2 *denselayer.Layer
	var dense1 *denselayer.Layer

	auxWeights1 := matrix.NewMatrix(bestDense1.Weights.Rows, bestDense1.Weights.Columns)
	auxBiases1 := matrix.NewMatrix(1, len(bestDense1.Biases))
	auxWeights2 := matrix.NewMatrix(bestDense2.Weights.Rows, bestDense2.Weights.Columns)
	auxBiases2 := matrix.NewMatrix(1, len(bestDense2.Biases))

	for i := 0; i < 1000; i++ {
		dense1 = bestDense1.Copy()
		dense2 = bestDense2.Copy()

		auxWeights1.Randomize()
		auxBiases1.Randomize()
		auxWeights2.Randomize()
		auxBiases2.Randomize()

		dense1.Weights = dense1.Weights.Add(auxWeights1.Scale(0.05))
		dense1.Biases = matrix.SumVectors(dense1.Biases, auxBiases1.Scale(0.05).Data)
		dense2.Weights = dense2.Weights.Add(auxWeights2.Scale(0.05))
		dense2.Biases = matrix.SumVectors(dense2.Biases, auxBiases2.Scale(0.05).Data)

		outputs1 := dense1.Forward(x)
		outputs2 := dense2.Forward(outputs1).SoftMax()

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
}
