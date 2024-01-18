// Linkagepack project Linkagepack.go
package Linkagepack

import (
	"io/ioutil"
	"strings"
)

//a matrix with float64 values, an array with column name
func MatrixLinkage(matrix [][]float64, colname []string) (result string) {
	//Linkage Algo
	test_tree2 := buildTree(matrix)
	retLeaf2 := enumerateLeafRoot(test_tree2)
	var response2 string
	for _, i_var := range retLeaf2 {
		response2 = processCurrentLeaf(test_tree2, i_var.Index, colname) + "," + response2
	}
	result = "(" + strings.TrimSuffix(response2, ",") + ");"
	return
}

func WriteFile(fileName, resultString string) (Err error) {
	Err = ioutil.WriteFile(fileName, []byte(resultString), 0640)
	return Err
}
