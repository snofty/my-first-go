package main

import "fmt"

func main(){
  fmt.Println(len("Hello world"))
  fmt.Println("Hello world"[1])
  fmt.Println("Hello "+"world")
  fmt.Println(32132 * 42452)
  fmt.Println((true && false) || (false && true) || !(false && false))
}
