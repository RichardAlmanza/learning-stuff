package spiral

import (
	"math"
	"math/rand"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// NewSpiralData generates spiral data. see: https://cs231n.github.io/neural-networks-case-study/
// Local adaptation from https://github.com/mbdeveci/NNfSiX/blob/master/Go/p005-ReLU-Activation.go
func NewSpiralData(numberOfPoints, numberOfClasses int, seed int) (*matrix.Matrix[float64], *matrix.Matrix[int]) {
	seededRand := rand.New(rand.NewSource(int64(seed)))
	rows := numberOfPoints * numberOfClasses

	X := matrix.NewMatrix[float64]([]int{rows, 2})
	y := matrix.NewMatrix[int]([]int{rows, 1})

	for c := 0; c < numberOfClasses; c++ {
		radius := linspace(0, 1, numberOfPoints)
		t := linspace(float64(c*4), float64((c+1)*4), numberOfPoints)

		for i := range t {
			t[i] += 0.2 * seededRand.NormFloat64()
		}

		for i := 0; i < numberOfPoints; i++ {
			X.Set([]int{c*numberOfPoints + i, 0}, radius[i]*math.Sin(t[i]*2.5))
			X.Set([]int{c*numberOfPoints + i, 1}, radius[i]*math.Cos(t[i]*2.5))
			y.Set([]int{c*numberOfPoints + i, 0}, c)
		}
	}

	return X, y
}

func linspace(start, end float64, num int) []float64 {
	result := make([]float64, num)
	step := (end - start) / float64(num-1)

	for i := range result {
		result[i] = start + float64(i)*step
	}

	return result
}

// Local adaptation from https://github.com/mbdeveci/NNfSiX/blob/master/Go/p005-ReLU-Activation.go
func Plot(X *matrix.Matrix[float64], y *matrix.Matrix[int]) {
	p := plot.New()

	p.Title.Text = "Spiral"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	c1 := make(plotter.XYs, 0)
	c2 := make(plotter.XYs, 0)
	c3 := make(plotter.XYs, 0)

	for i := 0; i < X.Shape()[0]; i++ {
		coordX := []int{i, 0}
		coordY := []int{i, 1}

		if y.GetAt(coordX) == 0 {
			c1 = append(c1, plotter.XY{
				X: X.GetAt(coordX),
				Y: X.GetAt(coordY),
			})
		} else if y.GetAt(coordX) == 1 {
			c2 = append(c2, plotter.XY{
				X: X.GetAt(coordX),
				Y: X.GetAt(coordY),
			})
		} else if y.GetAt(coordX) == 2 {
			c3 = append(c3, plotter.XY{
				X: X.GetAt(coordX),
				Y: X.GetAt(coordY),
			})
		}
	}

	_ = plotutil.AddLinePoints(p, "0", c1, "1", c2, "2", c3)

	if err := p.Save(8*vg.Centimeter, 8*vg.Centimeter, "points.png"); err != nil {
		panic(err)
	}
}
