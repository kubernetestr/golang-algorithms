package main
import "fmt"

func leastCM(x int, y int)int{
  i:=0
  if x<y{
    i=y
  } else {
    i=x
  }
  for i<=x*y{
    if i%x==0  && i%y==0{
      return i
    }
  i=i+1
  }

return x*y

}

func main (){
  a:= leastCM(15,30)
  fmt.Println(a)

}
