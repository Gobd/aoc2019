package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	var d1, d2 int
	for i := 387638; i <= 919123; i++ {
		d1g, d2g := testPassword(i)
		if d1g {
			d1++
		}
		if d2g {
			d2++
		}
	}

	fmt.Printf("Day 1: %d\n", d1) // 466
	fmt.Printf("Day 2: %d\n", d2) // 292
	fmt.Printf("Time elapsed: %s", time.Since(start))
}

func testPassword(password int) (bool, bool) {
	s := strconv.Itoa(password)
	if len(s) != 6 {
		return false, false
	}

	var prev int32 = -1
	var d1, d2 bool
	dm := make(map[int32]int)
	for _, cur := range s {
		if cur < prev {
			return false, false
		} else if cur == prev {
			dm[cur]++
			d1 = true
		}
		prev = cur
	}

	for _, d := range dm {
		if d == 1 {
			d2 = true
		}
		break
	}

	return d1, d2
}
