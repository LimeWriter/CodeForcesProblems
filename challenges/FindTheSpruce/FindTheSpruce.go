package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_num_spruce(matrix [][]string, n int, m int) int {
	var num_spruce, chain int = 0, 0
	var spruce_end bool = false

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if matrix[j][i] == "*" {
				num_spruce++
				chain = 1

				for l := j + 1; l < n; l++ {
					if i-chain >= 0 && i+chain < m {
						for c := i - chain; c <= i+chain; c++ {
							if matrix[l][c] == "." {
								spruce_end = true
								break
							}
						}

						if spruce_end == true {
							spruce_end = false
							break
						} else {
							num_spruce++
						}
					}
					chain++
				}
			}
		}
	}

	return num_spruce
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
