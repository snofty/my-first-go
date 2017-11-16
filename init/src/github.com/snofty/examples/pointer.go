package main

import "fmt"

func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)

	var x1 int = 12
	var x2 int = 45
	fmt.Println(x1, x2)
	fmt.Println("Swaping")
	swap(&x1, &x2)
	fmt.Println(x1, x2)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x1 *int, x2 *int) {
  temp := new(int)
	*temp = *x1
	*x1 = *x2
	*x2 = *temp
}
