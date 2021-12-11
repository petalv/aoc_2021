package main

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
)

type Coord struct {
	x   int
	y   int
	val int
}

func main() {
	part := os.Getenv("part")
	matrix := readInput("input.txt")

	if part == "part2" {
		println(part2(matrix))
	} else {
		sum := 0
		for _, c := range part1(matrix) {
			sum += c.val + 1
		}
		println(sum)
	}
}

func part1(matrix [][]int) []Coord {
	var visited []Coord
	var zap []Coord
	visit := []Coord{{0, 0, matrix[0][0]}}
	directions := []string{"W", "S", "E", "N"}
	var currentCoord Coord
	for len(visit) != 0 {
		currentCoord, visit = visit[0], visit[1:]
		visited = append(visited, currentCoord)
		for _, d := range directions {
			visitCoord, zapCoord, err := check(d, currentCoord, visited, matrix)
			if err == nil {
				if !contains(visit, visitCoord) {
					visit = append(visit, visitCoord)
				}
				if !contains(zap, zapCoord) {
					zap = append(zap, zapCoord)
				}
			}
		}
	}
	return getLowPoints(visited, zap)
}

func part2(matrix [][]int) int {
	lowPoints := part1(matrix)
	var visit []Coord
	var visited []Coord
	var currentCoord Coord
	directions := []string{"W", "S", "E", "N"}
	var basinSum []int
	for _, c := range lowPoints {
		keep := []Coord{c}
		visit = append(visit, c)
		for {
			if len(visit) == 0 {
				break
			}
			currentCoord, visit = visit[0], visit[1:]
			visited = append(visited, currentCoord)
			for _, d := range directions {
				visitCoord, zapCoord, err := check(d, currentCoord, visited, matrix)
				if err == nil && !equals(zapCoord, currentCoord) && visitCoord.val != 9 {
					if !contains(visit, visitCoord) {
						visit = append(visit, visitCoord)
					}

					if !contains(keep, zapCoord) {
						keep = append(keep, zapCoord)
					}
				}
			}
		}
		basinSum = append(basinSum, len(keep))
	}
	sort.Ints(basinSum)
	return basinSum[len(basinSum)-1] * basinSum[len(basinSum)-2] * basinSum[len(basinSum)-3]
}

func getLowPoints(visited []Coord, zap []Coord) []Coord {
	var lows []Coord
	for _, c := range visited {
		if !contains(zap, c) {
			lows = append(lows, c)
		}
	}
	return lows
}

func check(d string, coord Coord, visited []Coord, matrix [][]int) (Coord, Coord, error) {
	s, err := getCoord(d, coord.x, coord.y, matrix)
	if err != nil {
		return Coord{}, Coord{}, err
	}
	if contains(visited, s) {
		return Coord{}, Coord{}, errors.New("AlreadyVisited")
	}
	if coord.val < s.val {
		return s, s, nil
	} else {
		return s, coord, nil
	}
}

func getCoord(d string, x int, y int, matrix [][]int) (Coord, error) {
	var nx, ny int
	switch d {
	case "W":
		nx, ny = x-1, y
		break
	case "S":
		nx, ny = x, y+1
		break
	case "E":
		nx, ny = x+1, y
		break
	case "N":
		nx, ny = x, y-1
		break
	default:
		return Coord{}, errors.New("UnknownArgument " + d)
	}
	val, err := getValAt(nx, ny, matrix)
	return Coord{nx, ny, val}, err
}

func getValAt(x int, y int, matrix [][]int) (int, error) {
	if x >= 0 && y >= 0 && y < len(matrix) && x < len(matrix[y]) {
		return matrix[y][x], nil
	}
	return -1, errors.New("OutOfBounds")
}

func contains(slice []Coord, item Coord) bool {
	return indexOf(slice, item) != -1
}

func indexOf(slice []Coord, item Coord) int {
	for i, coord := range slice {
		if equals(coord, item) {
			return i
		}
	}
	return -1
}

func equals(a Coord, b Coord) bool {
	return a.x == b.x && a.y == b.y
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix = make([][]int, 100)
	y := 0
	for scanner.Scan() {
		values := extractNumbers(scanner.Text())
		matrix[y] = values
		y++
	}
	return matrix
}

func extractNumbers(line string) []int {
	var numbers []int
	for idx := range line {
		numVal, _ := strconv.Atoi(string(line[idx]))
		numbers = append(numbers, numVal)
	}
	return numbers
}
