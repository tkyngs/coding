package main

import (
  "fmt"
)

func main() {
  a,b:=ni(),ni()
  if a>0 && b>0{
    fmt.Println("Alloy")
  } else if a>0 && b==0 {
    fmt.Println("Gold")
  } else {
    fmt.Println("Silver")
  }
}

func ni() int {
  var n int
  fmt.Scan(&n)
  return n
}

func nis(n int) []int {
  a:=make([]int, n)
  for i:=0;i<n;i++ {
    fmt.Scan(&a[i])
  }
  return a
}
