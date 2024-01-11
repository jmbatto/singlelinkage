// Linkagepack project Linkagepack.go
package Linkagepack

import (
	"math"
)

//return a reduced matrix -without last column
func RemoveLineColMatrix(matrix [][]float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))
	// Parcourir chaque ligne de la matrice originale
	for i := 0; i < len(matrix); i++ {
		// Créer une nouvelle slice pour chaque ligne et copier ses éléments
		newMatrix[i] = make([]float64, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	columnIndex := len(newMatrix) - 1
	for i := 0; i < len(newMatrix); i++ {
		// Vérifie si la colonne à supprimer est valide
		if columnIndex >= 0 && columnIndex < len(matrix[i]) {
			newMatrix[i] = append(newMatrix[i][:columnIndex], newMatrix[i][columnIndex+1:]...)
		}
	}
	return newMatrix
}

//in place reduced matrix -without last column
func removeLineColMatrix(matrix [][]float64) [][]float64 {
	columnIndex := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		if columnIndex >= 0 && columnIndex < len(matrix[i]) {
			matrix[i] = append(matrix[i][:columnIndex], matrix[i][columnIndex+1:]...)
		}
	}
	//remove last column
	last := len(matrix) - 1
	matrix = matrix[:last]
	return matrix
}

//merge 2 column in place column max
func columnMerge(matrix [][]float64, min_col, max_col int) [][]float64 {
	for i := 0; i < max_col; i++ {
		if i != min_col && i != max_col {
			matrix[i][min_col] = findMini(matrix[i][min_col], matrix[i][max_col])
			matrix[min_col][i] = matrix[i][min_col]
		}
	}
	return removeLineColMatrix(matrix)
}

//return col A, col B, minDist, minDist-minDist1
func minMatrix(matrix [][]float64, size int) (int, int, float64, float64) {
	min := 1000.0  //math.MaxFloat64  //matrix[0][1]
	min1 := 1000.0 //math.MaxFloat64 //matrix[0][2]
	i_min := 0
	j_min := 1

	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if matrix[i][j] < min && i != j {
				min = matrix[i][j]
				i_min = i
				j_min = j
			}
		}
	}
	for i := 0; i < size; i++ {
		if matrix[i][j_min] < min1 && i != j_min && matrix[i][j_min] > min {
			min1 = matrix[i][j_min]
		}
	}
	return i_min, j_min, min, math.Abs(min1 - min)
}

//return value by increasing order
func minMaxMatrix(x, y int) (int, int) {
	if x < y {
		return x, y
	} else {
		return y, x
	}
}

func findMini(x, y float64) float64 {
	if x < y {
		return x
	} else {
		return y
	}
}
