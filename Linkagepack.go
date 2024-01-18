// Linkagepack project Linkagepack.go
package Linkagepack

import (
	"fmt"
	"io/ioutil"
)

func MatrixLinkage(matrix [][]float64, colname []string) (result string) {

	return
}

func columnsMerge(matrix [][]float64, old_index []int, min_col, max_col int) []int {
	for _, i_old := range old_index {
		if i_old != min_col && i_old != max_col {
			matrix[i_old][min_col] = findMini(matrix[i_old][min_col], matrix[i_old][max_col])
			matrix[min_col][i_old] = matrix[i_old][min_col]
		}
	}

	var new_index int
	for i, v := range old_index {
		if v == max_col {
			new_index = i
			break
		}
	}

	for i, _ := range old_index[new_index : len(old_index)-1] {
		ind := i + new_index
		old_index[ind] = old_index[ind+1]
	}

	old_index1 := old_index[:len(old_index)-1]
	matrix = removeLineColMatrix(matrix)
	return old_index1
}

func writeName(noms []string, dist []float64, min_col, max_col int, min float64) {
	//noms[min_col] = fmt.Sprintf("(%s:%2f,%s:%2f)", noms[min_col], min-dist[min_col], noms[max_col], min-dist[max_col])
	noms[min_col] = fmt.Sprintf("(%s:%2f,%s:%2f)", noms[min_col], dist[min_col], noms[max_col], min-dist[max_col])
	dist[min_col] = min
}

func WriteFile(fileName, resultString string) (Err error) {
	Err = ioutil.WriteFile(fileName, []byte(resultString), 0640)
	return Err
}
