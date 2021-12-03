package main

import (
	"fmt"
	"os"
	"petalv/aoc_2021/file"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStrArray("input.txt")
	if part == "part2" {
		fmt.Println(navigate(2, input))
	} else {
		fmt.Println(navigate(1, input))
	}
}

type Movement struct {
	Id        int
	Direction string
	Step      int
}

type Position struct {
	Depth int
	Reach int
	Aim   int
}

func navigate(part int, moves []string) int {
	mappedMovements := make(chan Movement, len(moves))
	finalPosition := make(chan Position)
	var wg sync.WaitGroup
	wg.Add(len(moves))
	for id, line := range moves {
		go func(move string, id int) {
			defer wg.Done()
			mappedMovements <- Map(move, id)
		}(line, id)
	}
	// This got ugly, because needed to order mappings
	if part == 2 {
		wg.Wait()
		go Reducer2(mappedMovements, finalPosition)
	} else {
		go Reducer1(mappedMovements, finalPosition)
		wg.Wait()
	}

	close(mappedMovements)
	res := <-finalPosition
	return res.Depth * res.Reach
}

func Map(rawMovement string, id int) Movement {
	movement := strings.Split(rawMovement, " ")
	step, _ := strconv.Atoi(movement[1])
	return Movement{
		Id:        id,
		Direction: movement[0],
		Step:      step,
	}
}

func Reducer1(movements chan Movement, pos chan Position) {
	finalPosition := Position{Reach: 0, Depth: 0, Aim: 0}
	for value := range movements {
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
	pos <- finalPosition
}

func Reducer2(movements chan Movement, pos chan Position) {
	finalPosition := Position{Reach: 0, Depth: 0, Aim: 0}
	var orderedMoves []Movement
	for value := range movements {
		orderedMoves = append(orderedMoves, value)
	}
	sort.Slice(orderedMoves, func(i, j int) bool {
		return orderedMoves[i].Id < orderedMoves[j].Id
	})
	for _, value := range orderedMoves {
		if value.Direction == "forward" {
			finalPosition.Reach += value.Step
			finalPosition.Depth += finalPosition.Aim * value.Step
		}
		if value.Direction == "down" {
			finalPosition.Aim += value.Step
		}
		if value.Direction == "up" {
			finalPosition.Aim -= value.Step
		}
	}
	pos <- finalPosition
}
