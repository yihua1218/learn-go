package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("len(os.Args: %v\n", len(os.Args))
	fmt.Printf("os.Args: %v\n", os.Args)

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	filename := os.Args[1]
	file, e := ioutil.ReadFile(filename)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(3)
	}

	fmt.Printf("file: %v\n", file)
}
