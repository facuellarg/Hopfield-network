package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func ReadImage(imageName *string) image.Image {
	reader, err := os.Open(*imageName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func WriteImage(name string, img image.Image) error {
	outputFile, err := os.Create(name + ".png")
	if err != nil {
		println("valio verg")
		return err
	}
	return png.Encode(outputFile, img)

}

func VectorToImage(vector []float64, h, w, scale int) (image.Image, error) {
	if len(vector) != h*w {
		return nil, errors.New("dimensiones de la imagen no coinciden longitud del vector")
	}
	recImg := image.Rect(0, 0, h*scale, w*scale)
	myImage := image.NewRGBA(recImg)
	count := 0
	for i := 0; i < h*scale; i += scale {
		for j := 0; j < w*scale; j += scale {
			c := color.RGBA{
				inverseInputMaped[vector[count]],
				inverseInputMaped[vector[count]],
				inverseInputMaped[vector[count]],
				255,
			}
			count++
			for xscale := 0; xscale < scale; xscale++ {
				for yscale := 0; yscale < scale; yscale++ {
					myImage.Set(j+xscale, i+yscale, c)
				}
			}

		}
	}
	return myImage.SubImage(recImg), nil
}

func ImageToVector(name string) []float64 {
	img := ReadImage(&name)
	bounds := img.Bounds()
	input := make([]float64, bounds.Size().X*bounds.Size().X)

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			r, _, _, _ := img.At(j, i).RGBA()
			input[i*bounds.Max.Y+j] = inputsMaped[r]
		}
	}
	return input
}

func IsEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func sigmoid(x float64) float64 {
	return 1 / (1 - math.Exp(-x))
}

func WriteVectorImage(name string, vector []float64, h, w, scale int) {
	img, _ := VectorToImage(vector, h, w, scale)
	WriteImage(name, img)
}
