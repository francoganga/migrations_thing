package main

import (
	"fmt"
	"minsi/comparator"
	"minsi/platform"
	"minsi/schema"
)

func main() {

	cmp := comparator.Comparator{}
	cmp.Platform = &platform.MariadbPlatform{}

	schema1 := schema.New("production")

	schema1.AddTable(schema.Table{
		Name: "user",
		Columns: map[string]schema.Column{
			"username": {
				Name: "username",
			},
		},
	})

	schema2 := schema.New("production-v2")
	diff := cmp.CompareSchemas(schema1, schema2)
	fmt.Printf("diff=%+v\n", diff)
}

