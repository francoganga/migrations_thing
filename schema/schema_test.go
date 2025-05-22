package schema

import (
	"fmt"
	"testing"
)



func TestMain(t *testing.T) {


	oldSchema := New([]Table{})

	old := Table {
		Name: "users",
		Columns: map[string]Column{
			"id": { Name: "id", Type: "int", AutoIncrement: true },
			"username": { Name: "username", Type: "string" },
			"age": { Name: "age", Type: "int" },
		},
		PrimaryKeyName: "id",
	}

	oldSchema.AddTable(old)

	newTable := Table {
		Name: "users",
		Columns: map[string]Column{
			"identifier": { Name: "identifier", Type: "int", AutoIncrement: true },
			"username": { Name: "username", Type: "string" },
			"age": { Name: "age", Type: "int" },
		},
		PrimaryKeyName: "id",
	}

	postTable := Table {
		Name: "posts",
		Columns: map[string]Column{
			"id": { Name: "id", Type: "int", AutoIncrement: true },
			"category": { Name: "category", Type: "string" },
			"content": { Name: "content", Type: "string" },
		},
		PrimaryKeyName: "id",

	}


	newSchema := New([]Table{})

	newSchema.AddTable(newTable)
	newSchema.AddTable(postTable)



	diff := DiffSchemas(*oldSchema, *newSchema)

	fmt.Printf("diff: %+v\n", diff)


}
