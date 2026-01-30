package main

import (
	"fmt"
	// "math"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	neuralnetwork "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/neural-network"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)
	x32 := matrix.NewMatrixFromSlice(x.Shape(), vector.ConvertTo[[]float64, []float64](x.Data))
	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)

	// original dynamic example

	dense1 := denselayer.NewLayer[float64](64, 2)
	dense2 := denselayer.NewLayer[float64](3, 64)

	relu := denselayer.NewReLu[float64]()
	softmax := denselayer.NewSoftMaxWithCrossEntropy[float64]()

	optimizer := neuralnetwork.NewSGD[float64](3, 1e-4)

	for i := 0; i < 10001; i++ {

		dense1.Forward(x32)
		relu.Forward(dense1.Output)
		dense2.Forward(relu.Output)
		softmax.Forward(dense2.Output)

		loss := softmax.Loss(expectedLabels)
		accuracy := softmax.Output.Accuracy(expectedLabels)

		if i%500 == 0 {
			fmt.Print("Epoch: ", i)
			fmt.Printf(", Current Learning Rate: %.3f", optimizer.LearningRate)
			fmt.Printf(", Accuracy: %.3f", accuracy)
			fmt.Printf(", Loss: %0.5f\n", loss)
		}

		// Backpropagation
		softmax.Backward(softmax.Output, expectedLabels)
		dense2.Backward(softmax.DInput)
		relu.Backward(dense2.DInputs)
		dense1.Backward(relu.DInput)

		// Update parameters using the optimizer
		optimizer.UpdateLearningRate(i)
		optimizer.UpdateParameters(dense1)
		optimizer.UpdateParameters(dense2)
	}
}
