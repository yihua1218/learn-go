package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	v "github.com/gima/govalid/v1"
)

func main() {
	filename := os.Args[1]
	var JSON map[string]interface{}
	file, e := ioutil.ReadFile(filename)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}
	json.Unmarshal(file, JSON)

	schema := v.Object(
		v.ObjKV("status", v.Boolean()),
		v.ObjKV("data", v.Object(
			v.ObjKV("debug", v.Number(v.NumMin(1), v.NumMax(99999))),
			v.ObjKV("items", v.Array(v.ArrEach(v.Object(
				v.ObjKV("url", v.String(v.StrMin(1))),
				v.ObjKV("comment", v.Optional(v.String())),
			)))),
			v.ObjKV("ghost", v.Optional(v.String())),
			v.ObjKV("ghost2", v.Optional(v.String())),
			v.ObjKV("meta", v.Object(
				v.ObjKeys(v.String()),
				v.ObjValues(v.Or(v.Number(v.NumMin(.01), v.NumMax(1.1)), v.String())),
			)),
		)),
	)

	if path, err := schema.Validate(JSON); err == nil {
		fmt.Println("Validation passed.")
	} else {
		fmt.Printf("Validation failed at %s. Error (%s)\n", path, err)
	}
}
