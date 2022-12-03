package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

var (
	waysOfAttack = map[string]map[string]int{
		"win": {
			"A": 2,
			"B": 3,
			"C": 1,
		},
		"lose": {
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"draw": {
			"A": 1,
			"B": 2,
			"C": 3,
		},
	}	

	count int
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        attacks := strings.Split(scanner.Text(), " ")

		/* draw */
		if attacks[1] == "Y" {
			count += (waysOfAttack["draw"][attacks[0]] + 3)
			continue;
		}

		/* wins */
		if attacks[1] == "Z" {
			count += (waysOfAttack["win"][attacks[0]] + 6)
			continue;
		}

		/* Lose */
		if attacks[1] == "X" {
			count += waysOfAttack["lose"][attacks[0]]
			continue;
		}
    }

	fmt.Println(count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}