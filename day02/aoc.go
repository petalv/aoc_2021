package main

import (
	"fmt"
	"os"
	"petalv/aoc_2021/file"
	"strconv"
	"strings"
	"sync"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStrArray("input.txt")
	if part == "part2" {
		// fmt.Println(checkDepthWindow(-1, input))
	} else {
		fmt.Println(part1(input))
	}
}

type Movement struct {
	Direction string
	Step      int
}

type Position struct {
	Depth int
	Reach int
}

func part1(moves []string) int {
	mappedMovements := make(chan []Movement)
	finalPosition := make(chan Position)
	var wg sync.WaitGroup
	wg.Add(len(moves))

	for _, line := range moves {
		go func(move []string) {
			defer wg.Done()
			mappedMovements <- Map(move)
		}(strings.Split(line, " "))
	}

	go Reducer(mappedMovements, finalPosition)
	wg.Wait()
	close(mappedMovements)
	res := <-finalPosition
	return res.Depth * res.Reach
}

func Map(movement []string) []Movement {
	var mappedMovements []Movement
	step, _ := strconv.Atoi(movement[1])
	return append(mappedMovements, Movement{
		Direction: movement[0],
		Step:      step,
	})
}

func Reducer(mappings chan []Movement, pos chan Position) {
	finalPosition := Position{Reach: 0, Depth: 0}
	for list := range mappings {
		for _, value := range list {
			if value.Direction == "forward" {
				finalPosition.Reach += value.Step
			}
			if value.Direction == "down" {
				finalPosition.Depth += value.Step
			}
			if value.Direction == "up" {
				finalPosition.Depth -= value.Step
			}
		}
	}
	pos <- finalPosition
}
