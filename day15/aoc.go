package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x   int
	y   int
	val int
}

type Node struct {
	parent *Node
	coord  *Point
	f      int
	g      int
	h      int
}

func main() {
	part := os.Getenv("part")
	data := readInput("input.txt", 100, 100)
	if part == "part2" {
		mapWidth := len(data[0])
		mapHeight := len(data)
		tiles := 5
		start := &Point{x: 0, y: 0, val: data[0][0]}

		goalX := mapWidth*tiles - 1
		goalY := mapHeight*tiles - 1
		goalVal := getData(data, goalY, goalX, tiles, tiles, tiles)
		goal := &Point{x: goalX, y: goalY, val: goalVal}
		f := func(a *Point, goal *Point) int {
			return int(math.Abs(float64(a.x-goal.x)) + math.Abs(float64(a.y-goal.y)))
		}
		res, _ := findPath(tiles, data, start, goal, mapWidth, mapHeight, f)
		println(res)
	} else {
		mapWidth := len(data[0])
		mapHeight := len(data)
		tiles := 1
		start := &Point{x: 0, y: 0, val: data[0][0]}
		goal := &Point{x: mapWidth*tiles - 1, y: mapHeight*tiles - 1, val: data[mapWidth*tiles-1][mapHeight*tiles-1]}
		f := func(a *Point, goal *Point) int {
			return int(math.Abs(float64(a.x-goal.x)) + math.Abs(float64(a.y-goal.y)))
		}
		res, node := findPath(tiles, data, start, goal, mapWidth, mapHeight, f)
		println(res)
		fmt.Printf("%s, %s", node.coord.x, node.coord.y)
	}
}

func findPath(tiles int, data [][]int, startCoord *Point, goalCoord *Point, mapWidth int, mapHeight int, fCalculateCost func(*Point, *Point) int) (int, *Node) {
	highscore := 99999999999999
	var openList []*Node
	var closedList []*Node
	openList = append(openList, &Node{parent: nil, coord: startCoord, f: 19999})
	var endNode *Node
	openListIndex := make(map[Point]*Node)
	closedListIndex := make(map[Point]*Node)
	for len(openList) > 0 {
		var nextNode *Node
		nextNode, openList = popBestOpenNode(openList, openListIndex)
		successors := generateSuccessors(tiles, data, nextNode, mapWidth, mapHeight)
		goalNode := checkForGoalState(successors, goalCoord)
		if goalNode == nil {
			for _, successor := range processSuccessors(successors, goalCoord, fCalculateCost) {

				if some(openListIndex, successor) {
					continue
				}

				if some(closedListIndex, successor) {
					continue
				}

				openList = append(openList, successor)
				openListIndex[*successor.coord] = successor
			}
			closedList = append(closedList, nextNode)
			closedListIndex[*nextNode.coord] = nextNode
		} else {
			res := backTrackPath(goalNode)
			sum := 0 - res[0].coord.val
			for _, val := range res {
				sum += val.coord.val
			}
			if sum < highscore {
				highscore = sum
				endNode = res[len(res)-1]
			}
		}
	}
	return highscore, endNode
}

func some(index map[Point]*Node, successor *Node) bool {
	/*for _, openNode := range list {
		if openNode.coord.x == successor.coord.x && openNode.coord.y == successor.coord.y && openNode.f <= successor.f {
			return true
		}
	}*/
	return index[*successor.coord] != nil && index[*successor.coord].f <= successor.f
}

func processSuccessors(successors []*Node, goalCoord *Point, fCalculateHCost func(*Point, *Point) int) []*Node {
	for sIdx := 0; sIdx < len(successors); sIdx++ {
		var successor = successors[sIdx]
		successor.g = successor.parent.g + successor.coord.val
		successor.h = fCalculateHCost(successor.coord, goalCoord)
		successor.f = successor.g + successor.h
	}
	return successors
}

func generateSuccessors(tiles int, data [][]int, node *Node, width int, height int) []*Node {
	nY := node.coord.y - 1
	sY := node.coord.y + 1
	wX := node.coord.x - 1
	eX := node.coord.x + 1

	tileY := getTile(height, node.coord.y)
	tileX := getTile(height, node.coord.x)

	tileNY := getTile(height, nY)
	tileSY := getTile(height, sY)
	tileEX := getTile(width, eX)
	tileWX := getTile(width, wX)

	var successorList []*Node
	if nY >= 0 {
		successorList = append(successorList, &Node{parent: node, coord: &Point{x: node.coord.x, y: nY, val: getData(data, nY, node.coord.x, tileX, tileNY, tiles)}})
	}
	if sY < height*tiles {
		successorList = append(successorList, &Node{parent: node, coord: &Point{x: node.coord.x, y: sY, val: getData(data, sY, node.coord.x, tileX, tileSY, tiles)}})
	}
	if eX < width*tiles {
		successorList = append(successorList, &Node{parent: node, coord: &Point{x: eX, y: node.coord.y, val: getData(data, node.coord.y, eX, tileEX, tileY, tiles)}})
	}
	if wX >= 0 {
		successorList = append(successorList, &Node{parent: node, coord: &Point{x: wX, y: node.coord.y, val: getData(data, node.coord.y, wX, tileWX, tileY, tiles)}})
	}
	return successorList
}

func getTile(max int, val int) int {
	return int(math.Floor(float64(val/max + 1)))
}

func checkForGoalState(nodes []*Node, goalCoord *Point) *Node {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].coord.x == goalCoord.x && nodes[i].coord.y == goalCoord.y {
			return nodes[i]
		}
	}
	return nil
}
func getData(data [][]int, y int, x int, tileX int, tileY int, tiles int) int {
	if y < 0 {
		return 99999
	}
	if x < 0 {
		return 99999
	}

	if y > len(data)*tiles {
		return 99999
	}

	if x > len(data[0])*tiles {
		return 99999
	}
	valX := int(math.Max(0, float64(x-((tileX-1)*len(data[0])-len(data[0]))-len(data[0]))))
	valY := int(math.Max(0, float64(y-((tileY-1)*len(data)-len(data))-len(data))))
	val := data[valY][valX]

	if tileY > 1 {
		val = val + tileY - 1
		if val > 9 {
			val = val - 9
		}
	}
	if tileX > 1 {
		val = val + tileX - 1
		if val > 9 {
			val = val - 9
		}
	}

	return val
}

func popBestOpenNode(openList []*Node, openListIndex map[Point]*Node) (*Node, []*Node) {
	bestNodeIdx := -1
	for i := 0; i < len(openList); i++ {
		if bestNodeIdx == -1 {
			bestNodeIdx = i
		} else if openList[i].f < openList[bestNodeIdx].f {
			bestNodeIdx = i
		}
	}

	node := openList[bestNodeIdx]
	openListIndex[*node.coord] = node
	openList = remIndx(openList, bestNodeIdx)
	return node, openList
}

func remIndx(list []*Node, i int) []*Node {
	var nl []*Node
	for idx, val := range list {
		if i != idx {
			nl = append(nl, val)
		}
	}
	return nl
}

func backTrackPath(node *Node) []*Node {
	if node.parent == nil {
		return []*Node{node}
	}
	return append(backTrackPath(node.parent), node)
}

func readInput(path string, width int, height int) [][]int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix = make([][]int, height, width)
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
