package layerbasedneurons

import (
	"math/rand"

	singleneuron "github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/layer-based-neurons/single-neuron"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/matrix"
	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

type LayerBasedNeurons[T vector.Real] struct {
	Size      int
	InputSize int
	Neurons   []singleneuron.Neuron[T]
}

func NewLayerBasedNeurons[T vector.Real](size, inputSize int, af func(T) T) *LayerBasedNeurons[T] {
	if size < 1 || inputSize < 1 {
		panic("WTF! an empty or negative layer/input size?")
	}

	layer := &LayerBasedNeurons[T]{
		Size:      size,
		InputSize: inputSize,
		Neurons:   make([]singleneuron.Neuron[T], size),
	}

	for i := 0; i < size; i++ {
		layer.Neurons[i] = *singleneuron.NewNeuron(inputSize, af)
	}

	return layer
}

func (l *LayerBasedNeurons[T]) SetWeights(m matrix.Matrix[T]) {
	if m.Shape()[1] != l.InputSize || m.Shape()[0] != l.Size {
		panic("Dimensions mismatch!")
	}

	for i := 0; i < l.Size; i++ {
		l.Neurons[i].SetWeights(m.GetRow(i))
	}
}

func (l *LayerBasedNeurons[T]) SetBiases(biases []T) {
	if len(biases) != l.Size {
		panic("Size mismatch!")
	}

	for i := 0; i < l.Size; i++ {
		l.Neurons[i].Bias = biases[i]
	}
}

func (l *LayerBasedNeurons[T]) RandomizeWeights() {
	for i := 0; i < l.Size; i++ {
		l.Neurons[i].RandomizeWeights()
	}
}

func (l *LayerBasedNeurons[T]) RandomizeBiases() {
	for i := 0; i < l.Size; i++ {
		l.Neurons[i].Bias = T((rand.Float64() - 0.5) * 10)
	}
}

func (l *LayerBasedNeurons[T]) Forward(input []T) []T {
	output := make([]T, l.Size)

	for i := 0; i < l.Size; i++ {
		out, err := l.Neurons[i].Synapsis(input)
		if err != nil {
			panic(err.Error())
		}

		output[i] = out
	}

	return output
}

func (l *LayerBasedNeurons[T]) ForwardBatch(input *matrix.Matrix[T]) *matrix.Matrix[T] {
	newShape := []int{input.Shape()[0], l.Size}
	output := matrix.NewMatrix[T](newShape)

	for sample := 0; sample < input.Shape()[0]; sample++ {

		outputNeurons := make([]T, l.Size)

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

func (l *LayerBasedNeurons[T]) Copy() *LayerBasedNeurons[T] {
	newNeurons := make([]singleneuron.Neuron[T], len(l.Neurons))

	for i := 0; i < len(l.Neurons); i++ {
		newNeurons[i] = *l.Neurons[i].Copy()
	}

	return &LayerBasedNeurons[T]{
		Size:      l.Size,
		InputSize: l.InputSize,
		Neurons:   newNeurons,
	}
}
