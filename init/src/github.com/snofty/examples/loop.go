package main

import "fmt"

func main(){
  i :=1
  for i<=10{
    fmt.Println(i)
    i=i+1
  }
  for j:=1; j<=10;j++{
    if(j%2==0){
      fmt.Printf("%d%s",j,"-Even")
    }else{
      fmt.Printf("%d%s",j,"-Odd")
    }
    fmt.Println()
  }

  fmt.Println("Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100")
  for k:=1; k<=100;k++{
    if k%3==0{
      fmt.Println(k)
    }
  }

  fmt.Println("Write a program that prints the numbers from 1 to 100. But for multiples of three print \"Fizz\" instead of the number and for the multiples of five print \"Buzz\". For numbers which are multiples of both three and five print \"FizzBuzz\".")
  var ty string =""
  for k:=1; k<=100;k++{
    if k%3==0{
      ty ="Fizz"
    }else if k%5==0{
      ty = ty+"Buzz"
    }
    if(ty!=""){
      fmt.Println(ty)
    }else{
      fmt.Println(k)
    }
    ty=""
  }
}
