/*
	https://codeforces.com/problemset/problem/1461/B

	B. Find the Spruce
	time limit per test
	1 second
	memory limit per test
	256 megabytes
	input
	standard input
	output
	standard output

	Holidays are coming up really soon. Rick realized that it's time to think about buying a traditional spruce tree.
	But Rick doesn't want real trees to get hurt so he decided to find some in an n×m matrix consisting of "*" and ".".

	To find every spruce first let's define what a spruce in the matrix is. A set of matrix cells is called a spruce of height k
	with origin at point (x,y)

	if:
		All cells in the set contain an "*".
		For each 1≤i≤k all cells with the row number x+i−1 and columns in range [y−i+1,y+i−1] must be a part of the set. All other cells cannot belong to the set.

	Now Rick wants to know how many spruces his n×m matrix contains. Help Rick solve this problem.

	Input
	Each test contains one or more test cases. The first line contains the number of test cases t
	(1≤t≤10).

	The first line of each test case contains two integers n and m (1≤n,m≤500) — matrix size.

	Next n lines of each test case contain m characters ci,j — matrix contents. It is guaranteed that ci,j is either a "." or an "*".

	It is guaranteed that the sum of n⋅m over all test cases does not exceed 5002 (∑n⋅m≤500^2).

	Output
	For each test case, print single integer — the total number of spruces in the matrix.

	Example
	Input
	4
	2 3
	.*.
	***
	2 3
	.*.
	**.
	4 5
	.***.
	*****
	*****
	*.*.*
	5 7
	..*.*..
	.*****.
	*******
	.*****.
	..*.*..

	Output
	5
	3
	23
	34
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_num_spruce(matrix [][]string, n int, m int) int {
	var num_spruces, depth int = 0, 0
	spruce_widths := make([][]int, n)

	// Prefix sum the number of spruces until an 'empty' cell is found
	for r := 0; r < n; r++ {
		spruce_widths[r] = make([]int, m)
		spruce_cells := 0

		for c := 0; c < m; c++ {

			if matrix[r][c] == "." {
				spruce_cells = 0
			} else {
				spruce_cells++
			}

			spruce_widths[r][c] = spruce_cells
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			depth = 0

			if spruce_widths[r][c] > 0 {
				num_spruces++

				for wr := r + 1; wr < n; wr++ {
					// We have a contiguous set of spruces with a width of depth+1 if the prefix sum of the left cell + (depth*1+1) is equal to the right cell - 1
					if (c-depth-1 >= 0 && c+depth+1 < m) && (spruce_widths[wr][c-depth-1] > 0) && (spruce_widths[wr][c-depth-1]+(depth*2+1) == spruce_widths[wr][c+depth+1]-1) {
						depth++
						num_spruces++
					} else {
						break
					}
				}
			}
		}
	}

	return num_spruces
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Discard first line
	sc.Scan()

	for sc.Scan() {
		var n, m int
		var err error
		size := strings.Split(sc.Text(), " ")

		if n, err = strconv.Atoi(size[0]); err != nil {
			panic(err)
		}
		if m, err = strconv.Atoi(size[1]); err != nil {
			panic(err)
		}

		spruce_matrix := make([][]string, n)

		for i := 0; i < n; i++ {
			sc.Scan()
			spruce_matrix[i] = strings.Split(sc.Text(), "")
		}

		fmt.Println(find_num_spruce(spruce_matrix, n, m))
	}
}
