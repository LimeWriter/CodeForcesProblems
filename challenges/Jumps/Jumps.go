/*
	https://codeforces.com/contest/1455/problem/B

	B. Jumps
	time limit per test
	1 second
	memory limit per test
	256 megabytes
	input
	standard input
	output
	standard output

	You are standing on the OX-axis at point 0 and you want to move to an integer point x>0.

	You can make several jumps. Suppose you're currently at point y (y may be negative) and jump for the k-th time. You can:

		either jump to the point y+k
		or jump to the point y−1.

	What is the minimum number of jumps you need to reach the point x?

	Input
	The first line contains a single integer t (1≤t≤1000) — the number of test cases.
	The first and only line of each test case contains the single integer x (1≤x≤106) — the destination point.

	Output
	For each test case, print the single integer — the minimum number of jumps to reach x.
	It can be proved that we can reach any integer point x.

	Example
	Input
	5
	1
	2
	3
	4
	5

	Output
	1
	3
	2
	3
	4
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var tests int
	fmt.Fscan(rd, &tests)

	for t := 0; t < tests; t++ {
		var d, p, jumps int = 0, 0, 0
		fmt.Fscan(rd, &d)

		for i := 1; p < d; i++ {
			jumps = i
			p += i
		}

		if p == d+1 {
			jumps++
		}

		fmt.Println(jumps)
	}
}
