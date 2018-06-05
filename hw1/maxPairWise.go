package main

import "fmt"

func main() {
  var n, biggest, biggest2, c int
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,105,
  }

  for _,v:=range x {
    if v>n {
      n = v
      biggest = n

    }
  }

  fmt.Println("The biggest number is ", biggest)
  for _,v:=range x {
    if v>c && v!=biggest {
      c = v
      biggest2 = c

    }
  }

  fmt.Println("The biggest number is ", biggest2)

  fmt.Println("max pairwise= ", biggest*biggest2)
}
