package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
var s  [][]uint8;
for i := 0; i < dy ; i++{
var temp  []uint8;
 for j  := 0 ; j < dx ; j++{
 	
 	temp = append(temp,uint8(i)+uint8(j))
 }
 s = append(s,temp)
}
return s;
}

func main() {
	pic.Show(Pic)
}

