package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Squid struct {
	point   Point
	energy  int
	flashed bool
}

type Heading int

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func main() {
	part := os.Getenv("part")
	data := readInput("input.txt")
	if part == "part2" {
		println(part1(data, true))
	} else {
		println(part1(data, false))
	}
}

func part1(data [][]int, sync bool) int {
	//fmt.Printf("data: %v", data)

	matrix := &matrixx{
		matrix: data,
		ymax:   len(data),
		xmax:   len(data[0]),
	}

	var flashers []Squid
	numTicks := 100
	if sync {
		numTicks = int(^uint(0) >> 1)
	}
	sum := 0
	for ticks := 0; ticks < numTicks; ticks++ {
		flashers = accumulate(matrix)
		newFlashers := animation(flashers, *matrix)
		flashers = append(flashers, newFlashers...)
		for len(newFlashers) != 0 {
			newFlashers = animation(newFlashers, *matrix)
			flashers = append(flashers, newFlashers...)
		}
		if sync && len(flashers) == matrix.ymax*matrix.xmax {
			return ticks + 1
		}
		sum += len(flashers)
	}
	return sum
}

func animation(flashers []Squid, m matrixx) []Squid {
	var lightUp []Squid
	var buddies []Squid
	for _, flasher := range flashers {
		buddies = getBordering(flasher, m)
		for _, buddy := range buddies {
			if !contains(buddy, lightUp) {
				if buddy.energy != 0 {
					increment(&buddy)
					setVal(&buddy, m.matrix)
					if buddy.energy == 0 {
						lightUp = append(lightUp, buddy)
					}
				}
			}

		}
	}
	return lightUp
}

func contains(s Squid, squids []Squid) bool {
	for _, x := range squids {
		if x.point.y == s.point.y && x.point.x == s.point.x {
			return true
		}
	}
	return false
}

func accumulate(m *matrixx) []Squid {
	iterator := m.createIterator()
	var lightUp []Squid
	for iterator.hasNext() {
		p, _ := iterator.getNext()
		increment(p)
		setVal(p, m.matrix)
		if p.energy == 0 {
			lightUp = append(lightUp, *p)
		}
	}
	return lightUp
}

func increment(p *Squid) {
	if p.energy == 9 {
		p.energy = 0
	} else {
		p.energy++
	}
}

func setVal(s *Squid, matrix [][]int) {
	matrix[s.point.y][s.point.x] = s.energy
}

func isValid(p Point, m matrixx) bool {
	return p.x >= 0 && p.y >= 0 && p.y < m.ymax && p.x < m.xmax
}

func getBordering(to Squid, m matrixx) []Squid {
	var positions []Point
	positions = []Point{getOffset(N, to),
		getOffset(NE, to),
		getOffset(E, to),
		getOffset(SE, to),
		getOffset(S, to),
		getOffset(SW, to),
		getOffset(W, to),
		getOffset(NW, to)}

	var squids []Squid
	for _, pos := range positions {
		if isValid(pos, m) {
			squids = append(squids, Squid{point: pos, energy: m.matrix[pos.y][pos.x]})
		}
	}
	return squids
}

func getOffset(heading Heading, coord Squid) Point {
	ox, oy := offset(heading)
	return Point{x: coord.point.x + ox, y: coord.point.y + oy}
}

func offset(heading Heading) (int, int) {
	switch heading {
	case N:
		return 0, -1
	case NE:
		return 1, -1
	case E:
		return 1, 0
	case SE:
		return 1, 1
	case S:
		return 0, 1
	case SW:
		return -1, 1
	case W:
		return -1, 0
	case NW:
		return -1, -1
	default:
		panic("Never.")
	}
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix = make([][]int, 10, 10)
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

type matrixIterator struct {
	yIndex int
	xIndex int
	matrix [][]int
}

type iterator interface {
	hasNext() bool
	getNext() (*Squid, error)
}

func (m *matrixIterator) hasNext() bool {
	return m.yIndex < len(m.matrix)
}

func (m *matrixIterator) getNext() (*Squid, error) {
	var val *Squid
	if m.hasNext() {
		if m.xIndex < len(m.matrix[m.yIndex]) {
			pos := Point{x: m.xIndex, y: m.yIndex}
			val = &Squid{point: pos, energy: m.matrix[m.yIndex][m.xIndex]}
			if m.xIndex+1 == len(m.matrix[m.yIndex]) {
				m.xIndex = 0
				m.yIndex++
			} else {
				m.xIndex++
			}
		}
		return val, nil
	}
	return nil, fmt.Errorf("OutOfBounds")
}

type matrixx struct {
	matrix [][]int
	ymax   int
	xmax   int
}

func (m *matrixx) createIterator() iterator {
	return &matrixIterator{
		matrix: m.matrix,
	}
}
