// Linkagepack_test
package Linkagepack

import (
	"strconv"
	"testing"
)

var matrix_start [][]float64 = [][]float64{
	{0, 7.40, 7.56, 5.01, 12.43},
	{7.40, 0, 8.62, 6.03, 6.55},
	{7.56, 8.62, 0, 12.46, 4.66},
	{5.01, 6.03, 12.46, 0, 9.28},
	{12.43, 6.55, 4.66, 9.28, 0},
}

var matrix_start1 [][]float64 = [][]float64{
	{0, 7.40, 7.56, 5.01, 12.43},
	{7.40, 0, 8.62, 6.03, 6.55},
	{7.56, 8.62, 0, 12.46, 4.66},
	{5.01, 6.03, 12.46, 0, 9.28},
	{12.43, 6.55, 4.66, 9.28, 0},
}

var matrix_intermediary1 [][]float64 = [][]float64{
	{0, 7.40, 7.56, 5.01},
	{7.40, 0, 6.55, 6.03},
	{7.56, 8.62, 0, 9.28},
	{5.01, 6.03, 9.28, 0},
}

var matrix_intermediary2 [][]float64 = [][]float64{
	{0, 6.03, 7.56},
	{6.03, 0, 6.55},
	{7.56, 6.55, 0},
}

var matrix_intermediary3 [][]float64 = [][]float64{
	{0, 6.55},
	{6.55, 0},
}

var matrix_colName []string = []string{"A", "B", "C", "D", "E"}

var target_result = "((C:4.66,E:4.66):1.89,((A:5.01,D:5.01):1.02,B:6.03):0.52);"

func TestWriteFile(t *testing.T) {
	_ = WriteFile("test.phb", target_result)
}

func TestMinMatrix1(t *testing.T) {
	a, b, c, d := minMatrix(matrix_start, 5)
	t.Logf("1: %v %v %v %v\n", a, b, c, d)
}

func TestMinMatrix2(t *testing.T) {
	a, b, c, d := minMatrix(matrix_intermediary1, 4)
	t.Logf("1: %v %v %v %v\n", a, b, c, d)
}

func TestMinMaxMatrix(t *testing.T) {
	a, b := minMaxMatrix(4, 3)
	t.Logf("1: %v %v \n", a, b)
}

func TestRemoveCol(t *testing.T) {
	t.Logf("1: %v \n", matrix_intermediary1)
	matrix := RemoveLineColMatrix(matrix_intermediary1)
	t.Logf("2: %v \n", matrix)
}

func TestColumnMerge1(t *testing.T) {
	//old_index := make([]int, 1)
	//old_index := []int{0, 1, 2, 3, 4}
	t.Logf("1: %v \n", matrix_start)
	matrix_start1 := columnMerge(matrix_start, 2, 4)
	t.Logf("2: %v \n", matrix_start1)

}

type TreeNode struct {
	Name string
	Data float64
}
type Tree map[TreeNode]Tree

func TestColumnMerge11(t *testing.T) {
	t.Logf("1: %v \n", matrix_start)
	i_start := len(matrix_start) - 1
	matrix_start_i := matrix_start1
	tree := Tree{
		TreeNode{"root", 0.0}: Tree{TreeNode{"sibling", 0.0}: nil},
	}
	//tree[TreeNode{"second", 1.0}] = Tree{}
	a, b, c, d := minMatrix(matrix_start_i, 5)
	t.Logf("1: %v %v %v %v\n", a, b, c, d)
	for i := i_start; i > 0; i-- {
		i1, i2, d, delta := minMatrix(matrix_start_i, i)
		tree[TreeNode{"i1_" + strconv.Itoa(i1), d}] = Tree{}
		matrix_start1 := columnMerge(matrix_start_i, i1, i2)
		matrix_start_i = matrix_start1
		t.Logf("2: %v %v %v %v\n", matrix_start_i, d, delta, tree)
	}
}

func TestColumnMerge2(t *testing.T) {
	//old_index := make([]int, 1)
	old_index := []int{0, 1, 2, 3}
	t.Logf("1: %v \n", matrix_intermediary1)
	columnsMerge(matrix_intermediary1, old_index, 0, 3)
	t.Logf("2: %v \n", matrix_intermediary1)
	t.Logf("3: %v \n", old_index)
}

func TestLinkage(t *testing.T) {
	t.Logf("1: %v \n", matrix_start1)
	result := Linkage(matrix_start1, matrix_colName, 4)
	t.Logf("2: %v \n", matrix_start1)
	t.Logf("2: %v \n", result)
	t.Logf("3: %v \n", target_result)
}
