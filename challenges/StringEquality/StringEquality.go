/*
	https://codeforces.com/problemset/problem/1451/C

	C. String Equality
	time limit per test
	2 seconds
	memory limit per test
	256 megabytes
	input
	standard input
	output
	standard output

	Ashish has two strings a and b, each of length n, and an integer k. The strings only contain lowercase English letters.

	He wants to convert string a into string b by performing some (possibly zero) operations on a.

	In one move, he can either

		choose an index i (1≤i≤n−1) and swap ai and ai+1, or
		choose an index i (1≤i≤n−k+1) and if ai,ai+1,…,ai+k−1 are all equal to some character c (c≠ 'z'),
			replace each one with the next character (c+1), that is, 'a' is replaced by 'b', 'b' is replaced by 'c' and so on.

	Note that he can perform any number of operations, and the operations can only be performed on string a.

	Help Ashish determine if it is possible to convert string a into b after performing some (possibly zero) operations on it.

	Input
	The first line contains a single integer t
	(1≤t≤105) — the number of test cases. The description of each test case is as follows.

	The first line of each test case contains two integers n (2≤n≤106) and k (1≤k≤n).

	The second line of each test case contains the string a of length n consisting of lowercase English letters.

	The third line of each test case contains the string b of length n consisting of lowercase English letters.

	It is guaranteed that the sum of values n among all test cases does not exceed 10^6.\

	Output
	For each test case, print "Yes" if Ashish can convert a into b after some moves, else print "No".

	You may print the letters of the answer in any case (upper or lower).

	Example
	Input
	4
	3 3
	abc
	bcd
	4 2
	abba
	azza
	2 1
	zz
	aa
	6 2
	aaabba
	ddddcc

	Output
	No
	Yes
	No
	Yes
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 4)
	sc.Buffer(buf, 1024*1024)

	sc.Scan()
	num_tests, _ := strconv.Atoi(sc.Text())

	for i := 0; i < num_tests; i++ {
		sc.Scan()

		nk := strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(nk[0])
		k, _ := strconv.Atoi(nk[1])

		sc.Scan()
		str_a := sc.Text()

		sc.Scan()
		str_b := sc.Text()

		a_count := make([]int, 26)
		b_count := make([]int, 26)

		for j := 0; j < n; j++ {
			a_count[str_a[j]-'a'] += 1
			b_count[str_b[j]-'a'] += 1
		}

		var str_eq bool = true

		for j := 0; j < 25; j++ {
			if b_count[j] > a_count[j] || (a_count[j]-b_count[j])%k != 0 {
				str_eq = false
				break
			}

			a_count[j+1] += a_count[j] - b_count[j]
		}

		if str_eq == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
