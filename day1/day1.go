package day1

import (
	"sort"
	"strconv"
	"os"
	"bufio"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// For Star1 just print max of all the sums after reading the ints from the file.  
func RunStar1(path string) int {
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	sum := 0
	maxSum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			maxSum = max(maxSum, sum)
			sum = 0
		}
		number, _ := strconv.Atoi(scanner.Text())
		sum += number
	}
	return maxSum
}

// for second star, read the file and print the sum of last three maximum sums. 
func RunStar2(path string) int {
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	numbers := make([]int, 0)

	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			numbers = append(numbers, sum)
			sum = 0
		}
		number, _ := strconv.Atoi(scanner.Text())
		sum += number
	}

	sort.Ints(numbers)
	length := len(numbers)
	return numbers[length-1] + numbers[length-2]+numbers[length-3]
}