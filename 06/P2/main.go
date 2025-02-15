package main

// Identified errors:
// - If a path is rejoined, it might be a path posterior to where the obstacle is put
//	Therefore, assuming that coming to the tile {x,y}, that has the same value as our-
//	current direction doesn't mean we are in a loop. We can be on the way out !
// - Every obstacle that's missing in this implementation involves a "straight 180°" pattern
//	 ie. when two existing obstacles are in diagonal, and the guard arrives between them
//			  ^ this causes a double 90 degree turn, hence the 180°

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

type X struct {
	x         int
	y         int
	direction rune
}

func main() {
	originalMatrix := getMatrix("../input")
	var startX = 0
	var startY = 0

	for y, row := range originalMatrix {
		for x, tile := range row {
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

// Mon x est une structure donnant ma position et ma direction actuelle
// Le f donne la structure suivante
func brent(f func(X) X, x0 X) (lambda int, mu int) {
	tortoise := x0
	hare := f(x0)
	i, lambda := 1, 1

	for ; tortoise != hare; hare, lambda = f(hare), lambda+1 {
		if i == lambda {
			tortoise = hare
			i *= 2
			lambda = 0
		}
	}

	tortoise, hare = x0, x0
	for range lambda {
		hare = f(hare)
	}

	for ; tortoise != hare; mu++ {
		tortoise = f(tortoise)
		hare = f(hare)
	}

	return lambda, mu
}

func getMatrix(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matrix := make([][]rune, 0, 131)

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		text := scanner.Text()
		matrix = append(matrix, []rune(text))
	}

	return matrix
}

func getGoodBadComparison(base [][]rune, good [][]rune, bad [][]rune) [][]rune {
	comparison := make([][]rune, 130)
	for y := range comparison {
		comparison[y] = make([]rune, 130)
		for x := range comparison[y] {
			if val := good[y][x]; val == '0' && bad[y][x] == '.' {
				comparison[y][x] = 'X'
			} else if val == '.' && bad[y][x] == '0' {
				comparison[y][x] = '+'
			} else if val == '0' {
				comparison[y][x] = '0'
			} else {
				comparison[y][x] = base[y][x]
			}
		}
	}
	return comparison
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
