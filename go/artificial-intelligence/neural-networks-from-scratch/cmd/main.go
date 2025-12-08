package main

import "fmt"

func main() {
	var input []float32 = []float32{1.0, 2.0, 3.0, 2.5}
	var weights []float32 = []float32{0.2, 0.8, -0.5, 1.0}
	var bias float32 = 2.0

	var output float32 = (
		input[0] * weights[0] +
		input[1] * weights[1] +
		input[2] * weights[2] +
		input[3] * weights[3] +
		bias)

	fmt.Printf("%f \n", output)
}