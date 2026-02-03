package main

import (
	"fmt"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	neuralnetwork "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/neural-network"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func main() {
	x, y := spiral.NewSpiralData(100, 3, 0)
	xTest, _ := spiral.NewSpiralData(500, 3, 156)
	x32 := matrix.NewMatrixFromSlice(x.Shape(), vector.ConvertTo[[]float64, []float64](x.Data)).Copy()
	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data).Copy()
	// expectedLabelsTest := matrix.NewMatrixOneHot[float64]([]int{1500, 3}, yTest.Data)

	// original dynamic example

	dense1 := denselayer.NewLayer[float64](64, 2)
	dense2 := denselayer.NewLayer[float64](3, 64)

	relu := denselayer.NewReLu[float64]()
	softmax := denselayer.NewSoftMaxWithCrossEntropy[float64]()

	optimizer := neuralnetwork.NewAdam[float64](5e-2, 5e-7, 1e-7, 0.9, 0.999)

	// fmt.Println(optimizer)

	for i := 0; i < 10001; i++ {

		dense1.Forward(x32)
		relu.Forward(dense1.IOStates.Output)
		dense2.Forward(relu.Output)
		softmax.Forward(dense2.IOStates.Output)

		loss := softmax.Loss(expectedLabels)
		accuracy := softmax.Output.Accuracy(expectedLabels)

		if i%500 == 0 {
			fmt.Print("Epoch: ", i)
			fmt.Printf(", Accuracy: %.3f", accuracy)
			fmt.Printf(", Loss: %0.5f", loss)
			fmt.Printf(", Current Learning Rate: %.6f\n", optimizer.LearningRate)
		}

		// Backpropagation
		softmax.Backward(softmax.Output, expectedLabels)
		dense2.Backward(softmax.Gradient)
		relu.Backward(dense2.IOStates.Gradient)
		dense1.Backward(relu.Gradient)

		// Update parameters using the optimizer
		optimizer.UpdateLearningRate(i)
		optimizer.UpdateParameters(dense1, i)
		optimizer.UpdateParameters(dense2, i)
	}

	// Test Model

	dense1.Forward(xTest)
	relu.Forward(dense1.IOStates.Output)
	dense2.Forward(relu.Output)
	softmax.Forward(dense2.IOStates.Output)

	// loss := softmax.Loss(expectedLabelsTest)
	// accuracy := softmax.Output.Accuracy(expectedLabelsTest)

// 	fmt.Printf("Test Accuracy: %.5f", accuracy)
// 	fmt.Println()
// 	fmt.Printf("Test Loss: %.5f", loss)
// 	fmt.Println()
}
