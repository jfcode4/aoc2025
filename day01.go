package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction bool

const (
	Right Direction = true
	Left  Direction = false
)

type Rotation struct {
	Direction
	Distance int
}

func (d Direction) Str() string {
	switch d {
	case Right: return "R"
	case Left:  return "L"
	}
	return ""
}

func ParseLine(line []byte) (Rotation, error) {
	var r Rotation
	dir := line[0]
	if dir == 'R' {
		r.Direction= Right
	} else if dir == 'L' {
		r.Direction = Left
	} else {
		return r, fmt.Errorf("Unexpected '%c': expected 'R' or 'L'", dir)
	}
	distanceStr := string(line[1:])
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return r, fmt.Errorf("Unexpected '%s': expected a number", distanceStr)
	}
	r.Distance = distance
	return r, nil
}

func ParseInput(filename string) ([]Rotation, error) {
	result := make([]Rotation, 0)

	file, err := os.Open(filename)
	if err != nil {return nil, err}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum += 1
		rotation, err := ParseLine(scanner.Bytes())
		if err != nil {
			return result, fmt.Errorf("%s:%d: Parse error: %v", filename, lineNum, err)
		}
		result = append(result, rotation)
	}
	return result, nil
}

func main() {
	rotations, err := ParseInput("input01.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Total rotations: %d\n", len(rotations))

	dial := 50
	count := 0
	for _, rotation := range rotations {
		switch rotation.Direction {
		case Right:
			dial += rotation.Distance
		case Left:
			dial -= rotation.Distance
		}
		dial = (dial%100+100) % 100 // euclidian modulo
		if dial == 0 {
			count += 1
		}
	}
	fmt.Printf("The password is: %d\n", count)

	fmt.Println("Part Two")

	dial = 50
	count = 0
	for _, rotation := range rotations {
		toZero := 0
		switch rotation.Direction {
		case Right:
			toZero = 100 - dial
			dial += rotation.Distance
		case Left:
			toZero = (dial + 99) % 100 + 1 // dial if dial != 0 else 100
			dial -= rotation.Distance
		}
		dial = (dial%100+100) % 100 // euclidian modulo
		if rotation.Distance >= toZero {
			count += (rotation.Distance - toZero) / 100 + 1
		}
		//fmt.Printf("%s%d -> %d (count %d)\n", rotation.Direction.Str(), rotation.Distance, dial, count)
	}
	fmt.Printf("Actually, the password is: %d\n", count)
}
