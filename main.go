package main

import (
	"fmt"

	"github.com/qascade/AoC22/day1"
	"github.com/qascade/AoC22/day2"
)

func main() {
 	fmt.Println("Answer for Star1, day1: ", day1.RunStar1("./day1/input.txt"))
	fmt.Println("Answer for Star2, day1: ", day1.RunStar2("./day1/input.txt"))
	fmt.Println("Answer for Star1, day2: ", day2.RunStar1("./day2/input.txt"))
	fmt.Println("Answer for Star2, day2: ", day2.RunStar2("./day2/input.txt"))

}