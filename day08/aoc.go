package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part := os.Getenv("part")
	displays := readInput("input.txt")
	if part == "part2" {
		println(part2(displays))
	} else {
		println(part1(displays))
	}
}

const ONE = 2
const TWO_THREE_FIVE = 5
const FOUR = 4
const SEVEN = 3
const EIGHT = 7

const DIGIT_ZERO = "abcefg"
const DIGIT_ONE = "cf"
const DIGIT_TWO = "acdeg"
const DIGIT_THREE = "acdfg"
const DIGIT_FOUR = "bcdf"
const DIGIT_FIVE = "abdfg"
const DIGIT_SIX = "abdefg"
const DIGIT_SEVEN = "acf"
const DIGIT_EIGHT = "abcdefg"
const DIGIT_NINE = "abcdfg"

type Display struct {
	patterns []string
	outputs  []string
}

func part1(displays []Display) int {
	sum := 0
	for _, display := range displays {
		for _, output := range display.outputs {
			l := len(output)
			if l == ONE || l == FOUR || l == SEVEN || l == EIGHT {
				sum++
			}
		}
	}
	return sum
}

func part2(displays []Display) int {
	digitSegments := []string{"a", "b", "c", "d", "e", "f", "g"}

	allSum := 0
	for _, display := range displays {
		// Heuristics start
		the1 := SortString(findLen(ONE, display.patterns)[0])
		the7 := SortString(findLen(SEVEN, display.patterns)[0])
		the4 := SortString(findLen(FOUR, display.patterns)[0])
		a := RemoveInString(the7, the1)
		the235 := SortStrings(findLen(TWO_THREE_FIVE, display.patterns))
		the3 := findLen(3, RemoveInStrings(the235, the1))[0]
		dOrG := strings.Replace(the3, a, "", 1)
		bOrD := RemoveInString(the4, the1)
		g := strings.Replace(strings.Replace(dOrG, string(bOrD[0]), "", 1), string(bOrD[1]), "", 1)
		d := strings.Replace(dOrG, g, "", 1)
		b := strings.Replace(bOrD, d, "", 1)
		lookingFor2 := RemoveInStrings(RemoveInStrings(RemoveInStrings(the235, a), d), g)
		lookingFor2 = RemoveInStrings(lookingFor2, string(the1[0]))
		lookingFor2 = RemoveInStrings(lookingFor2, string(the1[1]))
		lookingForE := RemoveInStrings(lookingFor2, b)
		e := lookingForE[0]
		f := findLen(1, RemoveInStrings(RemoveInStrings(RemoveInStrings(RemoveInStrings(the235, a), b), d), g))[0]
		c := strings.Replace(the1, f, "", 1)
		// heuristics ends

		mapping := strings.Join([]string{a, b, c, d, e, f, g}, "")
		outPutSum := ""
		for _, out := range display.outputs {
			res := ""
			for _, o := range strings.Split(out, "") {
				i := strings.Index(mapping, o)
				res += SortString(digitSegments[i])
			}
			outPutSum += strconv.Itoa(matchDigit(SortString(res)))
		}
		s, _ := strconv.Atoi(outPutSum)
		allSum += s
	}
	return allSum
}

func matchDigit(in string) int {
	switch in {
	case DIGIT_ONE:
		return 1
	case DIGIT_TWO:
		return 2
	case DIGIT_THREE:
		return 3
	case DIGIT_FOUR:
		return 4
	case DIGIT_FIVE:
		return 5
	case DIGIT_SIX:
		return 6
	case DIGIT_SEVEN:
		return 7
	case DIGIT_EIGHT:
		return 8
	case DIGIT_NINE:
		return 9
	case DIGIT_ZERO:
		return 0
	default:
		println("PANIC", in)
		return -1
	}
}

func RemoveInString(input string, pattern string) string {
	return RemoveInStrings([]string{input}, pattern)[0]
}

func RemoveInStrings(input []string, pattern string) []string {
	var result []string
	for _, zapMe := range input {
		for _, p := range strings.Split(pattern, "") {
			zapMe = strings.Replace(zapMe, p, "", 1)
		}
		if len(zapMe) != 0 {
			result = append(result, zapMe)
		}
	}
	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func SortStrings(toSort []string) []string {
	var sorted []string
	for _, sortMe := range toSort {
		sorted = append(sorted, SortString(sortMe))
	}
	return sorted
}

func findLen(needle int, patterns []string) []string {
	var founds []string
	for _, pattern := range patterns {
		if len(pattern) == needle {
			founds = append(founds, pattern)
		}
	}
	return founds
}

func readInput(path string) []Display {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var displays []Display
	for scanner.Scan() {
		displaySplits := strings.Split(scanner.Text(), " | ")
		display := Display{patterns: strings.Fields(displaySplits[0]),
			outputs: strings.Fields(displaySplits[1])}
		displays = append(displays, display)
	}
	return displays
}
