/*
	https://codeforces.com/contest/1462/problem/D

	D. Add to Neighbour and Remove
	time limit per test
	3 seconds
	memory limit per test
	256 megabytes
	input
	standard input
	output
	standard output

	Polycarp was given an array of a[1…n]
	of n integers. He can perform the following operation with the array a no more than n times:

	Polycarp selects the index i and adds the value ai to one of his choice of its neighbors. More formally, Polycarp adds the value of ai to ai−1 or to ai+1
	(if such a neighbor does not exist, then it is impossible to add to it).
	After adding it, Polycarp removes the i-th element from the a array. During this step the length of a is decreased by 1.

	The two items above together denote one single operation.

	For example, if Polycarp has an array a=[3,1,6,6,2], then it can perform the following sequence of operations with it:

	Polycarp selects i=2 and adds the value ai to (i−1)-th element: a=[4,6,6,2].
	Polycarp selects i=1 and adds the value ai to (i+1)-th element: a=[10,6,2].
	Polycarp selects i=3 and adds the value ai to (i−1)-th element: a=[10,8].
	Polycarp selects i=2 and adds the value ai to (i−1)-th element: a=[18].

	Note that Polycarp could stop performing operations at any time.

	Polycarp wondered how many minimum operations he would need to perform to make all the elements of a
	equal (i.e., he wants all ai are equal to each other).
	Input

	The first line contains a single integer t
	(1≤t≤3000) — the number of test cases in the test. Then t

	test cases follow.

	The first line of each test case contains a single integer n
	(1≤n≤3000) — the length of the array. The next line contains n integers a1,a2,…,an (1≤ai≤105) — array a.

	It is guaranteed that the sum of n
	over all test cases does not exceed 3000.

	Output
	For each test case, output a single number — the minimum number of operations that Polycarp needs to perform so that all elements of the a array are the same (equal).

	Example
	Input
	4
	5
	3 1 6 6 2
	4
	1 2 2 1
	3
	2 2 2
	4
	6 3 2 1

	Output
	4
	2
	0
	2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_additions(int_slice []int, slice_len int) int {
	var additions, sum int = 0, 0

	for _, v := range int_slice {
		sum += v
	}

	for {
		// Sum of all elements must be divisible by number of elements
		if sum%(slice_len-additions) == 0 {
			all_equal := true
			sub_arr_sum := sum / (slice_len - additions)
			sub_sum := 0

			for j := 0; j < slice_len; j++ {
				sub_sum += int_slice[j]

				// A sub array must sum to the quotient of the (array total)/(elements-additions made)
				if sub_sum > sub_arr_sum {
					all_equal = false
					break
				} else if sub_sum == sub_arr_sum {
					sub_sum = 0
				}
			}
			if all_equal == true {
				break
			}
		}

		additions++
	}

	return additions
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Discard first line
	sc.Scan()

	for sc.Scan() {
		slice_len, err := strconv.Atoi(sc.Text())

		if err != nil {
			break
		}

		int_slice := make([]int, slice_len)

		sc.Scan()
		str_slice := strings.Split(sc.Text(), " ")

		for i := 0; i < slice_len; i++ {
			int_slice[i], err = strconv.Atoi(str_slice[i])
		}

		fmt.Println(find_additions(int_slice, slice_len))
	}
}
