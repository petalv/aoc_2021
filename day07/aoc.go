package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part := os.Getenv("part")
	crabs := readInput("input.txt")
	if part == "part2" {
		println(part2(crabs))
	} else {
		println(part1(crabs))
	}
}

func part1(crabs []int) int {
	sort.Ints(crabs)
	median := crabs[(len(crabs)/2)-1]
	fuel := 0
	for _, num := range crabs {
		fuel += Abs(num - median)
	}
	return fuel
}

func part2(crabs []int) int {
	sum := 0
	for _, num := range crabs {
		sum += num
	}
	avg := int(math.Floor(float64(sum) / float64(len(crabs))))
	fuel := 0
	for _, num := range crabs {
		fuel += AP(1, Abs(num-avg))
	}
	return fuel
}

func AP(commonDiff int, elements int) int {
	// (d*k*(k+1))/2 = (1*11*(11+1))/2
	return (commonDiff * elements * (elements + 1)) / 2

}
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func extractNumbers(line string, sep string) []int {
	var numbers []int
	for _, val := range strings.Split(line, sep) {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	return numbers
}

func readInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		data = extractNumbers(scanner.Text(), ",")
	}
	return data
}
