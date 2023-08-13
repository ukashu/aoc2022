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

	// [shape my opponent chose][shape points i would get for loss, draw, win]
	twoD := [3][3]int{{3, 1, 2}, {1, 2, 3}, {2, 3, 1}}
	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		moves := scanner.Text()
		myMove := moves[2] - 88
		oppMove := moves[0] - 65
		total += twoD[oppMove][myMove]
		total += int(myMove) * 3 // points for loss/draw/win
	}

	log.Print(total)
}