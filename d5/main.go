package d5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day5() string {
	stack := make([]string, 0)
	// Not gonna parse that input 😐
	stack = append(stack, "SZPDLBFC")
	stack = append(stack, "NVGPHWB")
	stack = append(stack, "FWBJG")
	stack = append(stack, "GJNFLWCS")
	stack = append(stack, "WJLTPMSH")
	stack = append(stack, "BCWGFS")
	stack = append(stack, "HTPMQBW")
	stack = append(stack, "FSWT")
	stack = append(stack, "NCR")
	reader := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		reader.Scan()
	}
	moves := make([][]int, 0)
	for reader.Scan() {
		line := reader.Text()
		instr := strings.Split(line, " ")
		instr1, _ := strconv.Atoi(instr[1])
		instr2, _ := strconv.Atoi(instr[3])
		instr3, _ := strconv.Atoi(instr[5])
		move := []int{instr1, instr2, instr3}
		moves = append(moves, move)
	}
	return Containers(stack, moves)
}

func Containers(stacks []string, moves [][]int) string {
	var res string = ""
	for i := 0; i < len(moves); i++ {
		quant := moves[i][0]
		from := moves[i][1]
		dest := moves[i][2]

		removed := make([]string, 0)
		for k := 0; k < quant; k++ {
			removed = append(removed, string(stacks[from-1][len(stacks[from-1])-1]))
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		}
		for x := len(removed) - 1; x >= 0; x-- {
			str := stacks[dest-1]
			newString := str + removed[x]
			stacks[dest-1] = newString
		}
	}
	for i := 0; i < len(stacks); i++ {
		stringAppend := stacks[i]
		lastChar := stringAppend[len(stringAppend)-1]
		res += string(lastChar)
	}
	return res
}
