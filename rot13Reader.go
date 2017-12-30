package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rdr rot13Reader) Read(bytes []byte) (int, error) {
	 n,err := rdr.r.Read(bytes)
	 if( err != nil){
			fmt.Sprint("Error: %v",err)
			}
	 fmt.Sprint("Bytes read : %v",n)
	for i := range bytes {
		if((bytes[i] -65) <13){
		bytes[i] = bytes[i]+13
		}else{
		bytes[i] = bytes[i] - 13
		}
	}
	return len(bytes), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
