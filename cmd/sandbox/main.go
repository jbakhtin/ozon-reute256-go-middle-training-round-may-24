package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in := bufio.NewReader(inputFile)

	var buf bytes.Buffer
	count, err := buf.ReadFrom(in)

	//countBytes, err := in.Bytes()
	fmt.Println(count, buf.Bytes())
}
