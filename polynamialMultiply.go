package main
import "fmt"

func multiply (a []int, b []int , n int ) []int {

  product := []int{0,0,0,0,0}
  for i:=0 ; i< 3 ; i++ {
    for j:=0 ; j< 3 ; j++ {
      product [i+j] = product [i+j] + a[i]*b[j]
      
     }
   }

   return product
 }

func main(){
  a := []int{3, 2, 5}
  b := []int{5, 1, 2}

  product := (multiply (a, b, 3))

  fmt.Println(product)


}
