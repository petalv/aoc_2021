package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Fold struct {
	x int
	y int
}

func main() {
	part := os.Getenv("part")
	intX, intY, folds := readInput("input.txt")
	if part == "part2" {
		folder(intX, intY, folds)
	} else {
		println(folder(intX, intY, []Fold{folds[0]}))
	}
}

func folder(intX []int, intY []int, folds []Fold) int {
	_, xmax := MinMax(intX)
	_, ymax := MinMax(intY)
	var data = make([][]int, ymax+1)
	for idx := range data {
		var xs []int
		for i := 0; i < xmax+1; i++ {
			xs = append(xs, 0)
		}
		data[idx] = xs
	}
	for idx := range intX {
		data[intY[idx]][intX[idx]] = 1
	}

	matrixx := &matrixx{
		matrix: data,
		ymax:   ymax + 1,
		xmax:   xmax + 1,
	}

	for _, fold := range folds {
		if fold.y != 0 {
			foldUp(fold.y, matrixx)
		} else {
			foldLeft(fold.x, matrixx)
		}

	}
	if len(folds) > 1 {
		printMatrix(matrixx)
	}

	sum := 0
	for y := 0; y < matrixx.ymax; y++ {
		for x := 0; x < matrixx.xmax; x++ {
			sum += matrixx.matrix[y][x]
		}
	}
	return sum
}

func foldLeft(foldX int, matrixx *matrixx) {
	for x := foldX; x < matrixx.xmax; x++ {
		for y := 0; y < matrixx.ymax; y++ {
			if matrixx.matrix[y][x] == 1 {
				newX := foldX - x + foldX
				//fmt.Printf("move %d,%d to %d,%d \n", y, x, y, newX)
				matrixx.matrix[y][newX] = 1
			}
		}
	}
	matrixx.xmax = foldX
}

func foldUp(foldY int, matrixx *matrixx) {
	for i := foldY; i < matrixx.ymax; i++ {
		for x := 0; x < matrixx.xmax; x++ {
			if matrixx.matrix[i][x] == 1 {
				newY := foldY - i + foldY
				//fmt.Printf("move %d,%d to %d,%d \n", x, i, x, newY )
				matrixx.matrix[newY][x] = 1
			}
		}
	}
	matrixx.ymax = foldY
}

func printMatrix(matrixx *matrixx) {
	for y := 0; y < matrixx.ymax; y++ {
		line := ""
		for x := 0; x < matrixx.xmax; x++ {
			line += strconv.Itoa(matrixx.matrix[y][x])
		}
		println(strings.ReplaceAll(strings.ReplaceAll(line, "0", "."), "1", "#"))
	}
}
func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

type matrixx struct {
	matrix [][]int
	ymax   int
	xmax   int
}

func readInput(path string) ([]int, []int, []Fold) {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var valuesX []int
	var valuesY []int
	var folds []Fold
	isXY := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isXY = false
		} else {
			if isXY {
				parts := strings.Split(line, ",")
				numValX, _ := strconv.Atoi(parts[0])
				numValY, _ := strconv.Atoi(parts[1])
				valuesX = append(valuesX, numValX)
				valuesY = append(valuesY, numValY)
			} else {
				foldLine := strings.ReplaceAll(line, "fold along ", "")
				foldParts := strings.Split(foldLine, "=")
				val, _ := strconv.Atoi(foldParts[1])
				if foldParts[0] == "x" {
					folds = append(folds, Fold{x: val})
				} else {
					folds = append(folds, Fold{y: val})
				}
			}
		}
	}
	return valuesX, valuesY, folds
}
