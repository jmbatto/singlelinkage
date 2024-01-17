// Linkagepack project Linkagepack.go
package Linkagepack

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Rank1   int
	Length1 float64
	Rank2   int
	Length2 float64
}

type Leaf struct {
	Rank   int
	Length float64
}

type LeafIndex struct {
	Leaf
	Index int
}

type TreeNodeLeaf struct {
	TreeNode
	Leaf
}
type TreeRank map[int]TreeNodeLeaf
type Tree map[TreeNode]Leaf

// return a reduced matrix -without last column
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

// in place reduced matrix -without last column
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

// merge 2 column in place column max
func columnMerge(matrix [][]float64, min_col, max_col int) [][]float64 {
	for i := 0; i < max_col; i++ {
		if i != min_col && i != max_col {
			matrix[i][min_col] = findMini(matrix[i][min_col], matrix[i][max_col])
			matrix[min_col][i] = matrix[i][min_col]
		}
	}
	return removeLineColMatrix(matrix)
}

// return col A, col B, minDist, minDist-minDist1
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

func buildTree(matrix_start [][]float64) map[int]TreeNodeLeaf {
	T := map[int]TreeNodeLeaf{}
	i_start := len(matrix_start)
	matrix_start_i := matrix_start
	tree := Tree{
		//TreeNode{0, 0.0, 0, 0.0}: Leaf{},
	}
	insert_count := 0
	for i := i_start; i > 1; i-- {
		i1, i2, d, delta := minMatrix(matrix_start_i, i)
		memo_tree := TreeNode{i1, d, i2, d}
		memo_leaf := Leaf{i1, delta}
		if _, ok := tree[memo_tree]; ok {
			tmp := tree[memo_tree]
			var index_T int
			for kk, vv := range T {
				if vv.TreeNode == memo_tree {
					index_T = kk
					break
				}
			}
			if tmp.Length > delta {
				tree[memo_tree] = memo_leaf
				T[index_T] = TreeNodeLeaf{memo_tree, memo_leaf}
			}
		} else {
			tree[memo_tree] = memo_leaf
			T[insert_count] = TreeNodeLeaf{memo_tree, memo_leaf}
			insert_count++
		}
		matrix_start1 := columnMerge(matrix_start_i, i1, i2)
		matrix_start_i = matrix_start1
	}
	return T
}

//get a slice of sorted TreeNodeLeaf according Leaf Lenght for a Rank of interest
//could be improved with a buffer (to avoid recomputing the sorted TreeNodeLeaf
func sortTreeNodeLeaf(tree map[int]TreeNodeLeaf, index int) []TreeNodeLeaf {
	keysToSort := make([]TreeNodeLeaf, 0)
	for _, i_val := range tree {
		if i_val.Rank == tree[index].Rank {
			keysToSort = append(keysToSort, i_val)
		}
	}
	sort.Slice(keysToSort, func(i, j int) bool { return keysToSort[i].Length < keysToSort[j].Length })
	return keysToSort
}

func findIndexTreeNode(tree map[int]TreeNodeLeaf, T TreeNode, L Leaf) (int, bool) {
	for i_index, _ := range tree {
		if tree[i_index].Leaf == L && tree[i_index].TreeNode == T {
			return i_index, true
		}
	}
	return 0, false
}

func findNextLeaf(tree map[int]TreeNodeLeaf, index int) (T TreeNode, L Leaf, I int, found bool) {
	leaf := tree[index].Leaf
	sortedTree := sortTreeNodeLeaf(tree, index)
	for i_index, i_val := range sortedTree {
		if i_val.Length > leaf.Length { // find next
			L = i_val.Leaf
			//find in TreeNodeLeaf index of L
			T = sortedTree[i_index].TreeNode
			I, found := findIndexTreeNode(tree, T, L)
			return T, L, I, found
		}
	}
	return
}

func findNextLeafRecursive(tree map[int]TreeNodeLeaf, index int) (textleaf string, ok bool) {
	T, L, I, ok := findNextLeaf(tree, index)
	if ok == false {
		return "", ok
	}
	//recursive ?
	textleaf = fmt.Sprintf("(%d:%.2f,%d:%.2f)", T.Rank1, T.Length1, T.Rank2, T.Length2)
	memo_text, _ := findNextLeafRecursive(tree, I)
	return textleaf + memo_text + fmt.Sprintf(":%.2f", L.Length), ok
}

//recursive call to use a stack mechanism
func getLeafRecursive(tree map[int]TreeNodeLeaf, index int) (textleaf string, ok bool) {
	textleaf, ok = findNextLeafRecursive(tree, index)
	return textleaf, ok
}

//collect all the root Leaf (from where a subtree can be build)
func enumerateLeafRoot(tree map[int]TreeNodeLeaf) (retLeaf []LeafIndex) {
	tmp := map[int]LeafIndex{} //int = the Leaf Rank - to be changed
	for i_key, i_val := range tree {
		if _, ok := tmp[i_val.Rank]; ok {
			if tmp[i_val.Rank].Length > i_val.Length {
				tmp[i_val.Rank] = LeafIndex{i_val.Leaf, i_key}
			}
		} else {
			tmp[i_val.Rank] = LeafIndex{i_val.Leaf, i_key}
		}
	}
	for _, i_val := range tmp {
		retLeaf = append(retLeaf, i_val)
	}
	return
}

func processCurrentLeaf(tree map[int]TreeNodeLeaf, index int) (textleaf string) {
	T := tree[index]
	text1, ok := getLeafRecursive(tree, index)
	if ok {
		if T.Rank == T.Rank1 {
			textleaf = fmt.Sprintf("(%s,%d:%.2f):%.2f", text1, T.Rank2, T.Length2, T.Length)
		} else if T.Rank == T.Rank2 {
			textleaf = fmt.Sprintf("(%d:%.2f,%s):%.2f", T.Rank1, T.Length1, text1, T.Length)
		}
	} else {
		textleaf = fmt.Sprintf("(%d:%.2f,%d:%.2f):%.2f", T.Rank1, T.Length1, T.Rank2, T.Length2, T.Length)
	}
	return textleaf
}
