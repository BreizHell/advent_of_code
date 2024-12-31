package main

import (
	"bufio"
	"log"
	"os"
)

type coordinate struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	originalMatrix := make([][]rune, 0, 131)
	var posX = 0
	var posY = 0

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		originalMatrix = append(originalMatrix, []rune(text))

		for x, tile := range text {
			if tile == '^' || tile == '>' || tile == 'v' || tile == '<' {
				posX, posY = x, y
			}
		}
	}

	matrixBuffer := make([][]rune, len(originalMatrix))
	cloneMatrix(originalMatrix, matrixBuffer)

	originalPathCoordinates, _ := play(matrixBuffer, posX, posY)
	delete(originalPathCoordinates, coordinate{posX, posY}) // remove starting point

	nbSolutions := 0
	for coord := range originalPathCoordinates {
		cloneMatrix(originalMatrix, matrixBuffer)
		matrixBuffer[coord.y][coord.x] = '#'
		if _, hasLooped := play(matrixBuffer, posX, posY); hasLooped {
			nbSolutions++
			log.Printf("current nbSolutions: %v", nbSolutions)
		}
	}

	log.Printf("Solutions: %v", nbSolutions)
}

func play(matrix [][]rune, startX int, startY int) (map[coordinate]bool, bool) {
	path := make(map[coordinate]bool, len(matrix)*len(matrix[0]))

	x, y := startX, startY
	var nextTile rune
	var nextCoord coordinate
	var currentDirection rune
	var stillWithinBoundary = true
	defer func() {
		path[coordinate{x, y}] = true
		matrix[y][x] = currentDirection
	}()

	for {
		defer func(fromX, fromY int) {
			path[coordinate{fromX, fromY}] = true
		}(x, y)

		// TODO	: redo this chaos vvvv
		currentDirection = matrix[y][x]
		nextCoord, stillWithinBoundary = nextTileCoordinates(matrix, currentDirection, x, y)
		x = nextCoord.x
		y = nextCoord.y

		if !stillWithinBoundary
		//y--
		//	if y-1 < 0 {
		//		return path, false
		//	}
		//	nextTile = matrix[y-1][x]

		if !stillWithinBoundary

		// TODO	: redo this chaos ^^^^

		if nextTile == '#' {
			matrix[y][x] = ninetyDegreeClockwise(currentDirection)
			nextCoord, stillWithinBoundary = nextTileCoordinates(matrix, matrix[y][x], x, y)
			if stillWithinBoundary && matrix[nextCoord.y][nextCoord.x] == matrix[y][x] {
				return path, true
			}
		} else {
			matrix[y][x] = currentDirection
			if matrix[y][x] == nextTile {
				return path, true
			}
		}
	}
}

func cloneMatrix(source [][]rune, destination [][]rune) {
	for y, row := range source {
		if destination[y] == nil {
			destination[y] = make([]rune, len(row))
		}
		copy(destination[y], row)
	}
}

func ninetyDegreeClockwise(from rune) rune {
	switch from {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		log.Fatalf("Unkown direction '%v'", from)
		return 'X'
	}
}

func nextTileCoordinates(matrix [][]rune, direction rune, x, y int) (coord coordinate, respectsBoundary bool) {
	switch direction {
	case '^':
		if 0 < y-1 {
			respectsBoundary = true
		}
		coord = coordinate{x, y - 1}
	case '>':
		if x < len(matrix[y])-1 {
			respectsBoundary = true
		}
		coord = coordinate{x + 1, y}
	case 'v':
		if y < len(matrix)-1 {
			respectsBoundary = true
		}
		coord = coordinate{x, y + 1}
	case '<':
		if 0 < x-1 {
			respectsBoundary = true
		}
		coord = coordinate{x - 1, y}
	default:
		log.Fatalf("Unkown direction '%v'", direction)
		coord = coordinate{0, 0}
	}
	return coord, respectsBoundary
}
