package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func ( red MyReader) Read(out []byte) (int,error){
	for i := range out {
		out[i] = 65
	}
	return len(out), nil
}

func main() {
	reader.Validate(MyReader{})
}
