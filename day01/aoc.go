package main

import (
	"fmt"
	"os"
	"petalv/aoc_2021/file"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadIntArray("input.txt")
	if part == "part2" {
		fmt.Println(checkDepthWindow(-1, input))
	} else {
		fmt.Println(checkDepth(input))
	}
}

func checkDepth(depths []int) int {
	if len(depths) == 1 {
		return 0
	}
	head, tail := depths[0], depths[1:]
	if tail[0] > head {
		return checkDepth(tail) + 1
	} else {
		return checkDepth(tail)
	}
}

func checkDepthWindow(prevWindow int, depths []int) int {
	if len(depths) == 2 {
		return 0
	}
	a, b, c, tail := depths[0], depths[1], depths[2], depths[1:]
	window := a + b + c
	if prevWindow != -1 && window > prevWindow {
		return checkDepthWindow(window, tail) + 1
	} else {
		return checkDepthWindow(window, tail)
	}

}
