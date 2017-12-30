package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
m := make(map[string]int)
var strs []string = strings.Fields(s)

for _,value := range strs{

var elem = m[value];
elem += 1
m[value] = elem

}
	return m
}

func main() {
	wc.Test(WordCount)
}

