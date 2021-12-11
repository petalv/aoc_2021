package main

import (
	"bufio"
	"errors"
	"fmt"
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

func part2(matrix [][]int) int {
	lowPoints := part1(matrix)
	var visit []Coord
	var visited []Coord
	var currentCoord Coord
	directions := []string{"W", "S", "E", "N"}
	var keepSum []int
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
		keepSum = append(keepSum, len(keep))
	}
	sort.Ints(keepSum)
	return keepSum[len(keepSum)-1] * keepSum[len(keepSum)-2] * keepSum[len(keepSum)-3]
}

func part1(matrix [][]int) []Coord {
	var visited []Coord
	var zap []Coord
	visit := []Coord{Coord{0, 0, matrix[0][0]}}
	directions := []string{"W", "S", "E", "N"}
	var currentCoord Coord
	for {
		if len(visit) == 0 {
			break
		}
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

	var sc []string
	for _, c := range zap {
		sc = append(sc, fmt.Sprintf("%d%d", c.x, c.y))
	}
	//sort.Strings(sc)
	var vc []string
	for _, c := range visited {
		vc = append(vc, fmt.Sprintf("%d%d", c.x, c.y))
	}
	//sort.Strings(vc)
	// fmt.Printf("Zapped %v\n", strings.Join(sc, ","))
	// fmt.Printf("visited %v\n", strings.Join(vc, ","))
	return getLowPoints(visited, zap)
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
	s := getCoord(d, coord.x, coord.y, matrix)
	if s.x < 0 || s.y < 0 || s.y >= len(matrix) || s.x >= len(matrix[s.y]) {
		return Coord{}, Coord{}, errors.New("OutofBounds")
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

func getCoord(d string, x int, y int, matrix [][]int) Coord {
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
		panic("bad heading")
	}
	val := getValAt(nx, ny, matrix)
	return Coord{nx, ny, val}
}

func getValAt(x int, y int, matrix [][]int) int {
	if x >= 0 && y >= 0 && y < len(matrix) && x < len(matrix[y]) {
		return matrix[y][x]
	}
	return -1
}

func contains(slice []Coord, item Coord) bool {
	return indexOf(slice, item) != -1
}

func indexOf(slice []Coord, item Coord) int {
	for i, coord := range slice {
		if coord.x == item.x && coord.y == item.y {
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
	for idx, _ := range line {
		numVal, _ := strconv.Atoi(string(line[idx]))
		numbers = append(numbers, numVal)
	}
	return numbers
}
