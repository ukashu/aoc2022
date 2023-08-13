package main

import (
	//"fmt"
	"log"
	"os"
	"bufio"
	//"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	log.SetPrefix("day2: ")
	log.SetFlags(0)

	file, err := os.Open("./input.txt") // For read access.
	check(err)
	defer file.Close()

	twoD := [3][3]int{{3, 6, 0}, {0, 3, 6}, {6, 0, 3}}
	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		moves := scanner.Text()
		myMove := moves[2] - 88
		total += twoD[1][myMove]
		total += int(myMove) + 1
	}

	log.Print(total)
}