/*
	https://codeforces.com/problemset/problem/1468/K

	K. The Robot
	time limit per test
	2 seconds
	memory limit per test
	512 megabytes
	input
	standard input
	output
	standard output

	There is a robot on a checkered field that is endless in all directions. Initially, the robot is located in the cell with coordinates (0,0).
	He will execute commands which are described by a string of capital Latin letters 'L', 'R', 'D', 'U'. When a command is executed, the robot simply moves in the corresponding direction:

		'L': one cell to the left (the x-coordinate of the current cell decreases by 1);
		'R': one cell to the right (the x-coordinate of the current cell is increased by 1);
		'D': one cell down (the y-coordinate of the current cell decreases by 1);
		'U': one cell up (the y-coordinate of the current cell is increased by 1).

	Your task is to put an obstacle in one cell of the field so that after executing the commands, the robot will return to the original cell of its path (0,0).
	Of course, an obstacle cannot be placed in the starting cell (0,0). It is guaranteed that if the obstacle is not placed, then the robot will not return to the starting cell.

	An obstacle affects the movement of the robot in the following way: if it tries to go in a certain direction, and there is an obstacle,
	then it simply remains in place (the obstacle also remains, that is, it does not disappear).

	Find any such cell of the field (other than (0,0)) that if you put an obstacle there, the robot will return to the cell (0,0) after the execution of all commands.
	If there is no solution, then report it.

	Input

	The first line contains one integer t (1≤t≤500) — the number of test cases.

	Each test case consists of a single line containing s — the sequence of commands, which are uppercase Latin letters 'L', 'R', 'D', 'U' only.
	The length of s is between 1 and 5000, inclusive. Additional constraint on s: executing this sequence of commands leads the robot to some cell other than (0,0), if there are no obstacles.

	The sum of lengths of all s	in a test doesn't exceed 5000.

	Output
	For each test case print a single line:

		if there is a solution, print two integers x and y (−109≤x,y≤109) such that an obstacle in (x,y) will force the robot to return back to the cell (0,0);
		otherwise, print two zeroes (i. e. 0 0).

	If there are multiple answers, you can print any of them.

	Example
	Input
	4
	L
	RUUDL
	LLUU
	DDDUUUUU
	Output
	-1 0
	1 2
	0 0
	0 1
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func find_point(path string) (int, int) {
	var h, v int = 0, 0

	for _, c := range path {
		if c == 'R' {
			h++
		} else if c == 'L' {
			h--
		} else if c == 'U' {
			v++
		} else if c == 'D' {
			v--
		}

		var x, y int = 0, 0
		for _, p := range path {
			p_x, p_y := x, y
			if p == 'R' {
				x++
			} else if p == 'L' {
				x--
			} else if p == 'U' {
				y++
			} else if p == 'D' {
				y--
			}

			// Block this movement and simulate the rest
			if x == h && y == v {
				x, y = p_x, p_y
			}
		}

		if x == 0 && y == 0 {
			return h, v
		}
	}

	return 0, 0
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var tests int
	fmt.Fscan(rd, &tests)

	for t := 0; t < tests; t++ {
		var s string

		fmt.Fscan(rd, &s)
		fmt.Println(find_point(s))
	}
}
