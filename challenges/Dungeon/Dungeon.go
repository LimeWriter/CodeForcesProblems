/*
	https://codeforces.com/problemset/problem/1463/A
	A. Dungeon
	time limit per test
	2 seconds
	memory limit per test
	256 megabytes
	input
	standard input
	output
	standard output

	You are playing a new computer game in which you have to fight monsters. In a dungeon you are trying to clear, you met three monsters;
	the first of them has a health points, the second has b health points, and the third has c.

	To kill the monsters, you can use a cannon that, when fired, deals 1 damage to the selected monster.
	Every 7-th (i. e. shots with numbers 7, 14, 21 etc.) cannon shot is enhanced and deals 1 damage to all monsters, not just one of them.
	If some monster's current amount of health points is 0, it can't be targeted by a regular shot and does not receive damage from an enhanced shot.

	You want to pass the dungeon beautifully, i. e., kill all the monsters with the same enhanced shot
	(i. e. after some enhanced shot, the health points of each of the monsters should become equal to 0	for the first time).
	Each shot must hit a monster, i. e. each shot deals damage to at least one monster.

	Input
	The first line contains a single integer t
	(1≤t≤104) — the number of test cases.

	Each test case consists of a single line that contains three integers a, b and c (1≤a,b,c≤108) — the number of health points each monster has.

	Output
	For each test case, print YES if you can kill all the monsters with the same enhanced shot.
	Otherwise, print NO. You may print each letter in any case (for example, YES, Yes, yes, yEs will all be recognized as positive answer).

	Example
	Input
	3
	3 2 4
	1 1 1
	10 1 7

	Output
	YES
	NO
	NO

	Note

	In the first test case, you can do as follows: 1-th shot to the first monster, 2-th shot to the second monster,
	3-th shot to the third monster, 4-th shot to the first monster, 5-th shot to the third monster, 6-th shot to the third monster,
	and 7-th enhanced shot will kill all the monsters.

	In the second test case, you can't kill monsters with the same enhanced shot, because the total number of health points of monsters is 3, and you will kill them in the first 3 shots.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input_len = 3
const enhanced_shot_total_dmg = 9

func perfect_kill(hp_list string) bool {
	var hp_list_str_slice []string = strings.Split(hp_list, " ")
	var hp_list_int_slice [3]int64
	var rc bool = true
	var err error

	// Validate input length and convert input string into an int array
	if len(hp_list_str_slice) != input_len {
		rc = false
	} else {
		for i := 0; i < input_len; i++ {
			if hp_list_int_slice[i], err = strconv.ParseInt(hp_list_str_slice[i], 10, 64); err != nil {
				rc = false
			}
		}
	}

	if rc == true {
		var hp_sum, enhanced_shots int64
		var min_hp int64 = hp_list_int_slice[0]

		for i := 0; i < input_len; i++ {
			if min_hp > hp_list_int_slice[i] {
				min_hp = hp_list_int_slice[i]
			}

			hp_sum += hp_list_int_slice[i]
		}

		enhanced_shots = hp_sum / enhanced_shot_total_dmg

		// Total number of enhanced shots done needs to be greater than the monster with the lowest hp
		// If the total pool of hp between the 3 monsters is not divisible by 9 then it is not possble
		// to evenly distribute the damage on the enhanced shot
		if (enhanced_shots > min_hp) || (hp_sum%enhanced_shot_total_dmg != 0) {
			rc = false
		}
	}

	return rc
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Discard first line
	sc.Scan()

	for sc.Scan() {
		if perfect_kill(sc.Text()) == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
