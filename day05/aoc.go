package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func main() {
	part := os.Getenv("part")
	coords := readCoords("input.txt")
	intersectionMap := make(map[string]int)

	for _, coord := range coords {
		if coord.x1 == coord.x2 || coord.y1 == coord.y2 {
			if coord.x1 != coord.x2 {
				for x := min(coord.x1, coord.x2); x <= max(coord.x1, coord.x2); x++ {
					intersectionMap[strconv.Itoa(x)+"-"+strconv.Itoa(coord.y1)] += 1
				}
			}
			if coord.y1 != coord.y2 {
				for y := min(coord.y1, coord.y2); y <= max(coord.y1, coord.y2); y++ {
					intersectionMap[strconv.Itoa(coord.x2)+"-"+strconv.Itoa(y)] += 1
				}
			}
		} else if part == "part2" {
			for step := 0; step <= Abs(coord.x1-coord.x2); step++ {
				if coord.x1 > coord.x2 {
					if coord.y1 > coord.y2 {
						intersectionMap[strconv.Itoa(coord.x1-step)+"-"+strconv.Itoa(coord.y1-step)] += 1
					} else {
						intersectionMap[strconv.Itoa(coord.x1-step)+"-"+strconv.Itoa(coord.y1+step)] += 1
					}
				} else {
					if coord.y1 > coord.y2 {
						intersectionMap[strconv.Itoa(coord.x1+step)+"-"+strconv.Itoa(coord.y1-step)] += 1
					} else {
						intersectionMap[strconv.Itoa(coord.x1+step)+"-"+strconv.Itoa(coord.y1+step)] += 1
					}
				}
			}
		}
	}

	counter := 0
	for _, entry := range intersectionMap {
		if entry >= 2 {
			counter++
		}
	}

	println(counter)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readCoords(path string) []Coord {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var coords []Coord
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		c1parts := strings.Split(parts[0], ",")
		c2parts := strings.Split(parts[1], ",")
		x1 := toInt(c1parts[0])
		x2 := toInt(c2parts[0])
		y1 := toInt(c1parts[1])
		y2 := toInt(c2parts[1])
		coord := Coord{x1: x1, y1: y1, x2: x2, y2: y2}
		coords = append(coords, coord)
	}
	return coords
}

func min(i1 int, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func toInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
