package denselayer_test

import (
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

	dense1 := denselayer.NewLayer[float64](3, 2)
	dense2 := denselayer.NewLayer[float64](3, 3)

	relu := denselayer.NewReLu[float64]()
	softMax := denselayer.NewSoftMax[float64]()

	dense1.Forward(x)
	relu.Forward(dense1.IOStates.Output)
	dense2.Forward(relu.Output)
	softMax.Forward(dense2.IOStates.Output)

	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)
	entropy := matrix.CrossEntropyLossPerRow(softMax.Output, expectedLabels)
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

	dense1 := denselayer.NewLayer[float64](3, 2)
	dense2 := denselayer.NewLayer[float64](3, 3)

	relu := denselayer.NewReLu[float64]()
	softMax := denselayer.NewSoftMax[float64]()

	dense1.Forward(x)
	relu.Forward(dense1.IOStates.Output)
	dense2.Forward(relu.Output)
	softMax.Forward(dense2.IOStates.Output)

	expectedLabels := matrix.NewMatrixOneHot[float64]([]int{300, 3}, y.Data)
	accuracy := softMax.Output.Accuracy(expectedLabels)

	difference := math.Abs(accuracy - expectedAccuracy)

	t.Run("Accuracy", func(t *testing.T) {
		if difference > deltaTolerance {
			t.Errorf("\nexpected Accuracy: %v \ngot: \n%v", expectedAccuracy, accuracy)
		}
	})
}

func TestLayer_Derivative_Float64(t *testing.T) {
	deltaTolerance := 0.1
	shape := []int{3,3}

	expectedResult := matrix.WrapSlice(shape,
		[]float64{
			-0.1, 1.0 / 30, 2.0 / 30,
			1.0 / 30, -5.0 / 30, 4.0 / 30,
			2.0 / 300, -1.0 / 30, 8.0 / 300,
		},
	)


	softmaxOutput := matrix.WrapSlice(shape,
		[]float64{
			0.7, 0.1, 0.2,
			0.1, 0.5, 0.4,
			0.02, 0.9, 0.08,
		},
	)

	targets := matrix.NewMatrixOneHot[float64](shape, []int{0, 1, 1})
	derivative := matrix.DerivativeCrossEntropyLossSoftMaxPerRow(softmaxOutput, targets).Scale(1 / float64(shape[0]))

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
		return l1.Base.W.IsEqual(l2.Base.W) && vector.AreEqual(l1.Base.B, l2.Base.B)
	}

	dense := denselayer.NewLayer[float64](10, 2)
	denseCopy := dense.Copy()

	// verify copy
	copiedCorrectly := compareLayers(dense, denseCopy)

	// Modify original
	dense.Base.W = dense.Base.W.Scale(0.3)
	dense.Base.B[2] = 1.3

	areEquals := compareLayers(dense, denseCopy)

	t.Run("Copy", func(t *testing.T) {
		if !copiedCorrectly || areEquals {
			t.Errorf("Failed to deep copy\n%v \n%v", dense, denseCopy)
		}
	})
}
