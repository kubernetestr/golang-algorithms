package main

import  "fmt"


func Fibo(n int) int {
    array := make([]int, n+1, n+2)
    if n < 2 {
        array = array[0:2]
    }
    array[0] = 0
    array[1] = 1
    for i := 2; i <= n; i++ {
        array[i] = array[i-1] + array[i-2]
    }
    return array[n]
}


func main() {

  var x int

    x = Fibo(80)%10

    fmt.Println(Fibo(80))


    fmt.Println(x)
}
