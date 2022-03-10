package main

import "fmt"

func main() {
	queue := []int{}
	queue = append(queue, 42)
	queue = append(queue, 420)
	fmt.Println(queue)
	first := queue[0]
	queue = queue[1:]
	fmt.Println(first, queue)
}
