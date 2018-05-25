package main

import (
	"fmt"
	"go/build"
	"os"

	dynamo "github.com/yihua1218/dynamo-go"
)

func main() {
	gopath := os.Getenv("GOPATH")

	if gopath == "" {

		gopath = build.Default.GOPATH
	}
	fmt.Println(gopath)

	db := dynamo.New("us-east-2")
	db.ListTables()
	db.Query("g3_devices", "03AFB154")
	db.DescribeTable("g3_devices")
	fmt.Println(db.Region)
}
