package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type relationship struct {
	before map[int]bool
	after map[int]bool
}

var relationshipParser = regexp.MustCompile(`^(?P<before>[0-9]*)|(?P<after>[0-9]*)$`)
var beforeSEI, afterSEI = relationshipParser.SubexpIndex("before"), relationshipParser.SubexpIndex("right")

var updateParser = regexp.MustCompile(`(?P<number>[0-9]*),?`)
var numberSEI = updateParser.SubexpIndex("number")

func main() {
	// Instanciate the matrix
	file, err := os.Open("../easyinput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	relationshipsMap := map[int]relationship{}
	for scanner.Scan() {
		text := scanner.Text()
		
		if subMatches := relationshipParser.FindStringSubmatch(text); len(subMatches) > 0 {
			before, err := strconv.Atoi(subMatches[beforeSEI])
			if err != nil {
				log.Fatalf("Error while trying to parse left side of the rule %v", subMatches[beforeSEI])
			}
			after, err := strconv.Atoi(subMatches[afterSEI])
			if err != nil {
				log.Fatalf("Error while trying to parse left side of the rule %v", subMatches[beforeSEI])
			}

			_, exists := relationshipsMap[before]
			if !exists {
				relationshipsMap[before] = relationship{ before: map[int]bool{}, after: map[int]bool{}}
			}
			relationshipsMap[before].after[after] = true

			_, exists = relationshipsMap[after]
			if !exists {
				relationshipsMap[after] = relationship{ before: map[int]bool{}, after: map[int]bool{}}
			}
			relationshipsMap[after].before[before] = true
			} else if subMatches := updateParser.FindAllStringSubmatch(text, -1); len(subMatches) > 0 {
				updateLine := make([]int, 0, len(subMatches))
				for _, match := range(subMatches) {
					number, err := strconv.Atoi(match[numberSEI])
					if err != nil {
						log.Fatalf("Error while trying to parse update row value %v", match[numberSEI])
					}
					updateLine = append(updateLine, number)
				}
				// TODO: do stuff with that data
			} else {
				log.Println("emptyline")
			}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}
}