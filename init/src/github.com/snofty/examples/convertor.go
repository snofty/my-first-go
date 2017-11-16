package main

import "fmt"

func main(){
  var ft float64
  const ftInMtr = 0.3048
  fmt.Println("Enter feets to convert to meters")
  fmt.Scanf("%f",&ft)
  mtr := ft*ftInMtr
  fmt.Printf("%f%s%f%s",ft," feet =",mtr,"m")
}
