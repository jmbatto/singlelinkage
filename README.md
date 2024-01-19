# singlelinkage
This package provide a function to perform Single Linkage algorithm on a pairwise distance matrix.

```math
\begin{bmatrix}
A & B & C & D & E\cr
0 & 7.40 & 7.56 & 5.01 & 12.43 \cr
7.40 & 0 &  8.62 &  6.03 &  6.55 \cr
7.56 &  8.62 &  0 &  12.46 &  4.66 \cr
5.01 &  6.03 &  12.46 &  0 &  9.28 \cr
12.43 &  6.55 &  4.66 &  9.28 &  0 \cr
\end{bmatrix}
```
```golang
func MatrixLinkage(matrix [][]float64, colname []string) (result string)
```
The MatrixLinkage function is feeded with a [][]float64 matrix and a []string collecting column names {"A", "B", "C", "D", "E"}

The result is a string in [Newick format](https://en.wikipedia.org/wiki/Newick_format)

((C:4.66,E:4.66):1.89,((A:5.01,D:5.01):1.02,B:6.03):0.52);
