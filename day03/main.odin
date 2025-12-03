package main

import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

largest_joltage :: proc(s: string) -> int {
	largest_l := 0
	largest_index := 0
	for i := 0; i < len(s) - 1; i+=1 {
		num := strconv.parse_int(s[i:i+1]) or_else 0
		if num > largest_l {
			largest_l = num
			largest_index = i
		}
	}

	largest_r := -1
	for j := largest_index+1; j < len(s); j+=1 {
		num := strconv.parse_int(s[j:j+1]) or_else 0
		if num > largest_r {
			largest_r = num
		}
	}
	return largest_l * 10 + largest_r
}

// general function for joltage length of n
largest_joltage2 :: proc(s: string, n: int) -> i64 {
	assert(n < len(s))
	array := make([]int, len(s))
	for i in 0..<len(s) {
		array[i] = strconv.parse_int(s[i:i+1]) or_else 0
	}
	result: i64 = 0
	l_index := 0
	for i in 0..<n {
		largest := 0
		for j in l_index..<(len(s)-n+i+1) {
			if array[j] > largest {
				largest = array[j]
				l_index = j
			}
		}
		l_index += 1
		result *= 10
		result += i64(largest)
	}
	return result
}

main :: proc() {
	file, ok := os.read_entire_file("input.txt")
	if !ok {
		fmt.eprintln("couldn't open file")
		os.exit(1)
	}
	input := string(file)
	input2 := input

	fmt.println("--- Part 1 ---")

	total := 0
	for line in strings.split_lines_iterator(&input) {
		largest := largest_joltage(line)
		fmt.println(line, ":", "Largest joltage is:", largest)
		total += largest
	}
	fmt.println("Total joltage is:", total)

	fmt.println("--- Part 2 ---")

	total2: i64 = 0
	for line in strings.split_lines_iterator(&input2) {
		largest := largest_joltage2(line, 12)
		fmt.println(line, ":", "Largest joltage is:", largest)
		total2 += largest
	}
	fmt.println("Total joltage is:", total2)
}
