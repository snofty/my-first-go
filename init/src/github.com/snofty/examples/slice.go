package main

import "fmt"

func main(){
  var x []float64
  x = make([]float64,4,10)
  fmt.Println(len(x))

  slice1 := []int{1,2,3}
  slice2 := append(slice1,4,5)
  fmt.Println(slice1,slice2)
  slice3 := make([]int,2)
  copy(slice3,slice2)
  fmt.Println(slice3)
  fmt.Println(slice2[2:])
  fmt.Println(slice1[:2])

  xe := [6]string{"a","b","c","d","e","f"}
  //execluded the index elements.
  fmt.Println(xe[2:5])

  fmt.Println(len(make([]int, 3, 9)))
}
