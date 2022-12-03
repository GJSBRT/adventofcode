package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

var (
	attackPoints = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	waysOfAttack = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
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
		if attackPoints[attacks[0]] == attackPoints[attacks[1]] {
			count += (attackPoints[attacks[1]] + 3)
			continue;
		}

		/* wins */
		if waysOfAttack[attacks[0]] != attacks[1] {
			count += (attackPoints[attacks[1]] + 6)
			continue;
		}

		/* Lose */
		if waysOfAttack[attacks[0]] == attacks[1] {
			count += attackPoints[attacks[1]]
			continue;
		}
    }

	fmt.Println(count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}