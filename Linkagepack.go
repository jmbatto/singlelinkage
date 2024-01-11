// Linkagepack project Linkagepack.go
package Linkagepack

import (
	"fmt"
	"io/ioutil"
)

func Linkage(matrice [][]float64, noms []string, taille int) (result string) {

	var i_min int
	var j_min int
	var min_val float64

	var old_index []int
	var dist []float64
	old_index = make([]int, taille)
	dist = make([]float64, taille)

	for i, _ := range old_index {
		old_index[i] = i
		dist[i] = 0
	}
	lenindex := len(old_index)
	for i := 1; i < lenindex; i++ {
		fmt.Printf("1: iteration %v %v\n", i, lenindex)
		//for i := len(old_index) - 1; i > 0; i-- {
		if i != 0 {
			fmt.Printf("iteration %v %v %v\n", old_index, i, matrice)
			i_min, j_min, min_val, dist[i] = minMatrix(matrice, old_index[i])
			min_col, max_col := minMaxMatrix(i_min, j_min)
			columnsMerge(matrice, old_index, min_col, max_col)
			writeName(noms, dist, min_col, max_col, min_val)
			//-------------------------------------------------------
			if i == 0 {
				result = result + noms[min_col]
			} else {
				result = result + "," + noms[min_col]
			}
			//-------------------------------------------------------
		}
	}

	/* au debut on a fait :
	chaine = noms[0] + ";"
	et cette commande recupere uniquementle premier arbre
	alors que il fallait faire une concatenation qu'on vient d'ajouter
	afin d'ecrire tous les arbres avant le terminateur point virgule */

	//chaine = noms[0] + ";"
	result = result + ";"

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
