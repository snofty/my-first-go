package main

import "fmt"

func main() {
	var x = [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	ln := len(x)
	var total float64 = 0
	/*for i:=0;i<ln;i++{
	  total +=x[i]
	}*/

	//A single _ (underscore) is used to tell the compiler that we don't need this.
	//(In this case we don't need the iterator variable represents the current position i)
	// value = x[i]
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(ln))

	fmt.Println("Write a program that finds the smallest number in this list")
	xa := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
  var smallNum int = xa[0]
  for _,value := range xa{
    if(value<smallNum){
      smallNum=value;
    }
  }
  fmt.Printf("%s%d","Smaller is ",smallNum)
}
