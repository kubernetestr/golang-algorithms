package main

import  "fmt"


func Fibo(n int, m int) int {
    array := make([]int, n+1, n+2)
    if n < 2 {
        array = array[0:2]
    }
    array[0] = 0
    array[1] = 1
    for i := 2; i <= n; i++ {
        array[i] = array[i-1]%m + array[i-2]%m
    }
    return array[n]%m
}


func main() {





    fmt.Println(Fibo(281621, 7))



}
