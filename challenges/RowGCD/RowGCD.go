package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	var n, m int
	var gcd, gcd_a int64 = 0, 0
	var err error
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 4)
	sc.Buffer(buf, 1024*1024*4)

	sc.Scan()

	nm := strings.Split(sc.Text(), " ")
	if n, err = strconv.Atoi(nm[0]); err != nil {
	}
	if m, err = strconv.Atoi(nm[1]); err != nil {
	}

	sc.Scan()

	a := strings.Split(sc.Text(), " ")
	a_int := make([]int64, n)
	for i := 0; i < n; i++ {
		if a_int[i], err = strconv.ParseInt(a[i], 10, 64); err != nil {
			panic(err)
		}
	}

	sc.Scan()

	b := strings.Split(sc.Text(), " ")
	b_int := make([]int64, m)
	for i := 0; i < m; i++ {
		if b_int[i], err = strconv.ParseInt(b[i], 10, 64); err != nil {
			panic(err)
		}
	}

	for i := 1; i < n; i++ {
		a_dif := a_int[i] - a_int[0]
		if a_dif < 0 {
			a_dif = -a_dif
		}

		gcd_a = find_gcd(a_dif, gcd)
	}

	for i := 0; i < m; i++ {
		gcd = find_gcd(gcd_a, b_int[i]+a_int[0])
		fmt.Printf("%v ", gcd)
	}
}
