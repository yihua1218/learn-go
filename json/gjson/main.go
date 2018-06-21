package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
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

	fmt.Printf("file: %s\n", file)

	json := string(file)

	root := gjson.Parse(json)
	headers := gjson.Get(json, "headers")

	fmt.Printf("headers: %v\n", headers)

	requestContext := gjson.Get(string(file), "requestContext")

	requestContext.ForEach(func(key, value gjson.Result) bool {
		fmt.Printf("key: %s\n", key)
		return true
	})

	root.ForEach(func(key, value gjson.Result) bool {
		fmt.Printf("key: %s\n", key)
		return true
	})
}
