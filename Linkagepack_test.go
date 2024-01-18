// Linkagepack_test
package Linkagepack

import (
	"math"
	"strings"
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

// func TestColumnMerge1(t *testing.T) {
// 	//old_index := make([]int, 1)
// 	//old_index := []int{0, 1, 2, 3, 4}
// 	t.Logf("1: %v \n", matrix_start)
// 	matrix_start1 := columnMerge(matrix_start, 2, 4)
// 	t.Logf("2: %v \n", matrix_start1)

// }

func TestAlgoMatrixLinkage(t *testing.T) {
	t.Logf("Input Matrix : %v \n", matrix_start)
	i_start := len(matrix_start)
	matrix_start_i := matrix_start
	tree := Tree{
		//TreeNode{0, 0.0, 0, 0.0}: Leaf{},
	}

	for i := i_start; i > 1; i-- {
		i1, i2, d, delta := minMatrix(matrix_start_i, i)
		memo_tree := TreeNode{i1, d, i2, d}
		memo_leaf := Leaf{i1, delta}
		if _, ok := tree[memo_tree]; ok {
			tmp := tree[memo_tree]
			if tmp.Length > delta {
				tree[memo_tree] = memo_leaf
			}
		} else {
			tree[memo_tree] = memo_leaf
		}
		matrix_start1 := columnMerge(matrix_start_i, i1, i2)
		matrix_start_i = matrix_start1
		t.Logf("Intermediate matrix result: %v \n %v %v \n computed map %v\n", matrix_start_i, d, delta, tree)

	}

}

var test_tree = map[int]TreeNodeLeaf{ //ordered according to insertion, continuous
	//0: TreeNodeLeaf{TreeNode{0, 0, 0, 0}, Leaf{0, 1000.0}},
	2: TreeNodeLeaf{TreeNode{0, 5.01, 3, 5.01}, Leaf{0, 1.02}},
	3: TreeNodeLeaf{TreeNode{0, 6.03, 1, 6.03}, Leaf{0, 0.52}},
	1: TreeNodeLeaf{TreeNode{2, 4.66, 4, 4.66}, Leaf{2, 1.89}},
}

func TestAlgoMatrixLinkageProduceString(t *testing.T) {
	i_mini_old := minLength
	i_mini := i_mini_old
	var i_key_memo int
	var i_rank int
	var T TreeNode
	// var L Leaf
	// var I int
	// var ok bool
	for i_key, i_val := range test_tree {
		i_mini = math.Min(i_mini_old, i_val.Length)
		if i_mini != i_mini_old {
			i_key_memo = i_key
			i_rank = i_val.Rank
			T = i_val.TreeNode
		}
		i_mini_old = i_mini
	}
	t.Logf("Find root Leaf : %v %v %v %v \n", T, i_rank, i_mini, i_key_memo)
	// T, L, I, ok = findNextLeaf(test_tree, i_key_memo)
	// t.Logf("22 1: %v %v %v %v \n", T, L, I, ok)
	// T, L, I, ok = findNextLeaf(test_tree, I)
	// t.Logf("22 2: %v %v %v %v \n", T, L, I, ok)
	// T, L, I, ok = findNextLeaf(test_tree, I)
	// t.Logf("22 3: %v %v %v %v \n", T, L, I, ok)
	textfeaf, _ := getLeafRecursive(test_tree, i_key_memo)
	t.Logf("getLeafRecursive: %v \n", textfeaf)
	retLeaf := enumerateLeafRoot(test_tree)
	t.Logf("enumerateLeafRoot: %v \n", retLeaf)
	var response string
	for _, i_var := range retLeaf {
		response = processCurrentLeaf(test_tree, i_var.Index) + "," + response
	}
	t.Logf("computed String : %v \n", "("+strings.TrimSuffix(response, ",")+");")
	test_tree2 := buildTree(matrix_start1)
	t.Logf("targetTree with shifted index : %v \n", test_tree)
	t.Logf("computed buildTree (from index 0) : %v \n", test_tree2)
	//Linkage Algo
	retLeaf2 := enumerateLeafRoot(test_tree2)
	var response2 string
	for _, i_var := range retLeaf2 {
		response2 = processCurrentLeaf(test_tree2, i_var.Index) + "," + response2
	}
	t.Logf("Target result : %v \n", target_result)
	t.Logf("Linkage Algo String Computed : %v \n", "("+strings.TrimSuffix(response2, ",")+");")
}

func TestColumnsMerge(t *testing.T) {
	test_index := []int{0, 1, 2, 3}
	t.Logf("input matrix_intermediary1 : %v \n", matrix_intermediary1)
	columnsMerge(matrix_intermediary1, test_index, 0, 3)
	t.Logf("merged matrix_intermediary1 : %v \n", matrix_intermediary1)
	//t.Logf("test_index : %v \n", test_index)
}
