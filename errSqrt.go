package main

import (
	"fmt"
	
)


type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string {
 	var str string
	str = fmt.Sprintf("Cannot Sqrt neg num : ")
	str += fmt.Sprint(float64(e))
	return str
}
func Sqrt(x float64) (float64, error) {

	if (x < 1) {
		return 0,ErrNegativeSqrt(x)
	}
        z := x/2
        for  i := 0 ; i < 10 ; i++{
                //fmt.Println(i, z)
                z -= (z*z -x)/(2*z)
        }
       // fmt.Print(" Sqrt with math ", math.Sqrt(x))
        return z,nil
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

