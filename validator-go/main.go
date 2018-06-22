package main

import (
	"fmt"
	"os"

	"github.com/yihua1218/validator-go"
)

func main() {
	fmt.Printf("len(os.Args: %v\n", len(os.Args))
	fmt.Printf("os.Args: %v\n", os.Args)

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	v, e := validator.New(os.Args[1])

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(2)
	}

	v.Test()
}
