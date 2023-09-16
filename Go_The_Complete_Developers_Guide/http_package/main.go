package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)

	}

	// low level printing of the body data
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//this produces the same result as the code above, but uses a function thar takes a function that implements the reader and another that implments the writer interface to acess an output the data
	//io.Copy(os.Stdout, resp.Body)

	//now we write out custom function that implement a writer to use with io.Copy
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrotte this many bytes: ", len(bs))
	return len(bs), nil
}
