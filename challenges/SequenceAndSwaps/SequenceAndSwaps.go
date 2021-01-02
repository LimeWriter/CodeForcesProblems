/*
	https://codeforces.com/problemset/problem/1455/D

	D. Sequence and Swaps
	time limit per test
	1.5 seconds
	memory limit per test
	512 megabytes
	input
	standard input
	output
	standard output

	You are given a sequence a consisting of n integers a1,a2,…,an, and an integer x.
	Your task is to make the sequence a sorted (it is considered sorted if the condition a1≤a2≤a3≤⋯≤an holds).

	To make the sequence sorted, you may perform the following operation any number of times you want (possibly zero):
	choose an integer i such that 1≤i≤n and ai>x, and swap the values of ai and x.

	For example, if a=[0,2,3,5,4], x=1, the following sequence of operations is possible:

		choose i=2 (it is possible since a2>x), then a=[0,1,3,5,4], x=2;
		choose i=3 (it is possible since a3>x), then a=[0,1,2,5,4], x=3;
		choose i=4 (it is possible since a4>x), then a=[0,1,2,3,4], x=5.

	Calculate the minimum number of operations you have to perform so that a becomes sorted, or report that it is impossible.

	Input
	The first line contains one integer t (1≤t≤500) — the number of test cases.

	Each test case consists of two lines. The first line contains two integers n and x (1≤n≤500, 0≤x≤500) —
	the number of elements in the sequence and the initial value of x.

	The second line contains n integers a1, a2, ..., an (0≤ai≤500).

	The sum of values of n over all test cases in the input does not exceed 500.

	Output

	For each test case, print one integer — the minimum number of operations you have to perform to make a sorted, or −1, if it is impossible.

	Example
	Input
	6
	4 1
	2 3 5 4
	5 6
	1 1 3 4 4
	1 10
	2
	2 10
	11 9
	2 10
	12 11
	5 18
	81 324 218 413 324

	Output
	3
	0
	0
	-1
	1
	3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var tests int = 0
	fmt.Fscan(rd, &tests)

	for t := 0; t < tests; t++ {
		var n, x, a int
		swaps := 0

		fmt.Fscan(rd, &n, &x)

		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &a)
			arr[i] = a
		}

		for i := 0; i < n; i++ {
			if sort.IntsAreSorted(arr) == true {
				fmt.Println(swaps)
				break
			} else if arr[i] > x {
				tmp := arr[i]
				arr[i] = x
				x = tmp

				swaps++
			}
		}

		if sort.IntsAreSorted(arr) == false {
			fmt.Println(-1)
		}
	}
}
