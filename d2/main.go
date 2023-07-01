package d2

import (
	"bufio"
	"os"
	"strings"
)

func Score(eMove string, myMove string) int {
	score := 0
	switch eMove {
	case "A": // rock
		if myMove == "X" {
			score = 3
		}
		if myMove == "Y" {
			score = 3 + 1
		}
		if myMove == "Z" {
			score = 6 + 2
		}
	case "B": // paper
		if myMove == "X" {
			score = 1
		}
		if myMove == "Y" {
			score = 2 + 3
		}
		if myMove == "Z" {
			score = 6 + 3
		}
	case "C": //sciccors
		if myMove == "X" {
			score = 2
		}
		if myMove == "Y" {
			score = 3 + 3
		}
		if myMove == "Z" {
			score = 6 + 1
		}
	}

	return score
}

func Day2() int {
	reader := bufio.NewScanner(os.Stdin)

	score := 0

	for reader.Scan() {
		line := reader.Text()
		moves := strings.Split(line, " ")

		eMove := moves[0]
		myMove := moves[1]

		score += Score(eMove, myMove)

	}

	return score

}
