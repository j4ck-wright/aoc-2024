package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumOfSlice(s []int) int {
	n := 0

	for i := 0; i < len(s); i++ {
		n = n + s[i]
	}

	return n
}

func frequencyMap(s []int) map[int]int {
	fm := make(map[int]int)
	for i := 0; i < len(s); i++ {
		fm[s[i]]++
	}

	return fm
}

func main() {

	input, err := os.Open("day-1/input.txt")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	scanner := bufio.NewScanner(input)

	listLeft := make([]int, 0)
	listRight := make([]int, 0)
	distancesApart := make([]int, 0)
	similarityScore := 0

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")

		leftVal, _ := strconv.Atoi(values[0])
		rightVal, _ := strconv.Atoi(values[1])

		listLeft = append(listLeft, leftVal)
		listRight = append(listRight, rightVal)
	}

	sort.Ints(listLeft)
	sort.Ints(listRight)

	rightFreqMap := frequencyMap(listRight)

	for i := 0; i < len(listLeft); i++ {
		num := listLeft[i]
		distance := num - listRight[i]

		// +ve distance
		if distance < 0 {
			distance = distance * -1
		}

		distancesApart = append(distancesApart, distance)
		occurances := rightFreqMap[num]
		similarityScore += (num * occurances)

	}

	sumDistances := sumOfSlice(distancesApart)
	fmt.Printf("The sum of distances is: %d\n", sumDistances)
	fmt.Printf("The similarity score is: %d\n", similarityScore)
}
