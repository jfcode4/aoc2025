package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseRange(s string) (int, int, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Unexpected '%s': expected a range", s)
	}
	begin, err1 := strconv.Atoi(parts[0])
	end,   err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("Unexpected '%s': expected a range of numbers", s)
	}
	return begin, end, nil
}

func IsRepeating(num int) bool {
	str := strconv.Itoa(num)
	if len(str) % 2 == 1 {
		return false
	}
	middle := len(str) / 2
	left := str[0:middle]
	right := str[middle:]
	return left == right
}

func IsRepeatingMany(num int) bool {
	str := strconv.Itoa(num)
	for i := 1; i <= len(str)/2; i++ {
		if len(str) % i != 0 {
			continue
		}
		begin := str[0:i]
		repeating := true
		for j := i; j < len(str); j+=i {
			next := str[j:j+i]
			if (next != begin) {
				repeating = false
				break
			}
		}
		if repeating {
			return true
		}
	}
	return false
}

func main() {
	filename := "input02.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Cannot open file %s: %v", filename, err)
		os.Exit(1)
	}
	fileStr := strings.TrimSpace(string(file))

	ranges := strings.Split(fileStr, ",")


	fmt.Println("--- Part 1 --- ")
	total := 0

	for _, r := range ranges {
		begin, end, err := ParseRange(r)
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			return
		}
		for i := begin; i <= end; i++ {
			if IsRepeating(i) {
				//fmt.Println("Found repeating:", i)
				total += i
			}
		}
	}
	fmt.Println("Total is: ", total)

	fmt.Println("--- Part 2 --- ")
	total = 0

	for _, r := range ranges {
		begin, end, err := ParseRange(r)
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			return
		}
		for i := begin; i <= end; i++ {
			if IsRepeatingMany(i) {
				//fmt.Println("Found repeating:", i)
				total += i
			}
		}
	}
	fmt.Println("Total is: ", total)
}
