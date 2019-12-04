package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y float64
}

func main() {
	start := time.Now()
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	w1 := createWire(s.Text())

	s.Scan()
	w2 := createWire(s.Text())

	if err := s.Err(); err != nil {
		panic(err)
	}

	var intersections []float64
	var steps []int
	for pt, w1steps := range w1 {
		if w2steps, ok := w2[pt]; ok {
			intersections = append(intersections, math.Abs(pt.x)+math.Abs(pt.y))
			steps = append(steps, w1steps+w2steps)
		}
	}
	sort.Float64s(intersections)
	sort.Ints(steps)

	fmt.Printf("Day 1: %d\n", int(intersections[0])) // 258
	fmt.Printf("Day 2: %d\n", steps[0])              // 12304
	fmt.Printf("Time elapsed: %s", time.Since(start))
}

func createWire(instrs string) map[point]int {
	m := make(map[point]int)
	cur := point{}
	var steps int

	for _, instr := range strings.Split(instrs, ",") {
		dir := string(instr[0])
		dist, err := strconv.Atoi(instr[1:])
		if err != nil {
			panic(err)
		}

		for i := 1; i <= dist; i++ {
			steps++
			switch dir {
			case "R":
				cur.x++
			case "L":
				cur.x--
			case "U":
				cur.y++
			case "D":
				cur.y--
			default:
				panic(instr)
			}
			if previousSteps, ok := m[cur]; !ok || steps < previousSteps {
				m[cur] = steps
			}
		}
	}

	return m
}
