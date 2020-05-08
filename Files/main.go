package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file name will be index 1 of os.Args
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
