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
func NewSpiralData(numberOfPoints, numberOfClasses int, seed int) (*matrix.MatrixFloat64, *matrix.MatrixFloat64) {
	seededRand := rand.New(rand.NewSource(int64(seed)))

	X := matrix.NewMatrix(numberOfPoints*numberOfClasses, 2)
	y := matrix.NewMatrix(numberOfPoints*numberOfClasses, 1)

	for c := 0; c < numberOfClasses; c++ {
		radius := linspace(0, 1, numberOfPoints)
		t := linspace(float64(c*4), float64((c+1)*4), numberOfPoints)

		for i := range t {
			t[i] += 0.2 * seededRand.NormFloat64()
		}

		for i := 0; i < numberOfPoints; i++ {
			X.Set(c*numberOfPoints+i, 0, radius[i]*math.Sin(t[i]*2.5))
			X.Set(c*numberOfPoints+i, 1, radius[i]*math.Cos(t[i]*2.5))
			y.Set(c*numberOfPoints+i, 0, float64(c))
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
func Plot(X *matrix.MatrixFloat64, y *matrix.MatrixFloat64) {
	p := plot.New()

	p.Title.Text = "Spiral"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	c1 := make(plotter.XYs, 0)
	c2 := make(plotter.XYs, 0)
	c3 := make(plotter.XYs, 0)

	for i := 0; i < X.Rows; i++ {
		if y.GetAt(i, 0) == 0 {
			c1 = append(c1, plotter.XY{
				X: X.GetAt(i, 0),
				Y: X.GetAt(i, 1),
			})
		} else if y.GetAt(i, 0) == 1 {
			c2 = append(c2, plotter.XY{
				X: X.GetAt(i, 0),
				Y: X.GetAt(i, 1),
			})
		} else if y.GetAt(i, 0) == 2 {
			c3 = append(c3, plotter.XY{
				X: X.GetAt(i, 0),
				Y: X.GetAt(i, 1),
			})
		}
	}

	_ = plotutil.AddLinePoints(p, "0", c1, "1", c2, "2", c3)

	if err := p.Save(8*vg.Centimeter, 8*vg.Centimeter, "points.png"); err != nil {
		panic(err)
	}
}
