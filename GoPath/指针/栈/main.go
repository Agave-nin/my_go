package main

import "fmt"

func main() {
	//init
	stack := []int{}
	ele1 := 42
	stack = append([]int{ele1}, stack...)//创建一个新队列，和stack拼接
	ele2 := 420
	stack = append([]int{ele2}, stack...)
	fmt.Println(stack)

	pop := stack[0]
	fmt.Println(pop)
	stack = stack[1:]
	fmt.Println(stack)
}
