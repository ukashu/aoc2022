package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("./input.txt") // For read access.
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i int
	currentElf := 0
	third := 0
	second := 0
	first := 0
	total := 0

	for i = 0; scanner.Scan(); i++ {
		current := scanner.Text()
		if len(current) == 0 {
			if currentElf > first {
				third = second
				second = first
				first = currentElf
			} else if currentElf > second {
				third = second
				second = currentElf
			} else if currentElf > third {
				third = currentElf
			}
			currentElf = 0
		} else {
			int, err := strconv.Atoi(current)
			check(err)
			currentElf += int
		}
	}
	total = first + second + third
	fmt.Printf("First elf: %v, second elf: %v, third elf: %v\nTotal: %v\n", first, second, third, total)
}