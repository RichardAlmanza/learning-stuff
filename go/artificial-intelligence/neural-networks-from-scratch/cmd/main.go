package main

import (
	"fmt"
	// "math"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	// neuralnetwork "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/neural-network"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)
	x32 := matrix.NewMatrixFromSlice(x.Shape(), vector.ConvertTo[[]float64, []float32](x.Data))
	expectedLabels := matrix.NewMatrixOneHot[float32]([]int{300, 3}, y.Data)

	// original dynamic example

	dense1 := denselayer.NewLayer[float32](64, 2)
	dense2 := denselayer.NewLayer[float32](3, 64)

	relu := denselayer.NewReLu[float32]()
	softmax := denselayer.NewSoftMaxWithCrossEntropy[float32]()

	dense1.Forward(x32)
	relu.Forward(dense1.Output)
	dense2.Forward(dense1.Output)
	softmax.Forward(dense2.Output)

	loss := vector.Avg(matrix.CrossEntropyLossPerRow(softmax.Output, expectedLabels))
	accuracy := softmax.Output.Accuracy(expectedLabels)

	fmt.Println("Loss: ", loss)
	fmt.Println("Accuracy: ", accuracy)

	softmax.Backward(softmax.Output, y.Data)
	dense2.Backward(expectedLabels)
	relu.Backward(dense2.DInputs)
	dense1.Backward(relu.DInput)

	fmt.Println(dense1.DWeights)
	fmt.Println(dense1.DBiases)
	fmt.Println(dense2.DWeights)
	fmt.Println(dense2.DBiases)
}
