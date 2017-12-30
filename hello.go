package main

import "fmt"
import "math"

func Sqrt(x float64) float64{

	z := x/2
	for  i := 0 ; i < 10 ; i++{
	 	fmt.Println(i, z)
		z -= (z*z -x)/(2*z)
	}
	fmt.Print(" Sqrt with math ", math.Sqrt(x))
	return z
}

func main() {
    fmt.Printf("hello, world\n")
	fmt.Print("Value of sqrt is 2 ", Sqrt(2))
	fmt.Print("Value of sqrt 3 is " , Sqrt(3))
	fmt.Print("Value of sqrt 4 is ", Sqrt(4))
}
