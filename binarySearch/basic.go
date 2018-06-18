package main
import "fmt"
func search (array []int, low int, high int, key int) int {



  if array[high]<array[low] {
     return low
   }
  mid := low + (high-low)/2

  if key == array[mid] {
     return mid
 } else if   key< array[mid] {
    return search(array, low, mid-1,key)
  } else {
    return search(array, mid+1 , high, key)
  }

}

func main() {

  dizi := []int {10,15,17,18,19,25,28,29,31,32,39}
  sonuc := search(dizi, 0, 10, 25)
  fmt.Println(sonuc)
}
