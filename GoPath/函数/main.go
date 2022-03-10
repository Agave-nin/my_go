package main

import (
	"fmt"
	"time"
)

func swap(x, y int) (int, int) {
	x, y = y, x
	return x, y
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextnumber := getSequence()

	t := time.Now()

	for i := 0; i < 2000000; i++ {
		fmt.Println(nextnumber())
	}

	time_all := time.Since(t)
	fmt.Println(time_all)

}
