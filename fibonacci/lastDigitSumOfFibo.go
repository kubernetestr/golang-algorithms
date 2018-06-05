package main

import "fmt"

func Fibo(n int) int {
	sum := 0
	array := make([]int, n+1, n+2)
	if n < 2 {
		array = array[0:2]
	}
	array[0] = 0
	array[1] = 1
	for i := 2; i <= n; i++ {
		array[i] = array[i-1]%10 + array[i-2]%10

		sum = sum + array[i]
	}

	return (sum + 1) % 10
}

func main() {

	fmt.Println(Fibo(100))

}
