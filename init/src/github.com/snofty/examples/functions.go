package main

import "fmt"

func main(){
  x := [10]int{
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
  }
  fmt.Println(add(x[:5],x[5:]))
  fmt.Println("Enter a integer value")
  var input int
  fmt.Scanf("%d",&input)
  _, isEven := half(input)
  if isEven {
    fmt.Println("Even")
  }else{
    fmt.Println("Odd")
  }

  fmt.Println(findGreater(34,232,14,562,452))

  nextEvent := makeOddGenerator()
  fmt.Println(nextEvent())
  fmt.Println(nextEvent())
  fmt.Println(nextEvent())
  
  fmt.Println("Enter a integer value to find fib")
  fmt.Scanf("%d",&input)
  fmt.Printf("%s%d","Fib is ",fib(input))
}

//Write a function which takes an integer and halves it and returns true if it was even or false if it was odd
func half(in int)(int, bool){
  hal := in/2
  return hal,hal%2==0
}

//sum is a function which takes a slice of numbers and adds them together.
func add(slice1 []int, slice2 []int) int {
  var total int = 0
  for _,value := range slice1{
    total += value
  }
  for _,value := range slice2{
    total += value
  }
  return total
}

//Write a function with one variadic parameter that finds the greatest number in a list of numbers.
func findGreater(list ...int) int{
  great := 0
  for _,value := range list{
    if(value>great){
      great = value
    }
  }
  return great
}

//write a makeOddGenerator function that generates odd numbers
func makeOddGenerator() func () uint{
  i := uint(1)
  return func()(ret uint){
    ret = i
    i += 2
    return
  }
}

//The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n)
func fib(n int) int {
  if n==1{
    return 1
  }
  if n==0{
    return 0
  }
  return fib(n-1)+fib(n-2)
}
