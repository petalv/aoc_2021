package main

import (
	"os"
	"petalv/aoc_2021/file"
	"strconv"
)

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStrArray("input.txt")
	// still wip
	if part == "part2" {
		h20 := toDec(part2(12, input, false))
		co2 := toDec(part2(12, input, true))
		println(h20 * co2)
	} else {
		bitSize := len(input[0])
		bits := significant(bitSize, input)
		strBits := ""
		for _, bit := range bits {
			strBits += strconv.Itoa(bit)
		}
		gamma, _ := strconv.ParseUint(strBits, 2, bitSize)
		epsilon := (^gamma) & 0xFFF
		println(gamma * epsilon)
	}
}

func significant(bitSize int, input []string) []int {
	storage := [16]int{}
	slice := storage[0:bitSize]
	for _, line := range input {
		u64, _ := strconv.ParseUint(line, 2, bitSize)
		for idx := range slice {
			if (u64>>idx)&1 != 0 {
				slice[bitSize-1-idx]++
			} else {
				slice[bitSize-1-idx]--
			}
		}
	}
	var bits []int
	for i := range slice {
		bits = append(bits, 1+slice[i]>>31)
	}
	return bits
}

func part2(bitSize int, input []string, flip bool) []string {
	var sublist = input
	for i := 0; i < bitSize; i++ {
		si := significant(bitSize, sublist)[i]
		if flip {
			si = si<<0 ^ 0x1
		}
		sublist = startswith(strconv.Itoa(si), i, sublist)
		if len(sublist) == 1 {
			break
		}
	}
	return sublist
}

func startswith(i string, pos int, lines []string) []string {
	var sublist []string
	for _, line := range lines {
		if line[pos:pos+1] == i {
			sublist = append(sublist, line)
		}
	}
	return sublist
}

func toDec(bits []string) uint64 {
	strBits := ""
	for _, bit := range bits {
		strBits += bit
	}
	res, _ := strconv.ParseUint(strBits, 2, 16)
	return res
}
