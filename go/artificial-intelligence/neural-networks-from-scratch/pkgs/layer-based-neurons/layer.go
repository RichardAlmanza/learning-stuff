package layerbasedneurons

import (
	"math/rand"

	singleneuron "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons/single-neuron"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
)

type LayerBasedNeurons struct {
	Size      int
	InputSize int
	Neurons   []singleneuron.NeuronFloat64
}

func NewLayerBasedNeurons(size, inputSize int) *LayerBasedNeurons {
	if size < 1 || inputSize < 1 {
		panic("WTF! an empty or negative layer/input size?")
	}

	layer := &LayerBasedNeurons{
		Size:      size,
		InputSize: inputSize,
		Neurons:   make([]singleneuron.NeuronFloat64, size),
	}

	for i := 0; i < size; i++ {
		layer.Neurons[i] = *singleneuron.NewNeuron(inputSize)
	}

	return layer
}

func (l *LayerBasedNeurons) SetWeights(m matrix.MatrixFloat64) {
	if m.Columns != l.InputSize || m.Rows != l.Size {
		panic("Dimensions mismatch!")
	}

	for i := 0; i < l.Size; i++ {
		l.Neurons[i].SetWeights(m.GetRow(i))
	}
}

func (l *LayerBasedNeurons) SetBiases(biases []float64) {
	if len(biases) != l.Size {
		panic("Size mismatch!")
	}

	for i := 0; i < l.Size; i++ {
		l.Neurons[i].Bias = biases[i]
	}
}

func (l *LayerBasedNeurons) RandomizeWeights() {
	for i := 0; i < l.Size; i++ {
		l.Neurons[i].RandomizeWeights()
	}
}

func (l *LayerBasedNeurons) RandomizeBiases() {
	for i := 0; i < l.Size; i++ {
		l.Neurons[i].Bias = (rand.Float64() - 0.5) * 10
	}
}

func (l *LayerBasedNeurons) Forward(input []float64) []float64 {
	output := make([]float64, l.Size)

	for i := 0; i < l.Size; i++ {
		out, err := l.Neurons[i].Synapsis(input)
		if err != nil {
			panic(err.Error())
		}

		output[i] = out
	}

	return output
}

func (l *LayerBasedNeurons) ForwardBatch(input *matrix.MatrixFloat64) *matrix.MatrixFloat64 {
	output := matrix.NewMatrix(input.Rows, l.Size)

	for sample := 0; sample < input.Rows; sample++ {

		outputNeurons := make([]float64, l.Size)

		for i := 0; i < l.Size; i++ {
			out, err := l.Neurons[i].Synapsis(input.GetRow(sample))
			if err != nil {
				panic(err.Error())
			}

			outputNeurons[i] = out
		}

		outputRow := output.GetRow(sample)
		copy(outputRow, outputNeurons)
	}

	return output
}
