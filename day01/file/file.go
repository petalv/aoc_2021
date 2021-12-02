package file

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}
