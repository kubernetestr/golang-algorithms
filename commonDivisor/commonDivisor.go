package main
import "fmt"

func gcd(x int, y int)int {
  var current_gcd int
  current_gcd=1
for i := 2; i <= x || i <= y; i++ {

if x % i == 0 && y % i == 0 {
    if i > current_gcd {
      current_gcd = i;
    }

}

}
return current_gcd
}

func main(){
fmt.Println(gcd(28851538,1183019))

}
