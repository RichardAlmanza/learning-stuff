package denselayer_test

import (
	"math"
	"testing"

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

func TestLayer_CrossEntropyLoss(t *testing.T) {
	expectedAvgLoss := 1.0986123
	deltaTolerance := 2e-5

	x, y := spiral.NewSpiralData(100, 3, 0)

	dense1 := denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
	dense2 := denselayer.NewLayer(3, 3, denselayer.LinearFunction)

	outputs1 := dense1.Forward(x)
	outputs2 := dense2.Forward(outputs1).SoftMax()

	expectedLabels := matrix.NewMatrixOneHot(300, 3, vectorF64ToInt(y.Data))
	entropy := outputs2.CrossEntropyLossPerRow(expectedLabels)
	avgLoss := matrix.Avg(entropy)
	difference := math.Abs(avgLoss - expectedAvgLoss)

	t.Run("CrossEntropyLoss", func(t *testing.T) {
		if difference > deltaTolerance {
			t.Errorf("\nexpected Loss: %v \ngot: \n%v", expectedAvgLoss, avgLoss)
		}
	})
}

func TestLayer_Accuracy(t *testing.T) {
	expectedAccuracy := 0.33
	deltaTolerance := 0.1

	x, y := spiral.NewSpiralData(100, 3, 0)

	dense1 := denselayer.NewLayer(3, 2, denselayer.ReLuFunction)
	dense2 := denselayer.NewLayer(3, 3, denselayer.LinearFunction)

	outputs1 := dense1.Forward(x)
	outputs2 := dense2.Forward(outputs1).SoftMax()

	expectedLabels := matrix.NewMatrixOneHot(300, 3, vectorF64ToInt(y.Data))
	accuracy := outputs2.Accuracy(expectedLabels)

	difference := math.Abs(accuracy - expectedAccuracy)

	t.Run("CrossEntropyLoss", func(t *testing.T) {
		if difference > deltaTolerance {
			t.Errorf("\nexpected Accuracy: %v \ngot: \n%v", expectedAccuracy, accuracy)
		}
	})
}
