package main

import (
	"bufio"
	"fmt"
	"os"
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

	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Day 1: %d\n", d1)
	fmt.Printf("Day 2: %d\n", d2)
	fmt.Printf("Time elapsed: %s", time.Since(start))
}
