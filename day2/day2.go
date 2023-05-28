package day2

import (
	"os"
	"bufio"
	"strings"
)

func RunStar1(path string) int {
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	score := 0
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		if vals[0] == "A" {
			if vals[1] == "X" {
				score += 4
			} else if vals[1] == "Y" {
				score += 8
			} else if vals[1] == "Z" {
				score += 3
			}
		} else if vals[0] == "B" {
			if vals[1] == "X" {
				score += 1
			} else if vals[1] == "Y" {
				score += 5
			} else if vals[1] == "Z" {
				score += 9
			}
		} else if vals[0] == "C" {
			if vals[1] == "X" {
				score += 7
			} else if vals[1] == "Y" {
				score += 2
			} else if vals[1] == "Z" {
				score += 6
			}
		}
	}
	return score
}

func RunStar2(path string) int {
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	score := 0
	// X == loss , Y == Draw, Z == Win
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		if vals[0] == "A" { // Stone
			if vals[1] == "X" { // Scissors
				score += 3
			} else if vals[1] == "Y" { // Stone
				score += 4
			} else if vals[1] == "Z" { // Paper
				score += 8
			}
		} else if vals[0] == "B" { // Paper
			if vals[1] == "X" { // Stone
				score += 1
			} else if vals[1] == "Y" { // Paper
				score += 5
			} else if vals[1] == "Z" { // Scissors
				score += 9
			}
		} else if vals[0] == "C" {//Scissor
			if vals[1] == "X" { // Paper
				score += 2
			} else if vals[1] == "Y" { // Scissors
				score += 6
			} else if vals[1] == "Z" { // Stone
				score += 7
			}
		}
	}
	return score
}