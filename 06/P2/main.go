package main

import (
	"bufio"
	"fmt"
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
	var startX = 0
	var startY = 0

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		originalMatrix = append(originalMatrix, []rune(text))

		for x, tile := range text {
			if tile == '^' || tile == '>' || tile == 'v' || tile == '<' {
				startX, startY = x, y
			}
		}
	}

	matrixBuffer := make([][]rune, len(originalMatrix))
	cloneMatrix(originalMatrix, matrixBuffer)

	originalPathCoordinates, _ := play(matrixBuffer, startX, startY)

	// printMatrix(matrixBuffer)
	log.Printf("Original guard path goes over %v individual tiles", len(originalPathCoordinates))

	// "The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice."
	delete(originalPathCoordinates, coordinate{startX, startY})

	succesfullCoordinates := make([]coordinate, 0, len(originalPathCoordinates))
	nbSolutions := 0
	for coord := range originalPathCoordinates {
		cloneMatrix(originalMatrix, matrixBuffer)
		matrixBuffer[coord.y][coord.x] = '#'
		if _, hasLooped := play(matrixBuffer, startX, startY); hasLooped {
			succesfullCoordinates = append(succesfullCoordinates, coord)
			nbSolutions++
		}
	}

	log.Printf("Solutions: %v", nbSolutions)

	log.Print("Solution Map:")
	cloneMatrix(originalMatrix, matrixBuffer)
	for _, coord := range succesfullCoordinates {
		matrixBuffer[coord.y][coord.x] = '0'
	}
	for _, row := range matrixBuffer {
		fmt.Printf("%v\n", string(row))
	}

}

func play(matrix [][]rune, startX int, startY int) (map[coordinate]bool, bool) {
	path := make(map[coordinate]bool, len(matrix)*len(matrix[0]))

	x, y, stillWithinBoundary := startX, startY, true
	currentDirection := matrix[y][x]
	matrix[y][x] = '.'
	for ; stillWithinBoundary; x, y, stillWithinBoundary = nextTileCoordinates(matrix, currentDirection, x, y) {
		if matrix[y][x] == '#' {
			// Change direction
			currentDirection = ninetyDegreeClockwise(currentDirection)
			// Backtrack
			x, y, _ = nextTileCoordinates(matrix, ninetyDegreeClockwise(currentDirection), x, y)
			// Get to right side coordinate
			x, y, stillWithinBoundary = nextTileCoordinates(matrix, currentDirection, x, y)
			if stillWithinBoundary {
				matrix[y][x] = currentDirection
			}
		} else {
			if currentDirection == matrix[y][x] {
				return path, true
			}
			matrix[y][x] = currentDirection
		}
		path[coordinate{x, y}] = true
	}

	return path, false
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
		log.Fatalf("Unknown direction '%v'", string(from))
		return 'X'
	}
}

func nextTileCoordinates(matrix [][]rune, direction rune, x, y int) (nX int, nY int, respectsBoundary bool) {
	switch direction {
	case '^':
		if 0 <= y-1 {
			respectsBoundary = true
		}
		nX, nY = x, y-1
	case '>':
		if x < len(matrix[y])-1 {
			respectsBoundary = true
		}
		nX, nY = x+1, y
	case 'v':
		if y < len(matrix)-1 {
			respectsBoundary = true
		}
		nX, nY = x, y+1
	case '<':
		if 0 <= x-1 {
			respectsBoundary = true
		}
		nX, nY = x-1, y
	default:
		log.Fatalf("Unknown direction '%v'", string(direction))
		nX, nY = 0, 0
	}
	return nX, nY, respectsBoundary
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Printf("%v\n", string(row))
	}
}
