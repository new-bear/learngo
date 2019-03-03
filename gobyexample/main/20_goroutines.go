package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("go routine")

	// 匿名函数协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")
}
