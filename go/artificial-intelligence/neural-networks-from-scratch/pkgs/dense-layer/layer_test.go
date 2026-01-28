package denselayer_test

import (
	"fmt"
	"math"
	"testing"

	denselayer "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/dense-layer"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/spiral"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func TestLayer_CrossEntropyLoss_Float64(t *testing.T) {
	expectedAvgLoss := 1.0986123
	deltaTolerance := 2e-5

	x, y := spiral.NewSpiralData(100, 3, 0)

	dense1 := denselayer.NewLayer(3, 2, denselayer.NewReLuFloat64())
	dense2 := denselayer.NewLayer(3, 3, denselayer.NewLinearFloat64())

	outputs1 := dense1.Forward(x)
	outputs2 := matrix.SoftMax(dense2.Forward(outputs1))

	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)
	entropy := matrix.CrossEntropyLossPerRow(outputs2, expectedLabels)
	avgLoss := vector.Avg(entropy)
	difference := math.Abs(avgLoss - expectedAvgLoss)

	t.Run("CrossEntropyLoss", func(t *testing.T) {
		if difference > deltaTolerance {
			t.Errorf("\nexpected Loss: %v \ngot: \n%v", expectedAvgLoss, avgLoss)
		}
	})
}

func TestLayer_Accuracy_Float64(t *testing.T) {
	expectedAccuracy := 0.33
	deltaTolerance := 0.1

	x, y := spiral.NewSpiralData(100, 3, 0)

	dense1 := denselayer.NewLayer(3, 2, denselayer.NewReLuFloat64())
	dense2 := denselayer.NewLayer(3, 3, denselayer.NewReLuFloat64())

	outputs1 := dense1.Forward(x)
	outputs2 := matrix.SoftMax(dense2.Forward(outputs1))

	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)
	accuracy := outputs2.Accuracy(expectedLabels)

	difference := math.Abs(accuracy - expectedAccuracy)

	t.Run("Accuracy", func(t *testing.T) {
		if difference > deltaTolerance {
			t.Errorf("\nexpected Accuracy: %v \ngot: \n%v", expectedAccuracy, accuracy)
		}
	})
}

func TestLayer_Derivative_Float64(t *testing.T) {
	expectedResult := matrix.WrapSlice([]int{3, 3},
		[]float64{
			-0.1, 1.0 / 30, 2.0 / 30,
			1.0 / 30, -5.0 / 30, 4.0 / 30,
			2.0 / 300, -1.0 / 30, 8.0 / 300,
		},
	)

	deltaTolerance := 0.1

	softmaxOutput := matrix.WrapSlice([]int{3, 3},
		[]float64{
			0.7, 0.1, 0.2,
			0.1, 0.5, 0.4,
			0.02, 0.9, 0.08,
		},
	)

	onehot := []int{0, 1, 1}
	derivative := softmaxOutput.DerivativeCrossEntropyLossSoftMaxPerRow(onehot).Scale(1 / float64(len(onehot)))

	diff := derivative.Scale(-1).Add(expectedResult).MapFunc(func(_ int,f float64) float64 { return math.Abs(f) })
	_, maxdiff := vector.Max(diff.Data)

	t.Run("Derivative", func(t *testing.T) {
		if maxdiff > deltaTolerance {
			t.Errorf("\nexpected Result: %v \ngot: \n%v", expectedResult, derivative)
		}
	})
}

func TestLayer_Copy_Float64(t *testing.T) {
	compareLayers := func(l1, l2 *denselayer.Layer[float64]) bool {
		areEqual := l1.TWeights.IsEqual(l2.TWeights) && vector.AreEqual(l1.Biases, l2.Biases)

		if !areEqual {
			return false
		}

		return fmt.Sprintf("%v", l1.ActivationFunction) == fmt.Sprintf("%v", l2.ActivationFunction)
	}

	dense := denselayer.NewLayer(10, 2, denselayer.NewSigmoidFloat64())
	denseCopy := dense.Copy()

	// verify copy
	copiedCorrectly := compareLayers(dense, denseCopy)

	// Modify original
	dense.ActivationFunction = denselayer.NewLinearFloat64()
	dense.TWeights = dense.TWeights.Scale(0.3)
	dense.Biases[2] = 1.3

	areEquals := compareLayers(dense, denseCopy)

	t.Run("Copy", func(t *testing.T) {
		if !copiedCorrectly || areEquals {
			t.Errorf("Failed to deep copy\n%v \n%v", dense, denseCopy)
		}
	})
}
