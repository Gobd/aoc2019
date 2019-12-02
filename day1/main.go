package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var d1, d2 int
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		f := (i / 3) - 2
		d1 += f
		d2 += f

		for f = (f / 3) - 2; f > 0; f = (f / 3) - 2 {
			d2 += f
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Day 1: %d\n", d1) // 3337766
	fmt.Printf("Day 2: %d\n", d2) // 5003788
	fmt.Printf("Time elapsed: %s", time.Since(start))
}
