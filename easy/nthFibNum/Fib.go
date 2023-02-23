package nthFibNum

var fibNums [10000]int

func Fib(n int) int {
	if n == 1 || n == 2 {
		return fibNums[n-1]
	}
	if fibNums[n-1] == 0 {
		fibNums[n-1] = Fib(n-1) + Fib(n-2)
	}
	return fibNums[n-1]
}
