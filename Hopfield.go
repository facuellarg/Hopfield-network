package main

import (
	"math/rand"
	"time"
)

type HopfieldNN struct {
	Neurons      []float64
	Wieghts      Matrix
	Energy       int
	LearningRate float64
	Threshold    float64
}

func NewHopfieldNN(neurons int, lr, th float64) HopfieldNN {
	Neurons := make([]float64, neurons)
	wieghts := NewMatrix(neurons, neurons)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < neurons; i++ {
	// 	for j := i; j < neurons; j++ {
	// 		if i == j {
	// 			wieghts[i][j] = 0
	// 		} else {
	// 			wieghts[i][j] = rand.Float64()*2 - 2
	// 			wieghts[j][i] = wieghts[i][j]
	// 		}
	// 	}
	// }
	hp := HopfieldNN{
		Neurons:      Neurons,
		Wieghts:      wieghts,
		Energy:       0,
		LearningRate: lr,
		Threshold:    th,
	}
	return hp
}

func (nn *HopfieldNN) Train(inputs [][]float64) {
	for _, input := range inputs {
		p := [][]float64{input}
		p_t := Transpose(Matrix(p))
		_, w_inputs := MatrixMultiplicationSpecific(p_t, p, "hopfield")
		_, nn.Wieghts = MatrixAdition(w_inputs, nn.Wieghts)
	}
}

func Guess(nn *HopfieldNN, input []float64) [][]float64 {
	stepsHP := make([][]float64, 1)
	rand.Seed(time.Now().UnixNano())
	tmp := make([]float64, len(input))
	copy(tmp, input)
	stepsHP[0] = tmp

	for {

		err, output := MatrixMultiplication([][]float64{input}, nn.Wieghts)
		if err != nil {
			println(err.Error())
		}

		output = Map(output, func(a float64) float64 {
			if a > nn.Threshold {
				return 1.0
			}
			return -1.0
		})
		for _, r := range rand.Perm(len(input)) {
			if output[0][r] != input[r] {
				tmp := make([]float64, len(input))
				copy(tmp, stepsHP[len(stepsHP)-1])
				tmp[r] = output[0][r]
				stepsHP = append(stepsHP, tmp)
			}
		}
		if IsEqual(input, (output[0])) {
			break
		}
		copy(input, output[0])
	}
	println(len(stepsHP))
	return stepsHP
}
