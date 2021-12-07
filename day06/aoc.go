package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	part := os.Getenv("part")
	if part == "part2" {
		lanterns(256)
	} else {
		lanterns(80)
	}
}

func lanterns(days int) {
	data := readInput("input.txt")

	var queue []int
	queue = append(queue, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}...)

	for _, fish := range data {
		queue[fish] += 1
	}

	cycleDay := 0
	for day := 1; day <= days; day++ {
		fishes := queue[0]
		queue = queue[1:]
		queue = append(queue, fishes)
		queue[6] += fishes

		if cycleDay == 8 {
			cycleDay = 0
		} else {
			cycleDay++
		}
	}
	sum := 0
	for _, element := range queue {
		sum += element
	}
	println(sum)
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
