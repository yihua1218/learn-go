package main

import (
	"fmt"

	"github.com/yihua1218/dynamo-go"
)

func main() {
	db := dynamo.New("us-east-2")
	fmt.Println(db.Region)

}
