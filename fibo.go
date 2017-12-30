package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
 value0 := 0
 value1 := 1
 iter := 0
 return func() int {
 
 switch iter {
 case 0:
 	fib := 0
	iter +=1
 	return fib
case 1:
 	fib := 1
	iter+=1
 	return fib	

 default:
 fib := value0 +value1
 value0 = value1
 value1 = fib
 
  return fib
  }
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

