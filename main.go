package main

import (
	_ "image/png"

	"github.com/google/gxui/drivers/gl"
)

var (
	inputs            [][]float64
	imagesNames       = []string{"A.png", "C.png", "F.png"}
	inputsMaped       = map[uint32]float64{65535: -1.0, 0: 1}
	inverseInputMaped = map[float64]uint8{-1.0: 255, 1: 0}
	steps             [][]float64
)

func main() {
	nn := NewHopfieldNN(64, 0, 0)
	inputs = make([][]float64, len(imagesNames))
	for i, name := range imagesNames {
		inputs[i] = ImageToVector(name)

	}
	nn.Train(inputs)
	steps = Guess(&nn, ImageToVector("algo.png"))

	WriteVectorImage("final", steps[len(steps)-1], 8, 8, 8)
	gl.StartDriver(appMain)
}
