package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	part := os.Getenv("part")
	data := readInput("input.txt")
	if part == "part2" {
		r := part2(data)
		println(r[len(r)/2])
	} else {
		println(part1(data))
	}
}

func part1(data []string) int {
	sum := 0
	for _, line := range data {
		err := parseLine(line)
		if err != nil {
			se, seOk := err.(*SyntaxError)
			if seOk {
				sum += score1(se.token)
			}
		}
	}
	return sum
}

func part2(data []string) []int {
	var sum []int
	for _, line := range data {
		err := parseLine(line)
		if err != nil {
			se, seOk := err.(*IncompleteError)
			if seOk {
				sum = append(sum, calcScorePart2(se.completions))
			}
		}
	}
	sort.Ints(sum)
	return sum
}

func score1(t int) int {
	switch t {
	case RIGHT_PARAN:
		return 3
	case RIGHT_BRACKET:
		return 57
	case RIGHT_CURLY:
		return 1197
	default:
		return 25137
	}
}

func calcScorePart2(completions []int) int {
	sum := 0
	for _, t := range completions {
		sum = sum * 5
		sum += score2(t - 1)
	}
	return sum
}

func score2(t int) int {
	switch t {
	case RIGHT_PARAN:
		return 1
	case RIGHT_BRACKET:
		return 2
	case RIGHT_CURLY:
		return 3
	default:
		return 4
	}
}

const RIGHT_PARAN = 1
const LEFT_PARAN = 2
const RIGHT_BRACKET = 3
const LEFT_BRACKET = 4
const RIGHT_CURLY = 5
const LEFT_CURLY = 6
const RIGHT_ANGLE = 7
const LEFT_ANGLE = 8

const CHAR_RIGHT_PARAN = ')'
const CHAR_LEFT_PARAN = '('
const CHAR_RIGHT_BRACKET = ']'
const CHAR_LEFT_BRACKET = '['
const CHAR_RIGHT_CURLY = '}'
const CHAR_LEFT_CURLY = '{'
const CHAR_RIGHT_ANGLE = '>'
const CHAR_LEFT_ANGLE = '<'

type SyntaxError struct {
	pos   int
	token int
}

type IncompleteError struct {
	completions []int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("Error %d: %s not allowed", e.pos, e.token)
}

func (e *IncompleteError) Error() string {
	return fmt.Sprintf("Incomplete, missing %v", e.completions)
}

func parseLine(line string) error {
	var state []int
	var currentState int
	for pos, c := range string(line) {
		t := toToken(c)
		currentState, state = pop(state)
		a := allowed(t, currentState)
		if !a {
			return &SyntaxError{
				pos:   pos,
				token: t,
			}
		}

		if !isClose(t) {
			if currentState != -1 {
				state = append(state, currentState)
			}
			state = append(state, t)
		}
	}

	if len(state) != 0 {
		return &IncompleteError{
			completions: reverse(state),
		}
	}
	return nil
}

func allowed(t int, currentState int) bool {
	if currentState == -1 {
		return isOpen(t)
	} else if currentState == LEFT_BRACKET {
		if isOpen(t) || t == RIGHT_BRACKET {
			return true
		}
	} else if currentState == LEFT_PARAN {
		if isOpen(t) || t == RIGHT_PARAN {
			return true
		}
	} else if currentState == LEFT_CURLY {
		if isOpen(t) || t == RIGHT_CURLY {
			return true
		}
	} else if currentState == LEFT_ANGLE {
		if isOpen(t) || t == RIGHT_ANGLE {
			return true
		}
	}
	return false
}

func isOpen(t int) bool {
	return t == LEFT_PARAN || t == LEFT_BRACKET || t == LEFT_CURLY || t == LEFT_ANGLE
}

func isClose(t int) bool {
	return t == RIGHT_PARAN || t == RIGHT_BRACKET || t == RIGHT_CURLY || t == RIGHT_ANGLE
}

func toToken(r rune) int {
	switch r {
	case CHAR_RIGHT_PARAN:
		return RIGHT_PARAN
	case CHAR_LEFT_PARAN:
		return LEFT_PARAN
	case CHAR_RIGHT_BRACKET:
		return RIGHT_BRACKET
	case CHAR_LEFT_BRACKET:
		return LEFT_BRACKET
	case CHAR_RIGHT_CURLY:
		return RIGHT_CURLY
	case CHAR_LEFT_CURLY:
		return LEFT_CURLY
	case CHAR_RIGHT_ANGLE:
		return RIGHT_ANGLE
	case CHAR_LEFT_ANGLE:
		return LEFT_ANGLE
	default:
		return -1
	}
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func pop(a []int) (int, []int) {
	if len(a) == 0 {
		return -1, a
	}
	return a[len(a)-1], a[:len(a)-1]
}
