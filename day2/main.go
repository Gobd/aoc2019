package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inp, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	opcodesStr := strings.Split(string(inp), ",")
	opcodes := make([]int, len(opcodesStr))
	for i := range opcodesStr {
		opcodeInt, err := strconv.Atoi(opcodesStr[i])
		if err != nil {
			panic(err)
		}
		opcodes[i] = opcodeInt
	}

	target := make([]int, len(opcodes))
	d1 := add(get(opcodes, target, 12, 2))
	var d2 int

	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			ans := add(get(opcodes, target, n, v))
			if ans == 19690720 {
				d2 = (100 * n) + v
			}
		}
	}

	fmt.Printf("Day 1: %d\n", d1) // 3267740
	fmt.Printf("Day 2: %d\n", d2) // 7870
	fmt.Printf("Time elapsed: %s", time.Since(start))
}

func get(opcodes, target []int, n, v int) []int {
	copy(target, opcodes)
	target[1] = n
	target[2] = v
	return target
}

func add(opcodes []int) int {
	for i := 0; i < len(opcodes)-1; i += 4 {
		a := opcodes[opcodes[i+1]]
		b := opcodes[opcodes[i+2]]
		dest := opcodes[i+3]

		switch opcode := opcodes[i]; opcode {
		case 1:
			opcodes[dest] = a + b
		case 2:
			opcodes[dest] = a * b
		case 99:
			return opcodes[0]
		default:
			panic("Bad opcode")
		}
	}
	return opcodes[0]
}
