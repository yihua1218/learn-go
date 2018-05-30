package main

import (
	"fmt"

	dynamo "github.com/yihua1218/dynamo-go"
)

func main() {
	db := dynamo.New("us-west-2")

	fmt.Println("main")
}
