package fib

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	res := Fib(n-1) + Fib(n-2)
	return res
}
