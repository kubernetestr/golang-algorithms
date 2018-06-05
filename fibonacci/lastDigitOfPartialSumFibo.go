package main

import "fmt"

func Fibo(n int, m int) int {
	sum := 0
	array := make([]int, m+1, m+2)
	if n < 2 {
		array = array[0:2]
	}
	array[0] = 0
	array[1] = 1
	for i := 2; i <= m; i++ {
		array[i] = array[i-1]%10 + array[i-2]%10
	}
	for d := n; d <= m; d++ {
		sum = sum + array[d]

	}

	return (sum) % 10
}

func main() {

	fmt.Println(Fibo(10,200))

}
