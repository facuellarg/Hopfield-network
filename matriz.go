package main

import (
	"bytes"
	"errors"
	"fmt"
)

type Matrix [][]float64

func NewMatrix(rows, columns int) Matrix {
	m := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, columns)
	}
	return m
}

func MatrixAdition(a, b Matrix) (error, Matrix) {
	if (len(a) != len(b)) || (len(a[0]) != len(b[0])) {
		return errors.New("No se pueden sumar matrices de tamaños diferentes"), nil
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(b))
		for j := range a[i] {
			m[i][j] = a[i][j] + b[i][j]
		}
	}
	return nil, Matrix(m)
}

func MatrixMultiplication(a, b Matrix) (error, Matrix) {
	if len(a[0]) != len(b) {
		return errors.New("Los tamaños entre columnas y filas no coinciden"), nil
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(b[0]))
		for j := range b[0] {
			for k := range b {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return nil, Matrix(m)
}

func MatrixMultiplicationSpecific(a, b Matrix, specific string) (error, Matrix) {
	if specific == "hopfield" {
		if len(a[0]) != len(b) {
			return errors.New("Los tamaños entre columnas y filas no coinciden"), nil
		}
		m := make([][]float64, len(a))
		for i := range a {
			m[i] = make([]float64, len(b[0]))
			for j := range b[0] {
				if i == j {
					m[i][j] = 0
					continue
				}
				for k := range b {

					m[i][j] += a[i][k] * b[k][j]
				}
			}
		}
		return nil, Matrix(m)
	} else {
		return MatrixMultiplication(a, b)
	}
}

func Map(a Matrix, f func(float64) float64) Matrix {
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(a[i]))
		for j := range a[i] {
			m[i][j] = f(a[i][j])
		}
	}
	return m
}

func MatrixScalar(a Matrix, b float64) Matrix {
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = make([]float64, len(a[0]))
		for j := range a[i] {
			m[i][j] = a[i][j] * b
		}
	}
	return Matrix(m)
}

func PrintMatrix(a Matrix) {
	for i := range a {
		for j := range a[0] {
			print(a[i][j], " ")
		}
		println()
	}
}

func Transpose(a Matrix) Matrix {
	m := NewMatrix(len(a[0]), len(a))
	for i := range a {
		for j := range a[0] {
			m[j][i] = a[i][j]
		}
	}
	return m
}

func (m Matrix) String() string {
	var buffer bytes.Buffer
	for i := range m {
		for j := range m[0] {
			buffer.WriteString(fmt.Sprintf("%2.2f ", m[i][j]))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
