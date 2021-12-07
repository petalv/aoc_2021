package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	id  int
	sum int
}

func main() {
	part := os.Getenv("part")
	var calledNumbers, boards, _ = ReadBingoInput("input.txt")
	if part == "part2" {
		finishedBoards := playGame(calledNumbers, boards, true)
		println(finishedBoards[len(finishedBoards)-1].sum)
	} else {
		finishedBoards := playGame(calledNumbers, boards, false)
		println(finishedBoards[0].sum)
	}
}

func playGame(calledNumbers []int, boards map[int][]int, runAll bool) map[int]Bingo {
	var boardOverlays = make(map[int][]int)
	for boardId := range boards {
		boardOverlays[boardId] = make([]int, 25)
	}
	var finishedBoards = make(map[int]Bingo)
	winOrder := 0
	// This is not nice :)
out:
	for _, calledNumber := range calledNumbers {
		for boardId := range boards {
			markIndex := markNumber(calledNumber, boards[boardId], boardOverlays[boardId])
			if markIndex != -1 {
				if checkBingo(boardOverlays[boardId]) {
					sum := sumBoard(boards[boardId], boardOverlays[boardId])
					finishedBoards[winOrder] = Bingo{id: boardId, sum: sum * calledNumber}
					if !runAll {
						break out
					}
					delete(boards, boardId)
					winOrder++

				}
			}
		}
	}
	return finishedBoards
}

func sumBoard(board []int, boardOverlay []int) int {
	sum := 0
	for i := 0; i <= 24; i++ {
		if boardOverlay[i] != 1 {
			sum += board[i]
		}
	}
	return sum
}

func markNumber(number int, board []int, boardOverlay []int) int {
	found := Find(board, number)
	if found != -1 {
		boardOverlay[found] = 1
		return found
	}
	return -1
}

func checkBingo(board []int) bool {
	return checkBoardRow(board) || checkBoardCols(board)
}

func checkBoardRow(board []int) bool {
	return sumSlice(board[0:5]) == 5 ||
		sumSlice(board[5:10]) == 5 ||
		sumSlice(board[10:15]) == 5 ||
		sumSlice(board[15:20]) == 5 ||
		sumSlice(board[20:25]) == 5
}

func checkBoardCols(board []int) bool {
	colSum := [5]int{0, 0, 0, 0, 0}
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			colSum[i] += board[j*5+i]
			if colSum[i] == 5 {
				return true
			}
		}
	}
	return false
}

func sumSlice(slice []int) int {
	total := 0
	for _, val := range slice {
		total += val
	}
	return total
}

func Find(haystack []int, needle int) int {
	for i, n := range haystack {
		if needle == n {
			return i
		}
	}
	return -1
}

func ReadBingoInput(path string) ([]int, map[int][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNo := 0
	var calledNumbers []int
	var board []int
	var boards = make(map[int][]int)
	for scanner.Scan() {
		if lineNo == 0 {
			calledNumbers = extractNumbers(scanner.Text(), ",")
		} else if scanner.Text() == "" {
			if board != nil {
				boards[len(boards)] = board
			}
			board = []int{}
		} else {
			board = append(board, extractBoardLine(scanner.Text())...)
		}
		lineNo++
	}
	boards[len(boards)] = board
	return calledNumbers, boards, scanner.Err()
}

func extractNumbers(line string, sep string) []int {
	var numbers []int
	for _, val := range strings.Split(line, sep) {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	return numbers
}

func extractBoardLine(line string) []int {
	var numbers []int
	for _, val := range strings.Fields(line) {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	return numbers
}
