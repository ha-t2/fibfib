package main

import (
	"flag"
	"fmt"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	flag.Parse()
	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(fib(n))
}
