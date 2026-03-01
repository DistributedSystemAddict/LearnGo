package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main2() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	URL := os.Args[1]
	data, err := http.Get(URL) //dât: *http.Response
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Body.Close()
}
