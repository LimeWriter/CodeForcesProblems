/*
	https://codeforces.com/problemset/problem/1469/B

	B. Red and Blue
	time limit per test
	2 seconds
	memory limit per test
	512 megabytes
	input
	standard input
	output
	standard output

	Monocarp had a sequence a consisting of n+m integers a1,a2,…,an+m. He painted the elements into two colors, red and blue;
	n elements were painted red, all other m elements were painted blue.

	After painting the elements, he has written two sequences r1,r2,…,rn and b1,b2,…,bm. The sequence r consisted of all red elements of a in the order they appeared in a;
	similarly, the sequence b consisted of all blue elements of a in the order they appeared in a as well.

	Unfortunately, the original sequence was lost, and Monocarp only has the sequences r and b.
	He wants to restore the original sequence. In case there are multiple ways to restore it,
	he wants to choose a way to restore that maximizes the value of f(a)=max(0,a1,(a1+a2),(a1+a2+a3),…,(a1+a2+a3+⋯+an+m))

	Help Monocarp to calculate the maximum possible value of f(a).

	Input
	The first line contains one integer t (1≤t≤1000) — the number of test cases. Then the test cases follow. Each test case consists of four lines.

	The first line of each test case contains one integer n (1≤n≤100).

	The second line contains n integers r1,r2,…,rn (−100≤ri≤100).

	The third line contains one integer m (1≤m≤100).

	The fourth line contains m integers b1,b2,…,bm (−100≤bi≤100).

	Output
	For each test case, print one integer — the maximum possible value of f(a).

	Example
	Input
	4
	4
	6 -5 7 -3
	3
	2 3 -4
	2
	1 1
	4
	10 -3 2 2
	5
	-1 -2 -3 -4 -5
	5
	-1 -2 -3 -4 -5
	1
	0
	1
	0

	Output
	13
	13
	0
	0
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
		var r, b, rn, bn, sum, r_max, b_max int

		fmt.Fscan(rd, &rn)
		for i := 0; i < rn; i++ {
			fmt.Fscan(rd, &r)

			sum += r

			if sum > r_max {
				r_max = sum
			}
		}

		sum = 0

		fmt.Fscan(rd, &bn)
		for i := 0; i < bn; i++ {
			fmt.Fscan(rd, &b)

			sum += b

			if sum > b_max {
				b_max = sum
			}
		}

		fmt.Println(r_max + b_max)
	}
}
