package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

/*
*************************************************** NOTES ****************************************************
	In go, interfaces can be nested inside another interface.
		Using the Code: type ReadCloser interface { Reader Closer }
						type Reader interface { Read([]byte) (int error) }
						type Closer interface { Close() (error) }

		ReadCloser -> top level interface
		Reader     -> second level interface
		Closer     -> second level interface
		Read()     -> function that forms the Reader interface using normal interface syntax
		Close()    -> function that forms the Closer interface using normal interface sytnax

	interfaces solve the issue of many inputs many outputs. It simplifies to many inputs one common output

	the make function allows us to create a dynamic slice with a starting number of elements

*/
