package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type relationship struct {
	precededBy map[int]bool
	followedBy map[int]bool
}

var relationshipParser = regexp.MustCompile(`(?P<before>[0-9]*)\|(?P<after>[0-9]*)`)
var beforeSEI, afterSEI = relationshipParser.SubexpIndex("before"), relationshipParser.SubexpIndex("after")

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	relationshipsMap := map[int]relationship{}
	var validUpdatesTotal int64 = 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		} else if subMatches := relationshipParser.FindStringSubmatch(text); len(subMatches) > 0 {
			before, errL := strconv.Atoi(subMatches[beforeSEI])
			if errL != nil {
				log.Fatalf("Error while trying to parse left side of the rule %v\nerr: %v", subMatches[beforeSEI], errL)
			}
			after, errR := strconv.Atoi(subMatches[afterSEI])
			if errR != nil {
				log.Fatalf("Error while trying to parse right side of the rule %v\nerr: %v", subMatches[afterSEI], errR)
			}

			_, exists := relationshipsMap[before]
			if !exists {
				relationshipsMap[before] = relationship{precededBy: map[int]bool{}, followedBy: map[int]bool{}}
			}
			relationshipsMap[before].followedBy[after] = true

			_, exists = relationshipsMap[after]
			if !exists {
				relationshipsMap[after] = relationship{precededBy: map[int]bool{}, followedBy: map[int]bool{}}
			}
			relationshipsMap[after].precededBy[before] = true
		} else if updateStrs := strings.Split(text, ","); len(updateStrs) > 0 {
			updateLine := make([]int, len(updateStrs))
			for i, updateStr := range updateStrs {
				number, err := strconv.Atoi(updateStr)
				if err != nil {
					log.Fatalf("Error while trying to parse update row value %v, text='%v', row='%v' (asArr=%v)", updateStr, text, updateStrs, len(updateStrs))
				}
				updateLine[i] = number
			}
			if !isUpdateValid(updateLine, relationshipsMap) {
				quickSort(updateLine, relationshipsMap, 0, len(updateLine)-1)
				validUpdatesTotal += int64(updateLine[len(updateLine)/2])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %q", err)
	}

	log.Printf("total = %v\n", validUpdatesTotal)
}

func isUpdateValid(updateLine []int, rulesMap map[int]relationship) bool {
	for i, number := range updateLine {
		for j, comparedTo := range updateLine {
			if j == i {
				continue
			}
			if j < i && rulesMap[number].followedBy[comparedTo] {
				return false
			}
			if i < j && rulesMap[number].precededBy[comparedTo] {
				return false
			}
		}
	}
	return true
}

func quickSort(list []int, rulesMap map[int]relationship, lowerBound int, upperBound int) {
	if lowerBound < upperBound {
		pivot := list[upperBound]

		i := lowerBound - 1

		for j := lowerBound; j < upperBound; j++ {
			if rulesMap[list[j]].followedBy[pivot] {
				i++
				list[i], list[j] = list[j], list[i]
			}
		}

		list[i+1], list[upperBound] = list[upperBound], list[i+1]

		pivotIndex := i + 1
		quickSort(list, rulesMap, lowerBound, pivotIndex-1)
		quickSort(list, rulesMap, pivotIndex+1, upperBound)
	}
}
