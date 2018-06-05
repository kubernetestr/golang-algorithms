package main
import "fmt"


   func Fibonacci(n int) int {
      if n<=1 {
        return n
      }
      return Fibonacci(n-2)+ Fibonacci(n-1)

   }

   func main(){

     sum:= Fibonacci(8)
     fmt.Println(sum)


   }
